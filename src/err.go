package main

import "log"

func checkErr(iErr error) {
	if iErr != nil {
		log.Fatal(iErr)
	}
}

func checkWarn(iErr error) {
	if iErr != nil {
		log.Printf(iErr.Error())
	}
}
