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

func instrument(name string, tags map[string]string, err *error) func() {
	return func() {
		fmt.Fprintln(ioutil.Discard, "error")
		tags["result"] = "notOK"
		fmt.Fprintln(ioutil.Discard, "Error", (*err).Error())
	}
}

func methodWithError() (err error) {
	defer instrument("name", map[string]string{"1": "1"}, &err)()
	return returnsErr()
}

func handleRequest() {
	err := methodWithError()
	fmt.Fprintln(ioutil.Discard, "error", err.Error())
}
