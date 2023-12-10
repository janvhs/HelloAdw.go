package main

import (
	"os"

	"example.com/hello/internal/app"
)

// The application code lives in the internal directory to keep the directories
// root tidy.
func main() {
	if code := mainE(); code != 0 {
		os.Exit(code)
	}
}

// The func mainE is required, because main calls os.Exit and will prevent
// deferred functions to run.
// This adds the possibility to add deferred function calls in this function.
func mainE() int {
	app := app.New()
	return app.Run(os.Args)
}
