//go:generate go tool templ generate
//go:generate go run github.com/google/wire/cmd/wire

package main

func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	err = app.db.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	err = app.db.Migrate()
	if err != nil {
		panic(err)
	}

	app.webServer.StartWebServer()
}
