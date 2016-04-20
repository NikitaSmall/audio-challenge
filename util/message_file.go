package util

import (
	"io"
	"log"
	"os"
)

// SaveMessageFile saves provided file as a wav file
// and may return error. Message path stored in env variable
// (usually it stores near static)
func SaveMessageFile(file io.Reader) error {
	out, err := os.Create(os.Getenv("MESSAGE_FILE"))
	if err != nil {
		log.Print(err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
