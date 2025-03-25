package main

import "github.com/rafaLino/couple-wishes-api/api"

func main() {
	api.NewApp().
		Initialize().
		StartupDatabase().
		ConfigEndpoints().
		Run()
}
