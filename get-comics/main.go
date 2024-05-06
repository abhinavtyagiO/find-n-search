package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getComic(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	response, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant read file: %s\n", err)
		os.Exit(-1)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Can't read %d; got %d", i, response.StatusCode)
		return nil
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid body: %s\n", err)
		os.Exit(-1)
	}

	return body
}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		count  int
		fails  int
		data   []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; fails < 2; i++ {
		if data = getComic(i); data == nil {
			fails++
			continue
		}
		if count > 0 {
			fmt.Fprint(output, ",") // place a comma after every object
		}
		if _, err := io.Copy(output, bytes.NewBuffer(data)); err != nil {
			fmt.Fprintf(os.Stderr, "Stopped: %s\n", err)
			os.Exit(-1)
		}

		fails = 0
		count++

	}
	fmt.Fprintf(os.Stderr, "Read %d comics", count)

}
