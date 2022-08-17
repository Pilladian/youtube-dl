package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Pilladian/logger"
	"github.com/kkdai/youtube/v2"
)

func downloadVideo(id string, d string, f string) error {
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
