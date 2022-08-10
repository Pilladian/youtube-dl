package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/Pilladian/logger"
	"github.com/kkdai/youtube/v2"
)

func download(id string, d string, f string) error {
	logger.Info("create youtube client")
	client := youtube.Client{}

	logger.Info("download video information")
	video, err := client.GetVideo(id)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	title := video.Title
	quality := formats[0].QualityLabel
	if quality != "" {
		quality = formats[0].Quality
	}
	author := video.Author
	size := float64(formats[0].ContentLength) * 0.000001
	filepath := fmt.Sprintf("%s%s.mp4", d, video.Title)
	if f != "" {
		filepath = d + f
	}
	fmt.Printf("Title:\t %s\nAuthor:  %s\nQuality: %s\nSize:\t %.1fMB\nPath:\t%s\n", title, author, quality, size, filepath)

	logger.Info("get video stream for specific format")
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return err
	}

	logger.Info("create output file")
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	logger.Info("copy stream to output file")
	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}
	return nil
}

func extractID(url string) string {
	return strings.Split(url, "v=")[1]
}

func validateURL(u string) (string, error) {
	re, _ := regexp.Compile(` https:\/\/www\.youtube\.com\/watch\?v=\w+ `)
	if !re.Match([]byte(fmt.Sprintf(" %s ", u))) {
		return "", errors.New(fmt.Sprintf("Incorrect URL \"%s\"\nExample: https://www.youtube.com/watch?v=VIDEO_ID", u))
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

func validateOut(f string) (string, error) {
	if f == "" {
		return "", nil
	}

	re, _ := regexp.Compile(` [\w-() _]+\.mp4 `)
	if !re.Match([]byte(fmt.Sprintf(" %s ", f))) {
		return "", errors.New("Incorrect filename for output file. Example: my-video.mp4")
	}
	return f, nil
}

func main() {
	var url string
	var path string
	var filename string
	var loglevel int
	var logfile string
	var debug bool
	flag.StringVar(&url, "u", "", "[ mandatory  ] YouTube URL of video")
	flag.StringVar(&path, "p", "./", "[ optionally ] Destination path of output file")
	flag.StringVar(&filename, "f", "", "[ optionally ] Filename of output file (default \"YOUTUBE_TITLE.mp4\")")
	flag.IntVar(&loglevel, "ll", 0, "[optionally] Log Level between 0-2 (default 0)")
	flag.StringVar(&logfile, "lf", "./youtube-dl.log", "[optionally] Log File")
	flag.BoolVar(&debug, "debug", false, "[optionally] Enable Debug Mode (Log Level 2, Log File std.out)")
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
		logger.Error(url_err.Error())
		os.Exit(1)
	}
	logger.Info("validating user input: path")
	path, dest_err := validatePath(path)
	if dest_err != nil {
		logger.Error(dest_err.Error())
		os.Exit(1)
	}
	logger.Info("validating user input: filename")
	filename, output_err := validateOut(filename)
	if output_err != nil {
		logger.Error(output_err.Error())
		os.Exit(1)
	}

	logger.Info("extracting video id from url")
	id := extractID(url)
	if download_err := download(id, path, filename); download_err != nil {
		logger.Error(download_err.Error())
	}
}
