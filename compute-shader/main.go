package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	if err := withgl(run); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func withgl(fn func() error) error {
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("failed to run glfw init: %w", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(512, 512, "Image Processing", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create window: %w", err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return fmt.Errorf("init gl: %w", err)
	}

	return fn()
}

func run() error {
	fmt.Println("Renderer :", gl.GoStr(gl.GetString(gl.RENDERER)))
	fmt.Println("OpenGL   :", gl.GoStr(gl.GetString(gl.VERSION)))

	fmt.Printf("Work Group Info: %#v\n", getWorkGroupInfo())

	shader, err := compileShader(kernelSource)
	if err != nil {
		return fmt.Errorf("compile shader: %w", err)
	}
	defer gl.DeleteShader(shader)

	program, err := createProgram(shader)
	if err != nil {
		return fmt.Errorf("create program: %w", err)
	}
	defer gl.DeleteProgram(program)

	var maxTextureSize int32
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &maxTextureSize)
	fmt.Println("max texture size", maxTextureSize)

	inputData, _ := ioutil.ReadFile("input.png")
	xm, err := png.Decode(bytes.NewReader(inputData))
	if err != nil {
		return fmt.Errorf("create program: %w", err)
	}
	inputImage := convertToRGBA(xm)
	size := inputImage.Rect.Size()

	imageInput := createInputTexture(0, inputImage)
	defer gl.DeleteTextures(1, &imageInput)

	imageOutput := createOutputTexture(1, size)
	defer gl.DeleteTextures(1, &imageOutput)

	gl.UseProgram(program)
	gl.DispatchCompute(uint32(size.X), uint32(size.Y), 1)
	gl.MemoryBarrier(gl.SHADER_IMAGE_ACCESS_BARRIER_BIT)

	m := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	gl.GetTextureImage(imageOutput, 0, gl.RGBA, gl.UNSIGNED_BYTE, int32(len(m.Pix)), gl.Ptr(m.Pix))

	f, err := os.Create("output.png")
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	err = png.Encode(f, m)
	if err != nil {
		return fmt.Errorf("encode image: %w", err)
	}

	return nil
}

type WorkGroupInfo struct {
	MaxCount Vec3i
	MaxSize  Vec3i

	Invocations int32
}

type Vec3i struct{ X, Y, Z int32 }

func getWorkGroupInfo() WorkGroupInfo {
	var g WorkGroupInfo
	gl.GetIntegeri_v(gl.MAX_COMPUTE_WORK_GROUP_COUNT, 0, &g.MaxCount.X)
	gl.GetIntegeri_v(gl.MAX_COMPUTE_WORK_GROUP_COUNT, 1, &g.MaxCount.Y)
	gl.GetIntegeri_v(gl.MAX_COMPUTE_WORK_GROUP_COUNT, 2, &g.MaxCount.Z)

	gl.GetIntegeri_v(gl.MAX_COMPUTE_WORK_GROUP_SIZE, 0, &g.MaxSize.X)
	gl.GetIntegeri_v(gl.MAX_COMPUTE_WORK_GROUP_SIZE, 1, &g.MaxSize.Y)
	gl.GetIntegeri_v(gl.MAX_COMPUTE_WORK_GROUP_SIZE, 2, &g.MaxSize.Z)

	gl.GetIntegerv(gl.MAX_COMPUTE_WORK_GROUP_INVOCATIONS, &g.Invocations)
	return g
}

func convertToRGBA(m image.Image) *image.RGBA {
	if n, ok := m.(*image.RGBA); ok {
		return n
	}

	n := image.NewRGBA(m.Bounds())
	draw.Draw(n, n.Bounds(), m, image.Point{}, draw.Src)
	return n
}

func createInputTexture(unit uint32, m *image.RGBA) uint32 {
	var tex uint32
	gl.GenTextures(1, &tex)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	size := m.Rect.Size()
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA32F, int32(size.X), int32(size.Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(m.Pix))

	gl.BindImageTexture(unit, tex, 0, false, 0, gl.READ_ONLY, gl.RGBA32F)

	return tex
}

func createOutputTexture(unit uint32, size image.Point) uint32 {
	var tex uint32
	gl.GenTextures(1, &tex)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA32F, int32(size.X), int32(size.Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	gl.BindImageTexture(unit, tex, 0, false, 0, gl.WRITE_ONLY, gl.RGBA32F)

	return tex
}

func compileShader(source string) (uint32, error) {
	shader := gl.CreateShader(gl.COMPUTE_SHADER)

	sourceLen := int32(len(source))
	sourceStr, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, sourceStr, &sourceLen)
	free()

	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status != gl.TRUE {
		defer gl.DeleteShader(shader)

		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := make([]byte, logLength+1)
		gl.GetShaderInfoLog(shader, logLength, nil, &log[0])

		return 0, fmt.Errorf("CompileShader %v: %v", source, string(log))
	}

	return shader, nil
}

func createProgram(shader uint32) (uint32, error) {
	program := gl.CreateProgram()
	gl.AttachShader(program, shader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status != gl.TRUE {
		defer gl.DeleteProgram(program)

		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := make([]byte, logLength+1)
		gl.GetProgramInfoLog(program, logLength, nil, &log[0])

		return 0, fmt.Errorf("LinkProgram: %v", string(log))
	}

	return program, nil
}

//go:embed kernel.comp
var kernelSource string
