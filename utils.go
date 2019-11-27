package go_tools

import (
	// "os"
	// "bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
	// "time"
)

func DoCmd(cmdStr string) string {
	cmds := strings.Split(cmdStr, " ")
	name := cmds[0]
	var sdtout bytes.Buffer
	var sdterr bytes.Buffer

	cmd := &exec.Cmd{
		Path:   name,
		Args:   append([]string{name}, cmds[1:]...),
		Stdout: &sdtout, //os.Stdout,
		Stderr: &sdterr, //os.Stderr,
	}
	if filepath.Base(name) == name {
		if lp, err := exec.LookPath(name); err != nil {
			// cmd.lookPathErr = err
		} else {
			cmd.Path = lp
		}
	}
	err := cmd.Run()
	// stdout, err := cmd.StdoutPipe() //指向cmd命令的stdout
	// cmd.Start()
	// content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	return string(sdtout.Bytes())

	// var out bytes.Buffer
	// cmd.Stdout = &out

	// errReader := bufio.NewReader(cmd.Stderr)
	// stdOutReader := bufio.NewReader(cmd.Stdout)
	// cmd := exec.Command("cmd", "/C", cmds)
	// ch := make(chan int, 1)
	// var opBytes []byte
	// go func() {
	// 	var err error
	// 	opBytes, err = cmd.Output()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	ch <- 1
	// }()
	// for {
	// 	// println(out.String())

	// 	if <-ch != 0 {
	// 		break
	// 	}
	// 	time.Sleep(time.Millisecond)
	// }

	// return string(opBytes)
}

func GetAllFile(pathname string) ([]string, error) {
	retSlice := []string{}
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			retSlice = append(retSlice, fi.Name())
		}
	}
	return retSlice, err
}
