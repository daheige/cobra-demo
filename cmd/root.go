package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-demo",
	Short: "cobra-demo",
	Long:  `do some task or job`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("please run available commands")
	},
}

var (
	cfgFile string
)

// InitConfig 初始化配置
func InitConfig() {
	// 指定初始化函数
	cobra.OnInitialize(initConfig)

	// 初始化全局参数
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./app.yaml)")
	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := filepath.Abs("./")
		cobra.CheckErr(err)

		log.Println("config_dir: ", configDir)
		// Search config in ./ directory with name "app.yaml"
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("app")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute 执行函数
func Execute() {
	defer Shutdown()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Shutdown 平滑退出处理
func Shutdown() {
	log.Println("server shutdown success")
}
