package adt_test

import (
	"adt"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (s *StackSuite) SetupTest() {
	s.stack = adt.NewStack()
}

func (s *StackSuite) TestEmpty() {
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty() {
	s.stack.Bury("red")

	s.False(s.stack.IsEmpty())
}

func (s *StackSuite) TestEmptySizeZero() {
	s.Zero(s.stack.Size())
}

func (s *StackSuite) TestTwoSize() {
	s.stack.Bury("red")
	s.stack.Bury("blue")

	s.Equal(2, s.stack.Size())
}

func (s *StackSuite) TestUnburySingleItem() {
	s.stack.Bury("red")

	s.Equal("red", s.stack.Unbury())
	s.Zero(s.stack.Size())
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestUnburyMutlipleItems() {
	s.stack.Bury("red")
	s.stack.Bury("blue")

	s.Equal("blue", s.stack.Unbury())
	s.Equal(1, s.stack.Size())

	s.Equal("red", s.stack.Unbury())
	s.Zero(s.stack.Size())
	s.True(s.stack.IsEmpty())
}

type StackSuite struct {
	suite.Suite
	stack *adt.Stack
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}
