package cmd

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/forever-eight/archiver/pkg/vlc"
)

// todo: take extension from file
const unpackedExtension = ".txt"

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file with variable-length code",
	Run:   unpack,
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		log.Fatal("no file path")
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		log.Printf("open file error: %s", err)
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		log.Printf("read all data error: %s", err)
	}

	packed := vlc.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		log.Printf("write to file error: %s", err)
	}
}

// todo: do not repeat yourself
func unpackedFileName(path string) string {
	// full name
	filename := filepath.Base(path)
	// file extension
	ext := filepath.Ext(path)
	// delete extension
	basename := strings.TrimSuffix(filename, ext)

	return basename + unpackedExtension
}
