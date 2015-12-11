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
				handleRequest()
			}
		}()
	}
	select {}
}

func returnsErr() error {
	return errors.New("[err string]")
}

func instrument(err *error) func() {
	return func() {
		fmt.Fprintln(ioutil.Discard, "Error", (*err).Error())
	}
}

func methodWithError() (retErr error) {
	defer instrument(&retErr)()
	if err := returnsErr(); err != nil {
		return err
	}

	return nil
}

func handleRequest() {
	err := methodWithError()
	fmt.Fprintln(ioutil.Discard, "error", err.Error())
}
