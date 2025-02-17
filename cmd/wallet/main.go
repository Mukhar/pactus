package main

import (
	"fmt"

	"github.com/pactus-project/pactus/cmd"
	"github.com/pactus-project/pactus/wallet"
	"github.com/spf13/cobra"
)

var (
	pathOpt       *string
	offlineOpt    *bool
	serverAddrOpt *string
)

func addPasswordOption(c *cobra.Command) *string {
	return c.Flags().StringP("password", "p",
		"", "the wallet password")
}

func openWallet() (*wallet.Wallet, error) {
	if !*offlineOpt && *serverAddrOpt != "" {
		wallet, err := wallet.Open(*pathOpt, true)
		if err != nil {
			return nil, err
		}

		err = wallet.Connect(*serverAddrOpt)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return wallet, err
	}
	wallet, err := wallet.Open(*pathOpt, *offlineOpt)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:               "pactus-wallet",
		Short:             "Pactus wallet",
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	// Hide the "help" sub-command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	pathOpt = rootCmd.PersistentFlags().String("path", "default_wallet", "the path to the wallet file")
	offlineOpt = rootCmd.PersistentFlags().Bool("offline", false, "offline mode")
	serverAddrOpt = rootCmd.PersistentFlags().String("server", "", "server gRPC address")

	buildCreateCmd(rootCmd)
	buildRecoverCmd(rootCmd)
	buildGetSeedCmd(rootCmd)
	buildChangePasswordCmd(rootCmd)
	buildAllTransactionCmd(rootCmd)
	buildAllAddrCmd(rootCmd)
	buildAllHistoryCmd(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		cmd.PrintErrorMsgf("%s", err)
	}
}
