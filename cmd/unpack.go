package cmd

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/forever-eight/archiver/pkg/compression"
	"github.com/forever-eight/archiver/pkg/compression/vlc"
)

const unpackedExtension = ".txt"

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "unpack file",
	Run:   unpack,
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc")
	err := unpackCmd.MarkFlagRequired("method")
	if err != nil {
		panic("method is required")
	}
}

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		log.Fatal("no file path")
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
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

	packed := decoder.Decode(data)

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
