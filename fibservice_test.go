package main

import (
	"fmt"
	. "gopkg.in/check.v1"
	"log"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

type TestSuite struct {
	fibService FibService
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpSuite(c *C) {
	s.fibService = &FibServiceImpl{}
}

func (s *TestSuite) TestFibCalculation(c *C) {
	//given
	limit := 5

	//when
	elems, err := s.fibService.CountNValues(limit)

	//then
	c.Check(err, IsNil)
	c.Check(elems, DeepEquals, FibElems{1, 1, 2, 3, 5})
}

func (s *TestSuite) TestToString(c *C) {
	//given
	elems := &FibElems{1, 1, 2, 3, 5}

	//when
	str := elems.ToString()

	c.Check(str, Equals, "1, 1, 2, 3, 5")
}
