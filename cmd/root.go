package cmd

import (
	"log"
	"os"
	"time"

	"github.com/narsuf/dns-go-matic/pkg/config"
	dnsomatic "github.com/narsuf/dns-go-matic/pkg/dns-o-matic"
	"github.com/spf13/cobra"
)

var (
	loadedConfig   config.Config
	user, password string
	logFile        string
	sleep          uint
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "DNS-GO-MATIC",
	Short: "Service designed to notify your IP to DNS-O-Matic.",
	Long: `DNS-GO-MATIC

Service designed to notify your IP to DNS-O-Matic.

In order to work the program requires 2 env variables:
DNS_O_MATIC_USER (overriden by flag "--user -u")
DNS_O_MATIC_PASSWORD (overriden by flag "--password -p")

An optional env variable to specify a file to write logs to can be configured:
DNS_O_MATIC_LOG_FILE (overriden by flag "--log-file -l")
If no log file is configured, logs will be written on the default output

To install cronjob:
(crontab -l 2>/dev/null; echo "* * * * * path/to/dns-o-matic -u user -p password") | crontab -

To uninstall cronjob:
crontab -l| grep -v "path/to/dns-o-matic" | crontab -

The program can run in "no-cron" mode. If the flag --sleep -s is provided the program will run
in an infinite loop with a sleep time of the --sleep -s minutes`,
	Run: func(cmd *cobra.Command, args []string) {
		if user != "" {
			loadedConfig.User = user
		}
		if password != "" {
			loadedConfig.Password = password
		}
		if logFile != "" {
			loadedConfig.LogFile = logFile
		}

		if loadedConfig.LogFile != "" {
			f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Printf("error opening file %s %s", logFile, err.Error())
			} else {
				log.SetOutput(f)
			}
			defer f.Close()
		}

		if loadedConfig.User == "" {
			log.Println("a user for DNS-O-MATIC is mandatory")
			os.Exit(1)
		}
		if loadedConfig.Password == "" {
			log.Println("a password for DNS-O-MATIC is mandatory")
			os.Exit(1)
		}

		if sleep != 0 {
			for {
				dnsomatic.UpdateIP(loadedConfig.User, loadedConfig.Password)
				time.Sleep(time.Duration(sleep) * time.Minute)
			}
		} else {
			dnsomatic.UpdateIP(loadedConfig.User, loadedConfig.Password)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "DNS-O-MATIC username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "DNS-O-MATIC password")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log-file", "l", "", "File to write logs into")
	rootCmd.PersistentFlags().UintVarP(&sleep, "sleep", "s", 0, "Enable no-cron execution. Sleep time in minutes")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	c, err := config.LoadConfig(".")
	if err != nil {
		log.Printf("cannot load config %s", err.Error())
	}

	loadedConfig = c
}
