package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	dataCmd := exec.Command("date")

	dataOut, err := dataCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dataOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
