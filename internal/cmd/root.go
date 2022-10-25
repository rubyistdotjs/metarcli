package cmd

import (
	"fmt"
	"sync"

	"github.com/rubyistdotjs/metarcli/internal/checkwxapi"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var apiKey string
	var icaoCodes []string

	rootCmd := &cobra.Command{
		Use:   "metar",
		Short: "TBD",
		Long:  "TBD",
		Run: func(cmd *cobra.Command, args []string) {
			cwxClient := checkwxapi.New(cmd.Context(), apiKey)

			wg := sync.WaitGroup{}
			wg.Add(3)

			var stations map[string]checkwxapi.Station
			var metars map[string]string
			var tafs map[string]string

			go func() {
				stations = cwxClient.RetrieveStations(icaoCodes)
				wg.Done()
			}()

			go func() {
				metars = cwxClient.RetrieveMetars(icaoCodes)
				wg.Done()
			}()

			go func() {
				tafs = cwxClient.RetrieveTafs(icaoCodes)
				wg.Done()
			}()

			wg.Wait()

			fmt.Printf("\n")

			for _, icaoCode := range icaoCodes {
				if station, found := stations[icaoCode]; found {
					fmt.Println(station.Name)
				} else {
					fmt.Println(icaoCode)
				}

				if metar, found := metars[icaoCode]; found {
					fmt.Println(metar)
				} else {
					fmt.Println("-")
				}

				if taf, found := tafs[icaoCode]; found {
					fmt.Println(taf)
				} else {
					fmt.Println("-")
				}

				fmt.Printf("\n")
			}
		},
	}

	rootCmd.Flags().StringVarP(&apiKey, "apiKey", "k", "", "CheckWX API key")
	rootCmd.MarkFlagRequired("apiKey")

	rootCmd.Flags().StringSliceVarP(&icaoCodes, "icaoCodes", "c", []string{""}, "ICAO airport codes of which you wish to view the weather information")
	rootCmd.MarkFlagRequired("icaoCodes")

	return rootCmd
}
