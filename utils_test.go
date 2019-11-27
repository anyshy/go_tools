package go_tools

import (
	"testing"
	"time"
	// "github.com/stretchr/testify/assert"
)

func TestDoCmd(t *testing.T) {
	result := DoCmd("dir c:\\ ")
	println("result:" + result)
	// assert.True(t, err == nil)
	time.Sleep(time.Second)
}
