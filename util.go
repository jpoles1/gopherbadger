package main

import "log"

func errCheck(taskDescription string, err error) {
	if err != nil {
		log.Println("Error w/ " + taskDescription + ": " + err.Error())
	}
}
