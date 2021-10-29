package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"

	"github.com/manifoldco/promptui"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	path := filepath.Join(os.Getenv("HOME"), ".ssh", "config")
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	re := regexp.MustCompile(`Host\s+(\S+)`)
	matches := re.FindAllStringSubmatch(string(bs), -1)
	if matches == nil {
		fmt.Println("not found hosts")
		fmt.Println(string(bs))
		return 0
	}

	var nonPatternHost []string
	for _, match := range matches {
		if !strings.Contains(match[1], "*") {
			nonPatternHost = append(nonPatternHost, match[1])
		}
	}

	prompt := promptui.Select{
		Label: "Select Hosts",
		Items: nonPatternHost,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	bin, err := exec.LookPath("ssh")
	if err != nil {
		fmt.Println(err)
		return 1
	}

	args := []string{"ssh", result}
	env := os.Environ()
	if err := syscall.Exec(bin, args, env); err != nil {
		fmt.Println(err)
		return 1
	}

	// never reach here
	return 0
}
