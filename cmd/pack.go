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

const packedExtension = ".vlc"

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")
	err := packCmd.MarkFlagRequired("method")
	if err != nil {
		panic("method is required")
	}
}

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder
	if len(args) == 0 || args[0] == "" {
		log.Fatal("no file path")
	}

	method := cmd.Flag("method").Value.String()
	switch method {
	case "vlc":
		encoder = vlc.New()
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

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
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
