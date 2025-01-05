//go:build opengles

package main

import "fmt"

type OpenGLES struct {
	fbo          uint32
	fbo_texture  uint32
	enableModel  bool
	enableShadow bool
}

func (r *OpenGLES) GetName() string {
	return "OpenGLES 3.0"
}

func (r *OpenGLES) Init() {
	// Initialize OpenGLES
	fmt.Println("OpenGLES 3.0 initialized")
}

func (r *OpenGLES) Close() {
	// Close OpenGLES
	fmt.Println("OpenGLES 3.0 closed")
}
