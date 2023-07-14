package qfoxlibs

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/nicknailers69/qfoxlibs/pkg/utils"
)

/*
	@description just testing constants for fun :)
*/

func TestConstants(t *testing.T) {

	id, _ := utils.NewID()
	spew.Dump(id.String())
	spew.Dump(id.Base64())

	TestQFoxTestSuite(t)
}
