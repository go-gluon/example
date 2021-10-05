package main

//go:generate gluon-generator

import (
	_ "github.com/go-gluon/fiber"
	_ "github.com/go-gluon/gluon"
	_ "github.com/go-gluon/zerolog"
)

// func start() {
// 	app := fiber.New(fiber.Config{
// 		DisableStartupMessage: true,
// 	})
// 	app.Mount("/user", NewUserRestController())
// }
