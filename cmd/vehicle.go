/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go-cli-crud/internal"
)

// vehicleCmd represents the vehicle command
var (
	v internal.Vehicle

	vehicleCmd = &cobra.Command{
		Use:   "vehicle",
		Short: "Create a new vehicle for a company",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%#v\n", &v)
		},
	}
)

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vehicleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	vehicleCmd.Flags().StringVarP(&v.Name, "name", "n", "", "Sets vehicle model name")
	vehicleCmd.Flags().IntVarP(&v.Price, "price", "p", 0, "Sets vehicle model price")
	vehicleCmd.Flags().IntVarP(&v.Seats, "seats", "s", 2, "Sets vehicle model seats")
	vehicleCmd.Flags().StringVarP(&v.Type, "gear", "g", "", "Sets vehicle model gear")
	vehicleCmd.Flags().IntVarP(&v.MaxSpeed, "maxspeed", "m", 0, "Sets vehicle model max speed")

	vehicleCmd.MarkFlagRequired("name")
	vehicleCmd.MarkFlagRequired("price")
}
