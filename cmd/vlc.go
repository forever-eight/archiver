package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const packedExtension = ".vlc"

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file with variable-length code",
	Run:   pack,
}

func init() {
	packCmd.AddCommand(vlcCmd)
}

func pack(_ *cobra.Command, args []string) {
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

	packed := ""
	fmt.Println(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		log.Printf("write to file error: %s", err)
	}
}

func packedFileName(path string) string {
	// full name
	filename := filepath.Base(path)
	// file extension
	ext := filepath.Ext(path)
	// delete extension
	basename := strings.TrimSuffix(filename, ext)

	return basename + packedExtension
}
