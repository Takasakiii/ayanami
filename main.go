package main

import (
	"github.com/Takasakiii/ayanami/internal/config"
	"github.com/Takasakiii/ayanami/internal/database"
	"github.com/Takasakiii/ayanami/internal/filemanager"
	"github.com/Takasakiii/ayanami/internal/sender"
	"github.com/Takasakiii/ayanami/internal/server"
)

func main() {

	conf := config.GetConfig()
	sen, err := sender.NewFileBin(&conf.Senders.FileBin)
	if err != nil {
		panic(err)
	}

	fileMng, err := filemanager.NewFileManager(&conf.File, &sen)
	if err != nil {
		panic(err)
	}

	db := database.GetDatabase()

	webServer := server.Server{
		Config:      &conf.Server,
		FileManager: &fileMng,
		Database:    db,
	}

	webServer.StartWebServer()
}
