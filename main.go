package main

type Renderer interface {
	GetName() string
	Init()
	Close()
}

var Renderers []Renderer

func Init(Renderer_Setting string) Renderer {
	Renderers = append(Renderers, &OpenGL{})
	Renderers = append(Renderers, &OpenGLES{})
	Renderers = append(Renderers, &Vulkan{})

	for _, r := range Renderers {
		if r.GetName() == Renderer_Setting {
			return r
		}
	}
	return nil
}

func main() {
	gfx := Init("Vulkan")
	if gfx == nil {
		panic("Renderer not found")
	} else {
		gfx.Init()
		gfx.Close()
	}
}
