package cmd

import (
	"github.com/latifrons/commongo/program"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "net/http/pprof"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer program.DumpStack(true)
	_ = rootCmd.Execute()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ScriptCarrier",
	Short: "ScriptCarrier",
	Long:  `ScriptCarrier to da moon`,
}

func init() {
	rootCmd.PersistentFlags().StringP("dir-root", "r", "nodedata", "Folder for all data of one core")
	rootCmd.PersistentFlags().String("dir-log", "", "Log folder. Default to {dir.root}/log")
	rootCmd.PersistentFlags().String("dir-data", "", "Data folder. Default to {dir.root}/data")
	rootCmd.PersistentFlags().String("dir-config", "", "Config folder. Default to {dir.root}/config")
	rootCmd.PersistentFlags().String("dir-private", "", "Private folder. Default to {dir.root}/private")

	rootCmd.PersistentFlags().String("log-level", "info", "Logging verbosity, possible values:[panic, fatal, berror, warn, info, debug]")
	rootCmd.PersistentFlags().Bool("debug-return_detailed_error", false, "In Rpc response, return detailed berror message for debugging.")

	_ = viper.BindPFlag("dir.root", rootCmd.PersistentFlags().Lookup("dir-root"))
	_ = viper.BindPFlag("dir.log", rootCmd.PersistentFlags().Lookup("dir-log"))
	_ = viper.BindPFlag("dir.data", rootCmd.PersistentFlags().Lookup("dir-data"))
	_ = viper.BindPFlag("dir.config", rootCmd.PersistentFlags().Lookup("dir-config"))
	_ = viper.BindPFlag("dir.private", rootCmd.PersistentFlags().Lookup("dir-private"))
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	_ = viper.BindPFlag("debug.return_detailed_error", rootCmd.PersistentFlags().Lookup("debug-return_detailed_error"))

}
