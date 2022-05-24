package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
					return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
					return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
					return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
					return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize npelection",
	Long: `Sends npelection to the preferred path for respective platforms`,
	Run: func(cmd *cobra.Command, args []string) {
		ext, err := os.Executable()
    if err != nil {
        panic(err)
    }
		fn := filepath.Base(ext)
    fp := filepath.Clean(ext)
		fd := filepath.Dir(ext)

		r := runtime.GOOS
		switch r {
		case "windows":
			cmd := exec.Command("setx", "path", "%path%;"+fd)
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		case "linux", "darwin":
			copy(fp, "/usr/local/bin/"+fn)
			os.Chmod("/usr/local/bin/"+fn, 777)
		default:
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}