//go:generate go tool templ generate
//go:generate go tool wire

package main

import "github.com/Takasakiii/ayanami/internal/file"

func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	err = app.db.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	err = app.db.Migrate(&file.File{})
	if err != nil {
		panic(err)
	}

	app.webServer.StartWebServer()
}
