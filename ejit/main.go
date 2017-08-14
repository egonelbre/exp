//
//  WARNING
//
//  DO NOT USE IN PRODUCTION
//  IT IS VERY UNSAFE
//
//  AND MOSTLY FOR FUN
//
//  USE A PROPER PACKAGE FOR JIT INSTEAD OF WRITING YOUR OWN !!!
//
package main

import (
	"fmt"
	"math"
	"syscall"
	"time"
	"unsafe"
)

type Memory [1 << 8]float64

type Var byte
type Model struct {
	Vars   map[string]Var
	Eqs    map[string]*Eq
	Memory Memory

	lastVar Var
}

func NewModel() *Model {
	return &Model{
		Vars: make(map[string]Var, 512),
		Eqs:  make(map[string]*Eq, 512),
	}
}

func (m *Model) AddEq(name string, expr *Expr) {
	eq := &Eq{}
	m.Eqs[name] = eq
	eq.Compile(expr, m)
}

func (m *Model) nextVar() Var {
	m.lastVar++
	return m.lastVar
}

func (m *Model) Run(name string) float64 {
	eq, ok := m.Eqs[name]
	if !ok {
		panic("eq " + name + " not defined")
	}
	return eq.Run(m)
}

func (m *Model) SetVar(name string, value float64) Var {
	ix, ok := m.Vars[name]
	if !ok {
		ix = m.nextVar()
		m.Vars[name] = ix
	}
	m.Memory[ix] = value
	return ix
}

func (m *Model) GetVar(name string) (float64, Var) {
	ix, ok := m.Vars[name]
	if !ok {
		ix = m.nextVar()
		m.Vars[name] = ix
	}
	return m.Memory[ix], ix
}

type Eq struct {
	Expr   *Expr
	Code   []Statement
	Func   *Func
	Result Var
}

type Op byte

const (
	Add = Op('+')
	Sub = Op('-')
	Mul = Op('*')
	Div = Op('/')
)

type Statement struct {
	// D := A op B
	Op      Op
	D, A, B Var
}

// Compile assumes that *Expr represents a tree
func (eq *Eq) Compile(expr *Expr, model *Model) {
	eq.Expr = expr
	eq.Result = eq.compile(expr, model)

	fn := &Func{}
	fn.InitRAX()
	for _, stmt := range eq.Code {
		fn.BinaryOp(stmt.Op, stmt.D, stmt.A, stmt.B)
	}
	fn.Ret()
	fn.MarkExecutable()
	eq.Func = fn
}

func (eq *Eq) compile(expr *Expr, model *Model) Var {
	switch expr.Type {
	case ExprBinary:
		left := eq.compile(expr.A, model)
		right := eq.compile(expr.B, model)
		ix := model.nextVar()
		eq.Code = append(eq.Code, Statement{expr.Op, ix, left, right})
		return ix
	case ExprConst:
		ix := model.nextVar()
		model.Memory[ix] = expr.Value
		return ix
	case ExprVar:
		_, ix := model.GetVar(expr.VarName)
		return ix
	}
	panic("invalid expr")
}

func (eq *Eq) RunJIT(m *Model) float64 {
	mem := &m.Memory
	eq.Func.Call(mem)
	return mem[eq.Result]
}

func (eq *Eq) Run(m *Model) float64 {
	mem := &m.Memory
	for _, stmt := range eq.Code {
		mem[stmt.D] = stmt.Op.Eval(mem[stmt.A], mem[stmt.B])
	}
	return mem[eq.Result]
}

func (op Op) Eval(A, B float64) float64 {
	switch op {
	case Add:
		return A + B
	case Sub:
		return A - B
	case Mul:
		return A * B
	case Div:
		return A / B
	}
	return math.NaN()
}

type ExprType byte

const (
	ExprInvalid = ExprType(iota)
	ExprBinary
	ExprConst
	ExprVar
)

// Expr represents arbitrary ast node
type Expr struct {
	Type    ExprType
	VarName string
	Value   float64
	Op      Op
	A, B    *Expr
}

func (expr *Expr) Optimize() *Expr {
	switch expr.Type {
	case ExprBinary:
		left := expr.A.Optimize()
		right := expr.B.Optimize()

		if left.Type == ExprConst && right.Type == ExprConst {
			return &Expr{
				Type:  ExprConst,
				Value: expr.Op.Eval(left.Value, right.Value),
			}
		}
		return &Expr{
			Type: ExprBinary,
			Op:   expr.Op,
			A:    left,
			B:    right,
		}
	case ExprConst:
		return expr
	case ExprVar:
		return expr
	}
	panic("invalid expr")
}

func (expr *Expr) String() string {
	switch expr.Type {
	case ExprBinary:
		return fmt.Sprintf("(%v %v %v)", expr.A, string(expr.Op), expr.B)
	case ExprConst:
		return fmt.Sprintf("%v", expr.Value)
	case ExprVar:
		return expr.VarName
	}
	return "???"
}

