package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/scene"
	"cheezewiz/pkg/ebitenwrapper"
	"image/color"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	c := config.Get()
	game := &ebitenwrapper.Game{
		Scene:           scene.Init(),
		WindowTitle:     c.Title,
		WindowScale:     c.ScaleFactor,
		Width:           c.Window.Width,
		Height:          c.Window.Height,
		BackgroundColor: color.NRGBA{0x00, 0x40, 0x80, 0xff},
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM)
	go func() {
		<-sigc
		game.Exit()
	}()

	game.Run()
}
