/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
	"grpc-ldap-auth-cli/client"
	"grpc-ldap-auth-cli/prompts"
	"grpc-ldap-auth-cli/proto"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		yourName := prompts.PromptUser("What is your name?", nil)
		authToken := prompts.ConfirmSubmitAndGenerateAuth()
		grpcServiceClient := client.EstablishServerConnection()
		metaDataPairs := metadata.Pairs("authorization", fmt.Sprintf("Basic %s", authToken))
		ctx := metadata.NewOutgoingContext(context.Background(), metaDataPairs)
		response, err := grpcServiceClient.SayHi(ctx, &proto.SayHiRequest{MyName: yourName})
		if err != nil {
			logrus.Errorf("error occurred while submitting the request %v", err)
		} else {
			logrus.Info("response received")
			fmt.Println(response.GetGreetingResponse())
		}
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
