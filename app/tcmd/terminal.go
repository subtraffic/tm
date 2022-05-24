package tcmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tm",
	Short: "a web load testing tool",
	// TraverseChildren: true,
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to tm")
	},
}

func SetFlags() {
	rootCmd.Flags().StringP("url", "u", "http://127.0.0.1", "")
	rootCmd.Flags().IntP("cpu", "c", runtime.GOMAXPROCS(-1), "concurrency of req")
	rootCmd.Flags().IntP("rate", "q", 0, "rate of req")
	rootCmd.Flags().StringP("duration", "z", "", "duration of test")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	SetFlags()
}
