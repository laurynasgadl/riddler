package test

import (
	"math/rand"
	"time"

	"github.com/stretchr/testify/suite"
)

type BaseTestSuite struct {
	suite.Suite

	R *rand.Rand
}

func (suite *BaseTestSuite) SetupSuite() {
	suite.R = rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)
}
