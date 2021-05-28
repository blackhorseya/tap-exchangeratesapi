package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/entity/config"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	BaseUrl = "http://api.exchangeratesapi.io/v1"
)

var cfgFile string
var stateFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tap-exchangeratesapi",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		c := new(config.Config)
		err := viper.Unmarshal(c)
		cobra.CheckErr(err)

		uri := fmt.Sprintf("%s/%s", BaseUrl, c.StartDate)
		params := url.Values{}
		if len(c.ApiKey) != 0 {
			params.Add("access_key", c.ApiKey)
		}
		if len(c.Base) != 0 {
			params.Add("base", c.Base)
		}

		resp, err := request(uri, params.Encode())
		cobra.CheckErr(err)
		defer resp.Body.Close()

		payload, err := ioutil.ReadAll(resp.Body)
		cobra.CheckErr(err)
		fmt.Println(string(payload))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./configs/config.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&stateFile, "state", "s", "", "state file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".tap-exchangeratesapi" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tap-exchangeratesapi")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func request(url string, params string) (*http.Response, error) {
	uri := url + "?" + params
	resp, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func do(base string, startDate time.Time) {
	// todo: 2021-05-29|03:05|doggy|implement me
	panic("implement me")
}
