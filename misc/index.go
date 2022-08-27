package misc

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func YtbExecute(args ...string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	var errStdout, errStderr error
	cmd := exec.Command("youtube-dl", args...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func YtbMp4(url string) {
	YtbExecute(url, "-o", "%(title)s.%(ext)s", "-f", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/bestvideo+bestaudio")
}

func YtbMp3(url string) {
	YtbExecute(url, "-o", "%(title)s.%(ext)s", "-f", "bestaudio --extract-audio mp3")
}
