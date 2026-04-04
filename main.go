//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
package main

import (
	"os"
	"path/filepath"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/utl"
)

var (
	app *portapps.App
)

func init() {
	var err error

	// Init app
	if app, err = portapps.New("wireshark-portable", "Wireshark"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = filepath.Join(app.AppPath, "Wireshark.exe")
	app.Args = []string{
		"-o",
		"gui.update.enabled:FALSE",
	}

	os.Setenv("WIRESHARK_APPDATA", app.DataPath)

	defer app.Close()
	app.Launch(os.Args[1:])
}
