package collections

import (
	"io"
	"os"

	"github.com/cadigun/goeasyclient/easyhttp"
)

func CopyFileFromPath(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create the destination file
	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

func CopyFileFromURL(url, filename string) error {
	resp, err := easyhttp.Builder().WithRequestBody(url, nil, nil).Get()
	if err != nil {
		return err
	}
	respBody := resp.GetResponse().Body
	defer respBody.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, respBody)
	if err != nil {
		return err
	}
	return nil
}
