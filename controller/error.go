package controller

import "fmt"

func check(err error) {
	if err != nil {
		fmt.Println("Route error:", err)
	}
}
