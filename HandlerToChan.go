package glfwHelper

import (
	"github.com/go-gl/glfw"
)

type DigiState struct {
	Id    int // identification of input
	State int //whether the input is pressed or depressed
}

type Point struct {
	X int
	Y int
}

func WindowCloseChan() chan bool {
	out := make(chan bool)
	glfw.SetWindowCloseCallback(func() int {
		go func() { out <- true }()
		return 1
	})
	return out
}

func MouseButtonChan() chan DigiState {
	out := make(chan DigiState)
	glfw.SetMouseButtonCallback(func(button, state int) {
		go func() { out <- DigiState{button, state} }()
	})
	return out
}

func MousePosChan() chan Point {
	out := make(chan Point)
	glfw.SetMousePosCallback(func(x, y int) {
		go func() { out <- Point{x, y} }()
	})
	return out
}

func MouseWheelChan() chan int {
	out := make(chan int)
	glfw.SetMouseWheelCallback(func(pos int) {
		go func() { out <- pos }()
	})
	return out
}

func KeyChan() chan DigiState {
	out := make(chan DigiState)
	glfw.SetKeyCallback(func(key, state int) {
		go func() { out <- DigiState{key, state} }()
	})
	return out
}

func CharChan() chan rune {
	out := make(chan rune)
	glfw.SetCharCallback(func(char, state int) {
		go func() { out <- rune(char) }()
	})
	return out
}