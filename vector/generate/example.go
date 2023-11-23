package main

import (
	"fmt"
	"os"
)

func main() {
	ctx := NewNaive(Config{
		Package: "f32",
		Type: Type{
			Name: "float32",
			Size: 4,
		},
		Unroll:  4,
		Pointer: true,
	})

	All(ctx)

	for _, file := range ctx.Files {
		fmt.Println("===", file.Path, "===")

		content, err := file.Formatted()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(string(content))
	}
}

func All(ctx Context) {
	Axpy(ctx.In("axpy.go"))
	Scale(ctx.In("scale.go"))
}

func Axpy(ctx File) {
	ctx.Func("AxpyUnitary(alpha $Type, xs, ys []$Type)",
		Iterate{
			Ranges: []It{
				Slice("ys", 0, 1),
				Slice("xs", 0, 1),
			},
			For: "$ys += alpha * $xs",
		})

	ctx.Func("AxpyUnitaryTo(dst []$Type, alpha $Type, xs, ys []$Type)",
		Iterate{
			Ranges: []It{
				Slice("dst", 0, 1),
				Slice("ys", 0, 1),
				Slice("xs", 0, 1),
			},
			For: "$dst = $ys + alpha * $xs",
		})

	ctx.Func("AxpyIncTo(dst []$Type, incDst, idst uintptr, alpha $Type, xs, ys []$Type, n, incx, incy, ix, iy uintptr)",
		Iterate{
			Ranges: []It{
				Range("i", 0, "n"),
				Slice("dst", "idst", "incDst"),
				Slice("xs", "ix", "incx"),
				Slice("ys", "iy", "incy"),
			},
			For: "$dst = $ys + alpha * $xs",
		})
}

func Scale(ctx File) {
	ctx.Func("ScalIncUnitaryTo(dst []$Type, incdst uintptr, alpha $Type, xs []$Type, n, incx uintptr)",
		Iterate{
			Ranges: []It{
				Range("i", 0, "n"),
				Slice("dst", 0, "incdst"),
				Slice("xs", 0, "incx"),
			},
			For: "$dst = alpha * $xs",
		})
}
