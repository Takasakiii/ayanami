package main

import (
	"github.com/Takasakiii/ayanami/internal/config"
	"github.com/Takasakiii/ayanami/internal/filemanager"
	"github.com/Takasakiii/ayanami/internal/sender"
	"github.com/Takasakiii/ayanami/internal/server"
	"github.com/Takasakiii/ayanami/prisma/db"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	conf := config.GetConfig()
	sen, err := sender.NewFileBin(&conf.Senders.FileBin)
	if err != nil {
		panic(err)
	}

	fileMng, err := filemanager.NewFileManager(&conf.File, &sen)
	if err != nil {
		panic(err)
	}

	webServer := server.Server{
		Config:      &conf.Server,
		FileManager: &fileMng,
		Database:    client,
	}

	webServer.StartWebServer()
}
