package main

import (
	"math"

	ray "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1000)
	screenHeight := int32(1000)

	ray.SetConfigFlags(ray.FlagMsaa4xHint)
	ray.InitWindow(screenWidth, screenHeight, "xtra")

	camera := ray.Camera{}
	camera.Position = ray.NewVector3(0.0, 10.0, 10.0)
	camera.Target = ray.NewVector3(0.0, 5.0, 0.0)
	camera.Up = ray.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0

	ray.SetTargetFPS(60)

	const OrbitRadius = 20.0
	propeller := ray.LoadModelFromMesh(ray.GenMeshCube(1, 1, 1))
	propellerBase := ray.LoadModelFromMesh(ray.GenMeshCylinder(0.8, 1.5, 30))

	for !ray.WindowShouldClose() {
		sn, cs := math.Sincos((float64)(ray.GetTime()))
		camera.Position = ray.NewVector3(OrbitRadius*(float32)(sn), 20.0, OrbitRadius*(float32)(cs))

		ray.BeginDrawing()

		ray.ClearBackground(ray.RayWhite)
		ray.BeginMode3D(camera)
		{
			ray.DrawGrid(20, 1.0)

			ray.DrawCube(ray.NewVector3(0, 5, 0), 2.0, 2.5, 16.0, ray.Blue)
			ray.DrawCube(ray.NewVector3(0, 5.2, 3), 16.0, 0.5, 4.0, ray.SkyBlue)

			ray.DrawCube(ray.NewVector3(0, 5.2, -6), 8.0, 0.2, 3.0, ray.SkyBlue)
			ray.DrawCube(ray.NewVector3(0, 7, -6), 0.2, 3, 3.0, ray.SkyBlue)

			ray.DrawSphere(ray.NewVector3(0, 6, 3.5), 1, ray.LightGray)
			ray.DrawCube(ray.NewVector3(0, 7, -6), 0.2, 3, 3.0, ray.LightGray)

			ray.DrawModelEx(propeller,
				ray.NewVector3(0, 5, 8.5),
				ray.NewVector3(0, 0, 1), float32(ray.GetTime()*1000),
				ray.NewVector3(1, 7, 0.1),
				ray.Orange,
			)

			ray.DrawModelEx(propellerBase,
				ray.NewVector3(0, 5, 7.5),
				ray.NewVector3(1, 0, 0), 90,
				ray.NewVector3(1, 1, 1),
				ray.Orange,
			)

			/*
				ray.DrawSphere(ray.NewVector3(-1.0, 0.0, -2.0), 1.0, ray.Green)
				ray.DrawSphereWires(ray.NewVector3(1.0, 0.0, 2.0), 2.0, 16, 16, ray.Lime)

				ray.DrawCylinder(ray.NewVector3(4.0, 0.0, -2.0), 1.0, 2.0, 3.0, 4, ray.SkyBlue)
				ray.DrawCylinderWires(ray.NewVector3(4.0, 0.0, -2.0), 1.0, 2.0, 3.0, 4, ray.DarkBlue)
				ray.DrawCylinderWires(ray.NewVector3(4.5, -1.0, 2.0), 1.0, 1.0, 2.0, 6, ray.Brown)

				ray.DrawCylinder(ray.NewVector3(1.0, 0.0, -4.0), 0.0, 1.5, 3.0, 8, ray.Gold)
				ray.DrawCylinderWires(ray.NewVector3(1.0, 0.0, -4.0), 0.0, 1.5, 3.0, 8, ray.Pink)

				 // Draw a grid
			*/
		}
		ray.EndMode3D()
		ray.EndDrawing()
	}

	ray.CloseWindow()
}
