package main

import (
	"flag"
	"fmt"
)

// version is pgen version number
const version = "0.0.2"

// config holds pgen configuration settings
type config struct {
	kind string
}

// application holds pgen data
type application struct {
	config  config
	version string
}

func main() {
	// application configuration
	var cfg config

	// command line arguments
	flag.StringVar(&cfg.kind, "kind", "word", "Kind of password {word|random}")

	// parse command line arguments
	flag.Parse()

	// configure application
	app := &application{
		config:  cfg,
		version: version,
	}

	// switch over kind of password
	switch app.config.kind {
	case "word":
		fmt.Println(app.generateWord())
	case "random":
		fmt.Println(app.generateRandom())
	default:
		fmt.Println("Please enter valid arguments.")
		flag.PrintDefaults()
	}
}
