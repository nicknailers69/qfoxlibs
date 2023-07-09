package qfoxlibs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type QFoxTestSuite struct {
	suite.Suite
	TestData interface{}
}

func (suite *QFoxTestSuite) SetupTest() {
	suite.TestData = fmt.Sprintf("%c%c", Q_FOX, Qrationale)
}

func (suite *QFoxTestSuite) TestQfoxsymbol() {
	assert.Equal(suite.T(), QfoxSymbol(), suite.TestData)
}

func TestQFoxTestSuite(t *testing.T) {
	suite.Run(t, new(QFoxTestSuite))
}
