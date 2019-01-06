package keybind

import (
	"encoding/json"
	"errors"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"io/ioutil"
	"os"
)

// Interface of receiver of key events.
type KeyCatcher interface {
	Catch(signal string)
}

// The hotkey structure.
type Hotkey struct {
	// One or combination of some keysymdefs. See all possible keys here
	// https://github.com/BurntSushi/xgbutil/blob/master/keybind/keysymdef.go
	// Please note, some of combination doesn't work, eg: Control-Alt-x, ...
	Key    string `json:"key"`
	// Signal key that will be sent to receiver. May be a random string.
	Signal string `json:"signal"`
}

// Main keybind struct.
type Keybind struct {
	catcher KeyCatcher
	hotkeys []*Hotkey
	xu      *xgbutil.XUtil
	cb      []keybind.KeyPressFun
}

// The constructor.
func NewKeybind(catcher KeyCatcher) *Keybind {
	kb := Keybind{
		catcher: catcher,
		hotkeys: make([]*Hotkey, 0),
		cb:      make([]keybind.KeyPressFun, 0),
	}
	return &kb
}

// Set the list of hotkeys.
func (kb *Keybind) SetHotkeys(hotkeys []*Hotkey) {
	kb.hotkeys = hotkeys
}

// Load list of hotkeys from the file.
func (kb *Keybind) LoadFromFile(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return errors.New("file not found: " + filename)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &kb.hotkeys); err != nil {
		return err
	}

	return nil
}

// Initialize and attach hotkeys.
func (kb *Keybind) Init() (err error) {
	kb.xu, err = xgbutil.NewConn()
	if err != nil {
		return err
	}
	keybind.Initialize(kb.xu)

	for _, hk := range kb.hotkeys {
		go func(kb *Keybind, hk *Hotkey, err *error) {
			cb0 := keybind.KeyPressFun(
				func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
					kb.catcher.Catch(hk.Signal)
				})
			*err = cb0.Connect(kb.xu, kb.xu.RootWin(), hk.Key, true)
		}(kb, hk, &err)
		if err != nil {
			return err
		}
	}

	return nil
}

// Wait for keypress events.
func (kb *Keybind) Wait() {
	xevent.Main(kb.xu)
}

// Release hotkeys attachments.
func (kb *Keybind) Release() (err error) {
	for _, hk := range kb.hotkeys {
		err = keybind.KeyReleaseFun(
			func(X *xgbutil.XUtil, e xevent.KeyReleaseEvent) {
				keybind.Detach(X, X.RootWin())
			}).Connect(kb.xu, kb.xu.RootWin(), hk.Key, true)
		if err != nil {
			return err
		}
	}

	return nil
}
