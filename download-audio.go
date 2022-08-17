package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Pilladian/logger"
	"github.com/kkdai/youtube/v2"
)

func downloadAudio(id string, d string, f string) error {
	logger.Info("create youtube client")
	client := youtube.Client{}

	logger.Info("download video information")
	video, err := client.GetVideo(id)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	format_id := 0
	for i, f := range formats {
		if strings.Contains(f.MimeType, "audio/mp4") {
			format_id = i
			break
		}
	}
	title := video.Title
	author := video.Author
	size := float64(formats[format_id].ContentLength) * 0.000001
	filepath := fmt.Sprintf("%s%s.mp3", d, video.Title)
	if f != "" {
		filepath = d + f
	}
	fmt.Printf("Title:\t %s\nAuthor:  %s\nSize:\t %.1fMB\nPath:\t%s\n", title, author, size, filepath)

	logger.Info("get video stream for specific format")
	stream, _, err := client.GetStream(video, &formats[len(formats)-3])
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
