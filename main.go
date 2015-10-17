package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/odeke-em/go-uuid"
)

func getter(url string) (outfile string, err error) {
	dl, dlErr := http.Get(url)
	if dlErr != nil || dl == nil {
		err = dlErr
		return
	}

	outfile = fmt.Sprintf("%v", uuid.NewRandom())
	f, fErr := os.Create(outfile)
	if fErr != nil {
		err = fErr
		return
	}

	defer func() {
		if dl != nil && dl.Close {
			dl.Body.Close()
		}
	}()

	io.Copy(f, dl.Body)
	f.Close()

	return
}

func fReadLines(f io.Reader) chan string {
	strChan := make(chan string)
	go func() {
		defer func() {
			close(strChan)
		}()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {

			line := scanner.Text()
			line = strings.Trim(line, " ")
			line = strings.Trim(line, "\n")

			strChan <- line
		}
	}()

	return strChan
}

func main() {
	urlChan := fReadLines(os.Stdin)
	for url := range urlChan {
		outPath, err := getter(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v %q\n", err, outPath)
		} else {
			fmt.Println(outPath)
		}
	}
}
