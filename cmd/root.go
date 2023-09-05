/*
Copyright © 2023 Hava Pty Ltd support@hava.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teamhava/hava-ui-cli/printer"
	"github.com/teamhava/hava-ui-cli/version"
	"golang.org/x/term"
)

var (
	cfgFile string
	o       *printer.Output
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hava",
	Short: "A CLI to interface with the Hava platform",
	Long: `A CLI to interface with the Hava platform.

Hava CLI empowers engineers the ability to automate and interact
with the Hava platform.`,
	Version: version.String(),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Root Error Cmd: ", err)
		o.Close()
		os.Exit(1)
	} else {
		o.Close()
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Hava CLI Config file (default is ./.hava.env) or ~/.hava.env")

	// Output options, default is table output
	rootCmd.PersistentFlags().Bool("csv", false, "Will output command results as a CSV.")
	rootCmd.PersistentFlags().MarkHidden("csv")
	rootCmd.PersistentFlags().Bool("markdown", false, "Will output command results as a Markdown.")
	rootCmd.PersistentFlags().MarkHidden("markdown")
	rootCmd.PersistentFlags().Bool("html", false, "Will output command results as a HTML.")
	rootCmd.PersistentFlags().MarkHidden("html")
	rootCmd.PersistentFlags().Bool("json", false, "Will output command results as JSON.")
	rootCmd.PersistentFlags().MarkHidden("json")
	rootCmd.PersistentFlags().Bool("debug", false, "Print Debug messages")
	rootCmd.PersistentFlags().MarkHidden("debug")
	rootCmd.PersistentFlags().Bool("autoapprove", false, "Auto approve the hava command. --autoapprove, false by default")

}

// initConfig reads in config file and ENV variables if set.
// Leaving the ability to introduce config file in the future
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".hava" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/hava")
		viper.AddConfigPath("$HOME/.hava")
		viper.SetConfigName(".hava")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if havaToken := viper.GetString("HAVA_TOKEN"); len(havaToken) == 0 {

		if err := viper.ReadInConfig(); err != nil {

			if automationSet := viper.GetString("AUTOMATION"); len(automationSet) == 0 {
				// Interact with user to create a  ./hava.env config file
				createConfigFile()
			} else {
				fmt.Println("[Error] AUTOMATION variable set, with no HAVA_TOKEN \n\tIf running CLI interactively, unset AUTOMATION variable \n\tELSE\n\tPlease set HAVA_TOKEN environment variable or provide a valid hava token in a config file at ~/.hava.yaml or ./.hava.yaml")
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}

	}

	// Initialize output
	jsonOut, _ := rootCmd.Flags().GetBool("json")
	csvOut, _ := rootCmd.Flags().GetBool("csv")
	markdownOut, _ := rootCmd.Flags().GetBool("markdown")
	htmlOut, _ := rootCmd.Flags().GetBool("html")
	debugOut, _ := rootCmd.Flags().GetBool("debug")
	o = printer.New(jsonOut, csvOut, htmlOut, markdownOut, debugOut)
}

func createConfigFile() {
	fmt.Println("\n\n[CONFIG-Error] No hava_token found in config file at ~/.hava.yaml or ./.hava.yaml OR HAVA_TOKEN environment variable")
	fmt.Println("\n\n[CONFIG-CREATE] Creating hava CLI config file at ./.hava.yaml")
	havaToken := PasswordPrompt("[User Input] Enter your Hava API token:")
	havaEndpoint := StringPrompt("[User Input] Enter your Hava API Endpoint (Default https://api.hava.io):")
	havaConfigFile := StringPrompt("[User Input] Enter location of hava config file (Default ./.hava.yaml):")

	if len(havaEndpoint) == 0 {
		havaEndpoint = "https://api.hava.io"
	}

	if len(havaConfigFile) == 0 {
		havaConfigFile = "./.hava.yaml"
	}

	// Write config file
	confFile := []byte("---\nhava_token: " + havaToken + "\n" + "hava_endpoint: " + havaEndpoint)

	err := os.WriteFile(havaConfigFile, confFile, 0644)

	if err != nil {
		fmt.Println("error creating hava config file:" + err.Error())
		os.Exit(1)
	}

	fmt.Println(havaConfigFile + " config file has been created.")
	os.Exit(0)
}

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func PasswordPrompt(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(b)
		if s != "" {
			break
		}
	}
	fmt.Println()
	return s
}
