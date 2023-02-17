package main

import (
	"fmt"
	"os"
)

// version is pgen version number
const version = "0.0.3"

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

	// check if two arguments were provided
	if len(os.Args) > 2 {
		fmt.Println("Please enter valid arguments.")
		printArgs()
		return
	} else if len(os.Args) == 1 {
		cfg.kind = "word"
	} else {
		cfg.kind = os.Args[1]
	}

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
	case "pin":
		fmt.Println(app.generatePIN())
	case "help":
		fmt.Println("Pgen is a tool for generating passwords.\n")
		printArgs()
	default:
		fmt.Println(fmt.Sprintf("%s %s: unknown command", os.Args[0], os.Args[1]), "\n")
		printArgs()
	}
}

// printArgs prints all the acceptable arguments
func printArgs() {
	fmt.Println("Usage:")
	fmt.Println("        go <command>")

	fmt.Println("\nThe commands are:")
	fmt.Println("        word      non-existent words (default)")
	fmt.Println("        random    random characters")
	fmt.Println("        pin       four digit pin-code")
}
