package log

import "log"

// LogMessage logs a message
//
//	filename: the file name where the error occurred
//	funcName: the function name where the error occurred
//	str: the message to log
func LogMessage(filename, funcName string, str interface{}) {
	log.Printf("Message in %s/%s: %s\n\n", filename, funcName, str)
}
