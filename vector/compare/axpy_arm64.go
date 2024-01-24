package compare

func ArmAxpyUnsafe(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func ArmAxpyUnsafeX(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func ArmAxpyPointer(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func ArmAxpyPointerLoop(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func ArmAxpyPointerLoopX(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func ArmAxpyUnsafeXR4(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func ArmAxpyUnsafeInterleaveXR4(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func ArmAxpyPointerLoopXR4(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)

type armAxpyDecl struct {
	name string
	fn   func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
}

var armAxpyDecls = []armAxpyDecl{
	{name: "AxpyUnsafe", fn: ArmAxpyUnsafe},
	{name: "AxpyUnsafeX", fn: ArmAxpyUnsafeX},
	{name: "AxpyPointer", fn: ArmAxpyPointer},
	{name: "AxpyPointerLoop", fn: ArmAxpyPointerLoop},
	{name: "AxpyPointerLoopX", fn: ArmAxpyPointerLoopX},
	{name: "AxpyUnsafeXR4", fn: ArmAxpyUnsafeXR4},
	{name: "AxpyUnsafeInterleaveXR4", fn: ArmAxpyUnsafeInterleaveXR4},
	{name: "AxpyPointerLoopXR4", fn: ArmAxpyPointerLoopXR4},
}