// Helpers to make create expressions easier
func EOp(left *Expr, op Op, right *Expr) *Expr {
	return &Expr{Type: ExprBinary, Op: op, A: left, B: right}
}
func EAdd(left, right *Expr) *Expr { return EOp(left, Add, right) }
func ESub(left, right *Expr) *Expr { return EOp(left, Sub, right) }
func EMul(left, right *Expr) *Expr { return EOp(left, Mul, right) }
func EDiv(left, right *Expr) *Expr { return EOp(left, Div, right) }
func EVar(name string) *Expr       { return &Expr{Type: ExprVar, VarName: name} }
func EConst(value float64) *Expr   { return &Expr{Type: ExprConst, Value: value} }

func main() {
	// defer profile.Start(profile.CPUProfile).Stop()

	model := NewModel()

	// expr := "x + 2*y/(30-2/3) + (12-2)/5 + 1 + 2 + 3 + 4 + 5"
	// parsing left as an exercise :D
	expr := EAdd(
		EVar("x"),
		EAdd(
			EDiv(
				EMul(EConst(2), EVar("y")),
				ESub(EConst(30), EDiv(EConst(2), EConst(3))),
			),
			EAdd(
				EDiv(ESub(EConst(12), EConst(2)), EConst(5)),
				EAdd(EConst(1),
					EAdd(EConst(2),
						EAdd(EConst(3),
							EAdd(EConst(4), EConst(5))))),
			),
		),
	)

	fmt.Println(expr)
	expr = expr.Optimize()
	fmt.Println(expr)

	model.SetVar("x", 0)
	model.SetVar("y", 0)
	model.AddEq("testeq", expr)

	n := 1000 * 1000

	x := []float64{3, 4, 12, -3.4, 20}
	y := []float64{1, 2, 3, 4, 5}

	{
		// testing walking ASTs
		t := time.Now()
		result := float64(0)

		// cache the map lookups
		_, ix := model.GetVar("x")
		_, iy := model.GetVar("y")
		eq := model.Eqs["testeq"]

		for j := 0; j < n; j++ {
			model.Memory[ix] = x[3]
			model.Memory[iy] = y[3]
			result += eq.RunJIT(model)
		}
		fmt.Println(result)
		fmt.Println(time.Since(t))
	}

	{
		// testing just evaluating
		t := time.Now()
		result := float64(0)
		for j := 0; j < n; j++ {
			result += x[3] + 2*y[3]/(30-2.0/3.0) + (12-2)/5.0 + 1 + 2 + 3 + 4 + 5
		}
		fmt.Println(result)
		fmt.Println(time.Since(t))
	}
}

// must only use references, otherwise when the memory moves
type Func struct {
	body []byte
	Call func(*Memory)
}

func (fn *Func) append(code ...byte) { fn.body = append(fn.body, code...) }
func (fn *Func) append32(v uint32) {
	fn.append(byte(v>>0), byte(v>>8), byte(v>>16), byte(v>>24))
}

func (fn *Func) mov_rsp8_rax() { fn.append(0x48, 0x8b, 0x44, 0x24, 0x08) }
func (fn *Func) movsd_rax_X0(off uint32) {
	fn.append(0xf2, 0x0f, 0x10, 0x80)
	fn.append32(off)
}
func (fn *Func) movsd_rax_X1(off uint32) {
	fn.append(0xf2, 0x0f, 0x10, 0x88)
	fn.append32(off)
}
func (fn *Func) movsd_X0_rax(off uint32) {
	fn.append(0xf2, 0x0f, 0x11, 0x80)
	fn.append32(off)
}

func (fn *Func) add_X0_X1() { fn.append(0xf2, 0x0f, 0x58, 0xc1) }
func (fn *Func) sub_X0_X1() { fn.append(0xf2, 0x0f, 0x5c, 0xc1) }
func (fn *Func) mul_X0_X1() { fn.append(0xf2, 0x0f, 0x59, 0xc1) }
func (fn *Func) div_X0_X1() { fn.append(0xf2, 0x0f, 0x5e, 0xc1) }

func (fn *Func) InitRAX() { fn.mov_rsp8_rax() }
func (fn *Func) Ret()     { fn.append(0xc3) }
func (fn *Func) BinaryOp(op Op, dst, a, b Var) {
	fn.movsd_rax_X0(uint32(a) * 8)
	fn.movsd_rax_X1(uint32(b) * 8)
	switch op {
	case Add:
		fn.add_X0_X1()
	case Sub:
		fn.sub_X0_X1()
	case Mul:
		fn.mul_X0_X1()
	case Div:
		fn.div_X0_X1()
	}
	fn.movsd_X0_rax(uint32(dst) * 8)
}

func (fn *Func) MarkExecutable() {
	_, err := VirtualProtect(fn.body, 0x40)
	if err != nil {
		panic(err)
	}

	// OH GOD WHAT HAVE I DONE???
	type callstub struct{ fn func(*Memory) }
	var actual struct{ body **byte }
	pbody := &fn.body[0]
	actual.body = &pbody
	stub := (*callstub)(unsafe.Pointer(&actual))
	fn.Call = stub.fn
}

var (
	modkernel32        = syscall.NewLazyDLL("kernel32.dll")
	procVirtualProtect = modkernel32.NewProc("VirtualProtect")
)

func VirtualProtect(data []byte, newprotect uint) (oldprotect uint, err error) {
	var op uintptr
	r1, _, _ := procVirtualProtect.Call(uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)), uintptr(newprotect), uintptr(unsafe.Pointer(&op)))
	if r1 == 0 {
		err = fmt.Errorf("error")
	}
	return uint(op), err
}
