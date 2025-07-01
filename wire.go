//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Takasakiii/ayanami/internal/file"
	"github.com/Takasakiii/ayanami/internal/file/repository"
	"github.com/Takasakiii/ayanami/internal/file/service"
	"github.com/Takasakiii/ayanami/pkg/config"
	"github.com/Takasakiii/ayanami/pkg/cuid"
	"github.com/Takasakiii/ayanami/pkg/database"
	"github.com/Takasakiii/ayanami/pkg/sender"
	"github.com/Takasakiii/ayanami/pkg/server"
	"github.com/google/wire"
)

type App struct {
	webServer *server.Server
	db        database.Database
}

var configSet = wire.NewSet(config.GetConfig)
var senderSet = wire.NewSet(
	configSet,
	sender.NewS3Sender,
	wire.Bind(new(sender.Sender), new(*sender.S3Sender)),
	wire.Bind(new(sender.Downloader), new(*sender.S3Sender)))

var databaseSet = wire.NewSet(
	database.NewGormDatabase,
	wire.Bind(new(database.Database), new(*database.GormDatabase)))

var fileRepositorySet = wire.NewSet(
	databaseSet,
	repository.NewFileRepository,
	wire.Bind(new(file.Repository), new(*repository.FileRepository)))

var cuidSet = wire.NewSet(cuid.NewCuid, wire.Bind(new(cuid.Generator), new(*cuid.Cuid)))
var fileServiceSet = wire.NewSet(
	senderSet,
	cuidSet,
	fileRepositorySet,
	service.NewService,
	wire.Bind(new(file.Service), new(*service.FileService)))

var webServerSet = wire.NewSet(fileServiceSet, server.NewServer)

func newApp(webServer *server.Server, db database.Database) *App {
	return &App{
		webServer: webServer,
		db:        db,
	}
}

func InitializeApp() (*App, error) {
	wire.Build(webServerSet, newApp)
	return nil, nil
}
