package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lankaiyun/kaiyunchain/lively"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	lexerCmd.Flags().StringP("liv", "l", "", ".liv file")
}

var lexerCmd = &cobra.Command{
	Use:   "lexer",
	Short: "Show the lexer output for a given .liv file",
	Long: `
This command allows you to see how a given .liv file is
split into tokens by our lexer.`,
	Run: func(cmd *cobra.Command, args []string) {
		Lexer(cmd)
	},
}

func Lexer(cmd *cobra.Command) {
	file, _ := cmd.Flags().GetString("liv")
	isHas := strings.HasSuffix(file, ".liv")
	if !isHas {
		color.Red("The filetype must be an .liv file!")
		fmt.Println()
		return
	}
	bs, err := os.ReadFile(file)
	if err != nil {
		color.Red("File read failure: %s", err)
		fmt.Println()
		return
	}
	liv := lively.NewLively()
	toks, err := liv.ContractToTokens(string(bs))
	if err != nil {
		color.Red("%s", err)
	}
	for i := 0; i < len(toks); i++ {
		fmt.Println(toks[i])
	}
}
