package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cobra-demo",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取配置文件的值
		log.Println("app_env: ", viper.GetString("AppEnv"))
		log.Println("Cobra hello,version: v1.0.0")
	},
}
