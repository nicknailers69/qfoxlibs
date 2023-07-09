package qfoxlibs

import (
	"fmt"
)

var Q_FOX = '🦊'
var Qrationale = 'ℚ'

func QfoxSymbol() string {
	return fmt.Sprintf("%c%c", Q_FOX, Qrationale)
}

func QfoxUnicodeFox() string {
	return fmt.Sprintf("%U", Q_FOX)
}

func QfoxUnicodeQ() string {
	return fmt.Sprintf("%U", Qrationale)
}
