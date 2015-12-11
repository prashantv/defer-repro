package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				methodWithError()
			}
		}()
	}
	select {}
}

func instrument(err *error) func() {
	return func() {
		fmt.Fprintln(ioutil.Discard, "Error", (*err).Error())
	}
}

func methodWithError() (err error) {
	defer instrument(&err)()
	return errors.New("error")
}
