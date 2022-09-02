# ytbdown
> Youtube download based on youtube-dl.

## installation
```shell
# 安装 yt-dlp
pip install yt-dlp
# 或者用 brew 安装
brew install yt-dlp/taps/yt-dlp
brew install afeiship/jswork/ffmpeg
brew install afeiship/jswork/ytbdown
```

## update
```shell
cd /usr/local/Homebrew/Library/Taps/afeiship
rm -rf homebrew-jswork/
brew reinstall ytbdown
```

## usage
```shell
# download mp3
$ ytbdown -3 https://www.youtube.com/watch?v=BTfqMWSsrOY
$ ytbdown -3 https://www.youtube.com/watch?v=k2dGdKug7xY -n 周深-起風了
# download mp4
$ ytbdown -4 https://www.youtube.com/watch?v=BTfqMWSsrOY
# download bilibili video
$ ytbdown -4 https://www.bilibili.com/video/xxxxID
```

```conf
Download mp3/mp4 by youtube-dl/yt-dlp.

Usage:
  ytbdown [flags]

Flags:
  -h, --help          help for ytbdown
  -k, --keep          Download mpx keep original file.
  -3, --mp3           Download mp3 music.
  -4, --mp4           Download mp4 video.
  -n, --name string   Download filename. (default "%(title)s")
```