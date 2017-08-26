package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type operation int

const (
	opNone operation = iota
	opPlus
	opMinus
	opMultiply
	opDivide
)

func calculate(a, b int, op operation) (int, error) {
	switch op {
	// TODO calculate here
	default:
		return 0, fmt.Errorf("invalid operation: %s", op)
	}
}

func getValue(buf *bytes.Buffer) (int, error) {
	valString := buf.String()
	value, err := strconv.Atoi(valString)
	if err != nil {
		return 0, fmt.Errorf("error converting value: %s", err)
	}

	return value, nil
}
