package riddler

import (
	"testing"

	"github.com/laurynasgadl/riddler/strings"
	"github.com/laurynasgadl/riddler/test"
	"github.com/stretchr/testify/suite"
)

type RiddlerTestSuite struct {
	test.BaseTestSuite
}

func TestRiddlerTestSuite(t *testing.T) {
	suite.Run(t, new(RiddlerTestSuite))
}

func (suite *RiddlerTestSuite) Test_String() {
	length := 123
	charset := strings.SymbolsCharset

	generated := String(
		strings.WithLength(uint(length)),
		strings.WithCharset(charset),
	)

	suite.Assert().Len(generated, length)

	runes := strings.AlphaNumCharset
	for _, r := range runes {
		suite.Assert().NotContains(generated, string(r))
	}
}
