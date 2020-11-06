package main

import (
	"github.com/kataras/iris/v12"
	"log"
)

/**
 * sisdis was developed by Jose Buelvas (ihojose)
 *
 * @author        @ihojose
 * @author_url    dev.ihojose.com
 * @licence       Apache Licence v2.0
 * @year          2020
 * @donations     buymeacoff.ee/ihojose
 */
func main() {

	// Init Web Server
	app := iris.New()

	// Route APIs
	// TODO: API

	// Run Web Server
	if err := app.Listen(":8080"); err != nil {
		log.Panicln("Unexpected error in Recommendations API:", err)
	}
}
