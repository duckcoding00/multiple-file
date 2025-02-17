package api

import "github.com/duckcoding00/multiple-file/internal/handler"

func InitServer() {
	handler := handler.NewHandler()
	config := AppConfig{
		addr: ":8080",
		h:    handler,
	}

	app := NewApp(config)
	app.RegisterRouter()
	app.Run()
}
