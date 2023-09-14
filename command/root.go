/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package command

import (
	"github.com/Appkube-awsx/awsx-cloudelements/command/vpccmd"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "getVpcList",
	Short: "getVpcList command gets vpc list",
	Long:  `getVpcList command gets vpc list of an aws account`,
	Run:   runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
	if err != nil {
		cmd.Help()
		return
	}
	if authFlag {
		if _, err := vpccmd.GetVpcList(*clientAuth); err != nil {
			cmd.Help()
			return
		}
	} else {
		cmd.Help()
		return
	}

}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	rootCmd.PersistentFlags().String("vaultToken", "", "vault token")
	rootCmd.PersistentFlags().String("accountId", "", "aws account number")
	rootCmd.PersistentFlags().String("zone", "", "aws region")
	rootCmd.PersistentFlags().String("accessKey", "", "aws access key")
	rootCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	rootCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	rootCmd.PersistentFlags().String("externalId", "", "aws external id")
}
