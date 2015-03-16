package kangen

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writePid() {
	content := []byte(fmt.Sprintln(os.Getpid()))
	ioutil.WriteFile("/tmp/kangen.pid", content, os.ModePerm)
}
