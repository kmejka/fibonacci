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

var result []int
func benchmarkFib(n int, c *C, fibService FibService) {
	var r []int
	for i := 0; i < c.N; i++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		elems, _ := fibService.CountNValues(n)
		r = elems
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func (s *TestSuite) BenchmarkFib1(c *C) {
	benchmarkFib(1, c, s.fibService)
}

func (s *TestSuite) BenchmarkFib10(c *C) {
	benchmarkFib(10, c, s.fibService)
}

func (s *TestSuite) BenchmarkFib100(c *C) {
	benchmarkFib(100, c, s.fibService)
}


func (s *TestSuite) BenchmarkFib1000(c *C) {
	benchmarkFib(1000, c, s.fibService)
}
