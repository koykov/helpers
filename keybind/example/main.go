package main

import (
	"fmt"
	"github.com/koykov/helpers/keybind"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Implementation of concrete key catcher.
type ExampleCatch struct{}

// Just display the signal when it catch.
func (c ExampleCatch) Catch(signal string) error {
	fmt.Println("caught signal:", signal)
	return nil
}

func main() {
	var (
		catcher ExampleCatch
		err     error
		sigStop = make(chan os.Signal)
	)

	// Load hotkeys and initialize keybind.
	kb := keybind.NewKeybind(&catcher)
	if err = kb.LoadFromFile("hotkeys.json"); err != nil {
		log.Fatal(err)
	}
	if err = kb.Init(); err != nil {
		log.Fatal(err)
	}

	// Release keybind when application will stop.
	signal.Notify(sigStop, os.Interrupt, syscall.SIGTERM)
	signal.Notify(sigStop, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigStop
		err := kb.Release()
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		os.Exit(0)
	}()

	// Wait for hotkeys.
	kb.Wait()
}
