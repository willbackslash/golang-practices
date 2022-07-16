package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(message string) {
	fmt.Println(message)
}

func printMessage(message string) {
	fmt.Println(message)
}
