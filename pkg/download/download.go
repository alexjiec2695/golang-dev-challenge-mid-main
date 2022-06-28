package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Run(urlPath string) error {
	parse, err := url.Parse(urlPath)

	if err != nil {
		return fmt.Errorf("parsing url, url with problems: %v", err)
	}

	p := parse.Path

	if p == "/" || p == "" {
		return errors.New("URL is empty")
	}

	split := strings.Split(p, "/")

	if len(split) == 0 {
		return errors.New("file not exist")
	}

	name := split[len(split)-1]

	file, err := os.Create(name)

	if err != nil {
		return fmt.Errorf("creating file: %v", err)
	}

	resp, err := http.Get(urlPath)

	if err != nil {
		return fmt.Errorf("calling document: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("calling document: %v", resp.StatusCode)
	}

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		return err
	}

	defer func() {
		file.Close()
		resp.Body.Close()
	}()

	fmt.Printf("File %s with Size: %d", name, size)

	return nil

}
