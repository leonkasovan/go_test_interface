//go:build opengl

package main

import "fmt"

type OpenGL struct {
	fbo          uint32
	fbo_texture  uint32
	enableModel  bool
	enableShadow bool
}

func (r *OpenGL) GetName() string {
	return "OpenGL 3.2"
}

func (r *OpenGL) Init() {
	// Initialize OpenGL
	fmt.Println("OpenGL 3.2 initialized")
}

func (r *OpenGL) Close() {
	// Close OpenGL
	fmt.Println("OpenGL 3.2 closed")
}
