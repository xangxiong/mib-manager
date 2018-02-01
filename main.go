package main

import (
	"flag"
	"json"
	
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

var (
	AppName string
	BuiltAt string
	debug = flag.Bool("d", false, "enable debug mode")
	w *astilectron.Window
)

func main() {
	flag.Parse()
	astilog.FlagInit()

	// run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName: AppName,
			AppIconDarwinPath: "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug: *debug,
		Homepage: "index.html",
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astilectron.PtrStr("About")},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, iw *astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) {
			w = iw
			return nil
		},
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center: astilectron.PtrBool(true),
			Height: astilectron.PtrInt(700),
			Width: astilectron.PtrInt(700),
		},
		MessageHandler: handleMessages,
	});

	if err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "event.name":
		// unmarshal payload
		var s string
		if err = json.Unmarshal(m.Payload, &path); err != nil {
			payload = err.Error()
			return
		}

		payload = s + " world"
	}

	return
}