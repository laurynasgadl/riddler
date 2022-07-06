package strings

import (
	"testing"

	"github.com/laurynasgadl/riddler/test"
	"github.com/stretchr/testify/suite"
)

type StringsTestSuite struct {
	test.BaseTestSuite

	Strings *Strings
}

func TestStringsTestSuite(t *testing.T) {
	suite.Run(t, new(StringsTestSuite))
}

func (suite *StringsTestSuite) SetupTest() {
	suite.Strings = NewStrings(suite.R)
}

func (suite *StringsTestSuite) Test_Generate_withoutOptions() {
	generated := suite.Strings.Generate()

	suite.Assert().Len(generated, int(DefaultLength))

	runes := SymbolsCharset + "-_"
	for _, r := range runes {
		suite.Assert().NotContains(generated, string(r))
	}
}

func (suite *StringsTestSuite) Test_Generate_withLength() {
	l := 123
	generated := suite.Strings.Generate(WithLength(uint(l)))

	suite.Assert().Len(generated, l)
}

func (suite *StringsTestSuite) Test_Generate_withCharset() {
	generated := suite.Strings.Generate(WithCharset(SymbolsCharset))

	runes := AlphaNumDashCharset
	for _, r := range runes {
		suite.Assert().NotContains(generated, string(r))
	}
}
