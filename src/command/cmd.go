package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var book, qword, verse string
var chapter int
var all bool

const (
	BookListUsage             = "bible_reader book -l"
	TextBookUsage             = "bible_reader text -b 创"
	TextBookChapterUsage      = "bible_reader text -b 创 -c 1"
	TextBookChapterVerseUsage = "bible_reader text -b 创 -c 1 -v (1-3或 1,3,5,7)"
	OutBookUsage              = "bible_reader out -b 创"
	OutBookChapterUsage       = "bible_reader out -b 创 -c 1"
	OutBookChapterVerseUsage  = "bible_reader out -b 创 -c 1 -v (1-3或 1,3,5,7)"
)

var rootCmd = &cobra.Command{
	Use:   "bible_read",
	Short: "read bible at command ",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	rootCmd.AddCommand(GetBookCommand())
	rootCmd.AddCommand(GetTextCommand())
	rootCmd.AddCommand(GetOutCommand())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
