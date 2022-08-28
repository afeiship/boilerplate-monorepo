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

## usage
```shell
# download mp3
$ ytbdown -3 https://www.youtube.com/watch?v=BTfqMWSsrOY
# download mp4
$ ytbdown -4 https://www.youtube.com/watch?v=BTfqMWSsrOY
# download bilibili video
$ ytbdown -4 https://www.bilibili.com/video/xxxxID
```