package day13

import (
	"fmt"

	"github.com/pkg/term"
)

type Signal int

const (
	SigString Signal = iota
	SigChar
	SigClose
)

type Key int

const (
	KeyUp Key = iota + 257
	KeyDown
	KeyLeft
	KeyRight
)

type KeyboardListener struct {
	lastString string
	lastKey    Key

	keyChan        chan Key
	sigChan        chan Signal
	stringRespChan chan string
	keyRespChan    chan Key

	moribund bool
}

func createKeyboardListener() *KeyboardListener {
	newKeyboardListener := new(KeyboardListener)
	newKeyboardListener.lastKey = -1
	newKeyboardListener.lastString = ""

	newKeyboardListener.keyChan = make(chan Key)
	newKeyboardListener.sigChan = make(chan Signal, 1)
	newKeyboardListener.stringRespChan = make(chan string)
	newKeyboardListener.keyRespChan = make(chan Key)

	newKeyboardListener.moribund = false

	return newKeyboardListener
}

func keyCatcher(kl *KeyboardListener) {
	for {
		if kl.moribund {
			return
		}
		k, _ := getStdinChar()
		kl.keyChan <- k
	}
}

func GetString(kl *KeyboardListener) string {
	kl.sigChan <- SigString
	return <-kl.stringRespChan
}

func GetChar(kl *KeyboardListener) Key {
	kl.sigChan <- SigChar
	return <-kl.keyRespChan
}

func Close(kl *KeyboardListener) {
	kl.sigChan <- SigClose
}

func KeyboardListen(kl *KeyboardListener) {
	go keyCatcher(kl)

	sendNextChar := false
	sendNextString := false
	select {
	case k := <-kl.keyChan:
		if k == Key('\n') {
			if sendNextString {
				kl.stringRespChan <- kl.lastString
				kl.lastString = ""
				sendNextString = false
			}
		} else {
			if k < 257 {
				kl.lastString = fmt.Sprintf("%v%c", kl.keyRespChan, k)
			}
			kl.lastKey = k
			if sendNextChar {
				kl.keyRespChan <- kl.lastKey
				kl.lastKey = -1
				sendNextChar = false
			}
		}

	case sig := <-kl.sigChan:
		switch sig {
		case SigChar:
			if kl.lastKey != -1 {
				kl.keyRespChan <- kl.lastKey
				kl.lastKey = -1
				sendNextChar = false
			} else {
				sendNextChar = true
			}
		case SigString:
			if kl.lastKey == Key('\n') {
				kl.stringRespChan <- kl.lastString
				kl.lastString = ""
				sendNextString = false
			} else {
				sendNextString = true
			}
		case SigClose:
			kl.moribund = true
			return
		}
	}
}

func getStdinChar() (k Key, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".
		if bytes[2] == 65 {
			// Up
			k = KeyUp
			//keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			k = KeyDown
			//keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			k = KeyRight
			//keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			k = KeyLeft
			//keyCode = 37
		}
	} else if numRead == 1 {
		k = Key(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}
