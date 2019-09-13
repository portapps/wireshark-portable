//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("wireshark-portable", "Wireshark"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "Wireshark.exe")
	app.Args = []string{
		"-o",
		"gui.update.enabled:FALSE",
	}

	utl.OverrideEnv("WIRESHARK_APPDATA", app.DataPath)

	app.Launch(os.Args[1:])
}
