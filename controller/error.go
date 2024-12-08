package controller

import "fmt"

func check(err error) {
	if err != nil {
		fmt.Println("Controller error:", err)
	}
}
