# youtube-dl
YouTube Downloader written in Go

## Installation
Clone the repository from Github
```bash
git clone https://github.com/Pilladian/youtube-dl.git
```

Build the application
```bash
cd youtube-dl
go build -o youtube-dl *.go
```

## Usage
```bash
# ./youtube-dl -h

Usage of ./youtube-dl:
  -audio
        [ optionally ] Only download audio file
  -debug
        [ optionally ] Enable Debug Mode (Log Level 2, Log File std.out)
  -f string
        [ optionally ] Filename of output file (default YOUTUBE_TITLE)
  -lf string
        [ optionally ] Log File (default "./youtube-dl.log")
  -ll int
        [ optionally ] Log Level between 0-2 (default 0)
  -p string
        [ optionally ] Destination path of output file (default "./")
  -u string
        [ mandatory  ] YouTube URL of video
```