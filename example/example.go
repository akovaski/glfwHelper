package main

import (
	"fmt"
	"github.com/go-gl/glfw"
	"time"
	"github.com/akovaski/glfwHelper"
)

func main() {
	fmt.Print("GLFW Init Start.")
	glfw.Init()
	fmt.Println(" GLFW Init Done.")

	defer glfw.Terminate()
	defer fmt.Println("example Terminating.")

	fmt.Println("GLFW Set Hints.")
	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	fmt.Print("GLFW Open Window Start.")
	glfw.OpenWindow(640, 480, 8, 8, 8, 0, 0, 0, glfw.Windowed)
	glfw.SetWindowTitle("example")
	fmt.Println(" GLFW Open Window Done.")

	v1, v2, v3 := glfw.GLVersion()
	fmt.Printf("OpenGL version: %d.%d.%d\n", v1, v2, v3)
	fmt.Printf("GLFW version: %d.%d.%d\n", glfw.VersionMajor, glfw.VersionMinor, glfw.VersionRevision)

	glfw.SetSwapInterval(1)

	fps := time.Duration(30)
	fmt.Printf("Creating %d Hz Ticker.", fps)
	ticker := time.NewTicker(time.Second / fps)
	fmt.Printf(" %d Hz Ticker Created\n", fps)

	closedWindow := glfwHelper.WindowCloseChan()
	mouseButtons := glfwHelper.MouseButtonChan()
	mousePos := glfwHelper.MousePosChan()
	mouseWheel := glfwHelper.MouseWheelChan()
	keyButtons := glfwHelper.KeyChan()
	charButtons := glfwHelper.CharChan()

	for {
		select {
		case <-ticker.C:
			glfw.SwapBuffers()
		case <-closedWindow:
			return
		case input := <-mouseButtons:
			fmt.Println(input)
		case input := <-mousePos:
			fmt.Println(input)
		case input := <-mouseWheel:
			fmt.Println(input)
		case input := <-keyButtons:
			fmt.Println(input)
		case input := <-charButtons:
			fmt.Println(string(input))
		}
	}
}
