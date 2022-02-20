package common

import "net/http"

var HttpClient = &http.Client{
	Timeout: DefaultHttpTimeout,
}

var HttpDebug = false
