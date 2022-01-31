package errors

// ErrorParsing sent when the function was given a response that it could not be processed
var ErrorParsing = NewCryptError("error parsing the response body")

// ErrorReading sent when the function has an issue reading the response stream
var ErrorReading = NewCryptError("error reading the response data")

// ErrorRequest sent when the function has an error sending the http request
var ErrorRequest = NewCryptError("error sending the http request")

