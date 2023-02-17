package migration

import (
	"gomigrate-test/migration/common"
	"gomigrate-test/migration/services"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	baseMigrationCmd.AddCommand(migrateCmd)
	baseMigrationCmd.AddCommand(rollbackCmd)
	baseMigrationCmd.AddCommand(codeGeneratorCmd)

	migrateCmd.PersistentFlags().StringP("run", "r", "", "run migration")

	rollbackCmd.PersistentFlags().StringP("run", "r", "", "run rollback migration")

	codeGeneratorCmd.PersistentFlags().StringP("name", "n", "", "name of migration")
	codeGeneratorCmd.PersistentFlags().StringP("filepath", "f", "", "init filepath")
}

func Execute() *cobra.Command {
	return baseMigrationCmd
}

var baseMigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration for gorm",
	Run: func(cmd *cobra.Command, args []string) {
		common.ShowBanner()
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run migration for gorm",
	Run: func(cmd *cobra.Command, args []string) {

		common.ShowBanner()

		runCommad, _ := cmd.Flags().GetString("run")

		if runCommad == "all" {
			OpenConnectionMysql()
			if err := services.Migration(db); err != nil {
				common.GenerateError(common.MsgMigrationError + err.Error())
				os.Exit(0)
			}

			common.GenerateInfo(common.MsgMigrationSuccess)
			os.Exit(0)
		} else if runCommad != "" {
			OpenConnectionMysql()
			if err := services.MigrateTo(db, runCommad); err != nil {
				common.GenerateError(common.MsgMigrationError + err.Error())
				os.Exit(0)
			}

			common.GenerateInfo(common.MsgMigrationSuccess)
			os.Exit(0)
		}

		common.GenerateWarning(common.MsgMigrationCommandNotFound)
	},
}

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "run rollback for gorm",
	Run: func(cmd *cobra.Command, args []string) {
		common.ShowBanner()
		runCommad, _ := cmd.Flags().GetString("run")
		if runCommad == "last" {
			OpenConnectionMysql()
			if err := services.RollbackLast(db); err != nil {
				common.GenerateError(common.MsgRollbackError + err.Error())
				os.Exit(0)
			}

			common.GenerateInfo(common.MsgRollbackSuccess)
			os.Exit(0)
		} else if runCommad != "" {
			OpenConnectionMysql()
			if err := services.RollbackTo(db, runCommad); err != nil {
				common.GenerateError(common.MsgRollbackError + err.Error())
				os.Exit(0)
			}

			common.GenerateInfo(common.MsgRollbackError)
			os.Exit(0)
		}

		common.GenerateWarning(common.MsgRollbackCommandNotFound)
	},
}

var codeGeneratorCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate migration blueprint for gorm",
	Run: func(cmd *cobra.Command, args []string) {
		common.ShowBanner()

		filepath, _ := cmd.Flags().GetString("filepath")
		if filepath == "" {
			common.GenerateError("Filepath must be filled!")
			os.Exit(0)
		}

		migrationName, _ := cmd.Flags().GetString("name")
		if migrationName == "" {
			common.GenerateError("migration name must be filled!")
			os.Exit(0)
		}

		services.PreparePayload(migrationName, filepath)
		services.EngineGeneratorProcess()
	},
}
