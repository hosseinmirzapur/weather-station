package utils

import "log"

func PanicErr(err error, str string) {
	if err != nil {
		panic(str)
	}
}

func FatalErr(err error, str string) {
	if err != nil {
		log.Fatal(str)
	}
}
