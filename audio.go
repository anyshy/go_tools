package go_tools

import (
	"fmt"
)

func WavToMp3(inputFilePath string, outputFilePath string, ch chan int) {
	cmd := fmt.Sprintf("ffmpeg.exe -i %s -f mp3 -acodec libmp3lame -y %s", inputFilePath, outputFilePath)
	res := DoCmd(cmd)
	println(res)
	ch <- 1
}
