package main

import "strings"

func extractID(url string) string {
	return strings.Split(url, "v=")[1]
}
