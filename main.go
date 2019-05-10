package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "adm",
		Short:        "A Download Manager!",
		SilenceUsage: true,
	}
	cmd.AddCommand(getVersion(), getFileFromUrl())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func getFileFromUrl() *cobra.Command {
	return &cobra.Command{
		Use:     "get",
		Aliases: []string{"down", "download", "grap"},
		Short:   "Download a specific file (url) from the Internet",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//cmd.Println("Print: " + strings.Join(args, " "))
			ext := args[0][len(args[0])-3:]
			switch ext {
			case "jpg", "png", "gif":
				cmd.Println("Downloading the image...")
			case "mp4", "mkv", "3gp":
				cmd.Println("Downloading the video...")
			default:
				cmd.Println("Downloading the file...")
			}
			fileUrl := args[0]
			fileName := filepath.Base(fileUrl)
			if err := DownloadFile(fileName, fileUrl); err != nil {
				panic(err)
			}
			cmd.Println("Image Downloaded.")
			return nil
		},
	}
}

func getVersion() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Get the current version number of ADM",
		Aliases: []string{"v", "V", "Version", "VERSION"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("A Download Manager v0.1")
			return nil
		},
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
