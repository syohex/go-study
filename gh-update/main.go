package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func downloadURL(version string) string {
	return fmt.Sprintf("https://github.com/cli/cli/releases/download/v%s/gh_%s_linux_amd64.tar.gz", version, version)
}

func downloadAndExtract(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	gz, err := gzip.NewReader(res.Body)
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	tr := tar.NewReader(gz)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		if strings.HasSuffix(header.Name, "/bin/gh") {
			dest := path.Join(home, "bin", "gh")
			f, err := os.Create(dest)
			if err != nil {
				return err
			}

			defer f.Close()

			io.Copy(f, tr)

			if err := os.Chmod(dest, 0755); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("'gh' binary is not found in tar.gz")
}

func generateZshCompletion() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	gh := path.Join(home, "bin", "gh")
	cmd := exec.Command(gh, "completion", "--shell", "zsh")

	zshComp := path.Join(home, ".zsh", "completions", "_gh")
	f, err := os.Create(zshComp)
	if err != nil {
		return err
	}

	defer f.Close()

	cmd.Stdout = f

	return cmd.Run()
}

func _main() int {
	if len(os.Args) < 2 {
		fmt.Printf("Usage gh-update version\n")
		return 1
	}

	version := os.Args[1]
	url := downloadURL(version)

	if err := downloadAndExtract(url); err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Printf("Success to download and extract. version %s\n", version)

	if err := generateZshCompletion(); err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println("Success to update zsh completion file")
	return 0
}

func main() {
	os.Exit(_main())
}
