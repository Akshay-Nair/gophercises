package main

import "testing"

func TestM(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("error occured while calling the main function")
		}
	}()
	main()
}
