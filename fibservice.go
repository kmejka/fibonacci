package main

import (
	"errors"
	"strconv"
)

type FibService interface {
	CountNValues(limit int) (FibElems, error)
}

type FibServiceImpl struct{}

type FibElems []int

func fibCounter(c chan int) {
	x, y := 1, 1
	for {
		c <- x
		x, y = y, x+y
	}
}

func (srv *FibServiceImpl) CountNValues(n int) (FibElems, error) {
	if n == 0 {
		return nil, errors.New("Amount of fibonacci values to be counted has to be larger than 0")
	}
	ret := make([]int, n)
	//channel contains max 10 elements at once
	c := make(chan int, 10)
	//todo close the fibCounter thread
	//quit := make(chan int)
	go fibCounter(c)
	for i := 0; i < n; i++ {
		elem := <-c
		ret[i] = elem
	}
	return ret, nil
}

func (elems FibElems) ToString() string {
	var ret string
	for index, elem := range elems {
		ret += strconv.Itoa(elem)
		if index != len(elems) -1 {
			ret += ", "
		}
	}
	return ret
}
