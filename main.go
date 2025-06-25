package main

import (
	filePkg "github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/internal/file/repository"
	"github.com/Takasakiii/ayanami/internal/file/service"
	"github.com/Takasakiii/ayanami/pkg/config"
	"github.com/Takasakiii/ayanami/pkg/cuid"
	"github.com/Takasakiii/ayanami/pkg/database"
	"github.com/Takasakiii/ayanami/pkg/sender"
	"github.com/Takasakiii/ayanami/pkg/server"
)

func main() {
	conf := config.GetConfig()
	sen, err := sender.NewS3Sender(&conf.Senders.S3)
	if err != nil {
		panic(err)
	}
	cuidGenerator, err := cuid.NewCuid()
	if err != nil {
		panic(err)
	}

	db := database.NewGormDatabase()
	err = db.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	err = db.Migrate(&filePkg.File{})
	if err != nil {
		panic(err)
	}
	repo := repository.NewFileRepository(db)

	file := service.NewService(&conf.File, &sen, cuidGenerator, &sen, repo)

	webServer := server.Server{
		Config: &conf.Server,
		File:   file,
	}

	webServer.StartWebServer()
}
