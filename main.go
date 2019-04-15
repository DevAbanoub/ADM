package main

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "sidm",
		Short:        "Smart Internet Download Manager!",
		SilenceUsage: true,
	}
	cmd.AddCommand(downloadUrlCmd(), getFile())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func downloadUrlCmd() *cobra.Command {
	return &cobra.Command{
		Use: "get",
		RunE: func(cmd *cobra.Command, args []string) error {
			// now := time.Now()
			// prettyTime := now.Format(time.RubyDate)
			cmd.Println("Downloading...")
			return nil
		},
	}
}

func getFile() *cobra.Command {
	return &cobra.Command{
		Use: "getfile",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Image Downloading..")
			//edit to make it dynamic
			fileUrl := "https://4.bp.blogspot.com/-TOpMgE8dQKc/W-s52dGFFXI/AAAAAAAAyaU/Z1OHgo2ZSa88E3IaF9ztZ5UsnISSi82IgCK4BGAYYCw/s1600/logo.png"
			if err := DownloadFile("logo.jpg", fileUrl); err != nil {
				panic(err)
			}
			cmd.Println("Image Downloaded.")
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
