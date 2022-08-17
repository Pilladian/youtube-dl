package main

import (
	"flag"

	"github.com/Pilladian/logger"
)

func main() {
	var url string
	var path string
	var filename string
	var audio bool
	var loglevel int
	var logfile string
	var debug bool

	flag.StringVar(&url, "u", "", "[ mandatory  ] YouTube URL of video")
	flag.StringVar(&path, "p", "./", "[ optionally ] Destination path of output file")
	flag.StringVar(&filename, "f", "", "[ optionally ] Filename of output file (default YOUTUBE_TITLE)")
	flag.BoolVar(&audio, "audio", false, "[ optionally ] Only download audio file")
	flag.IntVar(&loglevel, "ll", 0, "[ optionally ] Log Level between 0-2 (default 0)")
	flag.StringVar(&logfile, "lf", "./youtube-dl.log", "[ optionally ] Log File")
	flag.BoolVar(&debug, "debug", false, "[ optionally ] Enable Debug Mode (Log Level 2, Log File std.out)")
	flag.Parse()

	if debug {
		logger.SetLogLevel(2)
	} else {
		logger.SetLogLevel(loglevel)
		logger.SetLogFilename(logfile)
	}

	logger.Info("---- Starting new Youtube Download ----")
	logger.Info("cli arguments parsed successfully")

	logger.Info("validating user input: URL")
	url, url_err := validateURL(url)
	if url_err != nil {
		logger.Fatal(url_err.Error())
	}
	logger.Info("validating user input: path")
	path, dest_err := validatePath(path)
	if dest_err != nil {
		logger.Fatal(dest_err.Error())
	}
	logger.Info("validating user input: filename")
	filename, output_err := validateOut(filename, audio)
	if output_err != nil {
		logger.Fatal(output_err.Error())
	}

	logger.Info("extracting video id from url")
	id := extractID(url)

	if audio {
		logger.Info("downloading audio only")
		if download_err := downloadAudio(id, path, filename); download_err != nil {
			logger.Error(download_err.Error())
		}
	} else {
		logger.Info("downloading video")
		if download_err := downloadVideo(id, path, filename); download_err != nil {
			logger.Error(download_err.Error())
		}
	}
}
