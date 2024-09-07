/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/

package main

import (
	"orders/cmd"
	"orders/internal/app/mylogger"
	"orders/internal/config"
	"orders/internal/lib/logger/sl"
	"orders/internal/storage/jsondb"
)

func main() {

	// init config: cleanenv
	cfg := config.MustLoad()

	// init logger: slog
	log := mylogger.SetupLogger(cfg.Env)

	// init storage: JSON
	storage, err := jsondb.GetStorage(cfg.StoragePath)
	if err != nil {
		log.Error("can't init storage", sl.Err(err))
		return
	}

	// init cobra-cli
	err = cmd.Execute(&storage, log)
	if err != nil {
		log.Error("can't init commands", sl.Err(err))
		return
	}

}
