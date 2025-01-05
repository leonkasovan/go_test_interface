Test go interface  

Edit tags in Makefile:  
`go build -tags=opengl,opengles,vulkan .`

Edit initialization in main.go
```
Renderers = append(Renderers, &OpenGL{})
Renderers = append(Renderers, &OpenGLES{})
Renderers = append(Renderers, &Vulkan{})
```

Build and run  
```
make
./go_test_interface
```
