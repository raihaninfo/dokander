package main

import (
	"strconv"
)

// This function helps to convert string to int
func (a *application) toInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	return n, err
}
