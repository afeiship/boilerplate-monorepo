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
	cmd := exec.Command("yt-dlp", args...)
	fmt.Println("[🐠]: ", cmd)
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
	if errStr != "" {
		fmt.Printf("💥ERROR: %s", errStr)
	}
	fmt.Printf("🌈out: %s", outStr)
}

func YtbMpx(mpx, url string, keep bool, name string) {
	var args = []string{url}
	var output = fmt.Sprintf("%s.(ext)s", name)
	if keep {
		args = append(args, "-k")
	}

	if mpx == "mp4" {
		args = append(args, "-o", output, "-f", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/bestvideo+bestaudio")
	}

	if mpx == "mp3" {
		args = append(args, "-o", output, "--extract-audio", "--audio-format", "mp3", "--audio-quality", "0")
	}

	YtbExecute(args...)
}
