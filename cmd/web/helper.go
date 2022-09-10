package main

import (
	"strconv"
)

func (a *application) toInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	return n, err
}
