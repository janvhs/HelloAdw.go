// Code from https://www.gtk.org/docs/getting-started/hello-world/ ported
// to Go and gotk4.
package app

import (
	"fmt"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// Application struct, holding the app's state.
// Embeds the gtk.Application to inherit their methods.
type Application struct {
	*gtk.Application
}

// Initializes a new Application.
func New() *Application {
	appId := "com.example.Hello"
	gtkApp := gtk.NewApplication(appId, gio.ApplicationFlagsNone)
	app := &Application{
		gtkApp,
	}

	app.ConnectActivate(app.activate)
	return app
}

// Get's called, when the "activate" signal fires.
func (a *Application) activate() {
	window := gtk.NewApplicationWindow(a.Application)
	window.SetTitle("Hello")
	window.SetDefaultSize(200, 200)

	button := gtk.NewButtonWithLabel("Print a message to the console")
	button.ConnectClicked(a.printHello)
	window.SetChild(button)
	
	window.Present()
}

func (a *Application) printHello() {
	fmt.Printf("A friendly hello in version %s from GTK4 and Go!\n", Version)
}
