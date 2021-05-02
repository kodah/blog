package cmd

import (
	"fmt"
	"os"

	"github.com/kodah/blog/server"

	"golang.org/x/crypto/bcrypt"

	"github.com/kodah/blog/dto"
	"github.com/kodah/blog/service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blogctl",
	Short: "blogctl is the cli for kodah's blog.",
	Long:  `blogctl is the cli for kodah's blog.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs the web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		var configService service.ConfigService = service.ConfigurationService("config.yaml")
		if configService.Error() != nil {
			fmt.Fprintf(os.Stderr, "Unable to load configuration file: %s", configService.Error())
			os.Exit(1)
		}

		var dbService service.DBService = service.SQLiteDBService(configService.GetDBPath())
		if dbService.Error() != nil {
			fmt.Fprintf(os.Stderr, "Unable to load configuration file: %s", configService.Error())
			os.Exit(1)
		}

		// migrate tables before connections are made
		err := dbService.Conn().AutoMigrate(dto.User{}, dto.Post{}, dto.Series{})
		if err != nil {
			return err
		}

		svr, err := server.NewWebServer()
		if err != nil {
			return err
		}

		return svr.Run(":8080")
	},
}

var mkpasswdCmd = &cobra.Command{
	Use:   "mkpasswd",
	Short: "mkpasswd creates a bcrypt hashed password.",
	Long:  `mkpasswd creates a bcrypt hashed password with a cost of 10.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := bcrypt.GenerateFromPassword([]byte(args[0]), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", data)

		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(mkpasswdCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
