package cmd

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rubyistdotjs/metarcli/internal/checkwxapi"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var apiKey string

	rootCmd := &cobra.Command{
		Use:   "metarcli icaoCodes...",
		Short: "Retrieve the latest METAR and TAF",
		Long:  "Retrieve the latest METAR and TAF messages for one or multiple airports",
		Args:  cobra.MatchAll(cobra.RangeArgs(1, 10), validateArgs),
		Run: func(cmd *cobra.Command, args []string) {
			cwxClient := checkwxapi.New(cmd.Context(), apiKey)
			icaoCodes := formatArgs(args)

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

	return rootCmd
}

func validateArgs(cmd *cobra.Command, args []string) error {
	for _, arg := range args {
		if len(arg) != 4 {
			return fmt.Errorf("invalid argument %q. Expected an ICAO airport code (4 letters)", arg)
		}
	}

	return nil
}

func formatArgs(args []string) []string {
	icaoCodes := make([]string, len(args))

	for i, arg := range args {
		icaoCodes[i] = strings.ToUpper(arg)
	}

	return icaoCodes
}
