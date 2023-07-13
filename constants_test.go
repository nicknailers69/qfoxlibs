package qfoxlibs

import (
	"fmt"
	"testing"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

/*
	@description just testing constants for fun :)
*/

func TestConstants(t *testing.T) {
	fmt.Println(oqs.EnabledKEMs())
	TestQFoxTestSuite(t)
}
