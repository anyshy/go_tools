package go_tools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWavToMp3(t *testing.T) {
	// var privateKey *ecdsa.PrivateKey
	filepath := "G:\\download\\2015沈阳\\"
	files, err := GetAllFile(filepath)
	ch := make(chan int, cap(files))
	for _, v := range files {
		inputFile := filepath + v
		outputFile := filepath + v + ".mp3"
		go WavToMp3(inputFile, outputFile, ch)
	}
	for i := 0; i < cap(ch); i++ {
		<-ch
	}
	assert.True(t, err == nil)
	time.Sleep(time.Second)
}
