// Copyright © 2018 Francesc Campoy <francesc@sourced.tech>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/src-d/engine/cmd/srcd/daemon"
)

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "srcd",
	Short: "The Code as Data solution by source{d}",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.srcd.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "if true, log all of the things")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".srcd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".srcd")
	}

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var logMsgRegex = regexp.MustCompile(`.*msg="(.+)"`)

func logAfterTimeout(msg string, timeout time.Duration) func() {
	d := &Defered{
		Timeout: timeout,
		Msg:     msg,
	}

	return d.Print()
}

func logAfterTimeoutWithSpinner(msg string, timeout time.Duration, spinnerInterval time.Duration) func() {
	d := &Defered{
		Timeout:         timeout,
		Msg:             msg,
		Spinner:         true,
		SpinnerInterval: spinnerInterval,
	}

	return d.Print()
}

func logAfterTimeoutWithServerLogs(msg string, timeout time.Duration) func() {
	d := &Defered{
		Timeout: timeout,
		Msg:     msg,
		InputFn: readDaemonLogs,
	}

	return d.Print()
}

func readDaemonLogs(stop <-chan bool) <-chan string {
	logs, err := daemon.GetLogs()
	if err != nil {
		logrus.Errorf("could not get logs from server container: %v", err)
		return nil
	}

	ch := make(chan string)
	go func() {
		defer logs.Close()
		scanner := bufio.NewScanner(logs)

		c := make(chan bool)
		scan := func() {
			c <- scanner.Scan()
		}

		go scan()
		for {
			select {
			case <-stop:
				close(ch)
				return
			case more := <-c:
				if !more {
					close(ch)
					if err := scanner.Err(); err != nil && err != context.Canceled {
						logrus.Errorf("can't read logs from server: %s", err)
					}

					return
				}

				match := logMsgRegex.FindStringSubmatch(scanner.Text())
				if len(match) == 2 {
					ch <- match[1]
				}

				go scan()
			}
		}

	}()

	return ch
}
