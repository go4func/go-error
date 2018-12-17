package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type newError struct {
	errorCode int
	errorDesc string
}

func TestStringMapping() {
	s := `usbotgpsn!uijt-qmfbtf"`

	// log.Println(strings.Map(f, s))
	log.Printf("%*s %s", 10, "", strings.Map(func(r rune) rune {
		return r - 1
	}, s))
}

func (err *newError) Error() string {
	return fmt.Sprintln(err.errorCode, err.errorDesc)
}

func main() {
	// if err := f1(); err != nil {
	// 	fmt.Println(err)
	// 	if _, ok := errors.Cause(err).(*newError); ok {
	// 		fmt.Println("err is newError")
	// 	}
	// }

	TestStringMapping()
}

func f1() error {
	if err := f2(); err != nil {
		return errors.Wrap(err, "error from f2()")
	}
	return nil
}

func f2() error {
	return &newError{errorCode: 101, errorDesc: "error from f1()"}
}
