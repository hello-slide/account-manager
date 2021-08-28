package handler

import (
	"os"
	"strings"
)

var url = os.Getenv("API_URL")

var domain = strings.Split(url, "//")[1]
