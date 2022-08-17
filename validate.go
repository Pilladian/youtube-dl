package main

import (
	"errors"
	"fmt"
	"regexp"
)

func validateURL(u string) (string, error) {
	re, _ := regexp.Compile(` https:\/\/www\.youtube\.com\/watch\?v=\w+ `)
	if !re.Match([]byte(fmt.Sprintf(" %s ", u))) {
		return "", errors.New(fmt.Sprintf("Incorrect URL \"%s\" - Example: https://www.youtube.com/watch?v=VIDEO_ID", u))
	}
	return u, nil
}

func validatePath(p string) (string, error) {
	re, _ := regexp.Compile(` (\.{0,2}\/)?(\.{1,2}\/)*(\w+\/)* `)
	if !re.Match([]byte(fmt.Sprintf(" %s ", p))) {
		return "", errors.New(fmt.Sprintf("Incorrect destination \"%s\"", p))
	}
	return p, nil
}

func validateOut(f string, audio bool) (string, error) {
	if f == "" {
		return "", nil
	}

	if audio {
		re, _ := regexp.Compile(` [\w-() _]+\.mp3 `)
		if !re.Match([]byte(fmt.Sprintf(" %s ", f))) {
			return "", errors.New("Incorrect filename for output file. Example: my-audio.mp3")
		}
		return f, nil
	} else {
		re, _ := regexp.Compile(` [\w-() _]+\.mp4 `)
		if !re.Match([]byte(fmt.Sprintf(" %s ", f))) {
			return "", errors.New("Incorrect filename for output file. Example: my-video.mp4")
		}
		return f, nil
	}
}
