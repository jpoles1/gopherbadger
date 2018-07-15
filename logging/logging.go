package logging

import (
	"os"

	"github.com/fatih/color"
)

//Success is a helper function to report a process finished successfully
func Success(taskDescription string) {
	color.Green("Success - " + taskDescription)
}

//Error is a helper function to check if an error has been produced, and respond appropriately
func Error(taskDescription string, err error) {
	if err != nil {
		color.Yellow("Error - " + taskDescription + ": " + err.Error())
	}
}

//Fatal is a helper function to check if a fatal error has been produced, and respond appropriately
func Fatal(taskDescription string, err error) {
	if err != nil {
		color.Red("Fatal Error - " + taskDescription + ": " + err.Error())
		os.Exit(1)
	}
}
