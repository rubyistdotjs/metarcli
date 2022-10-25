package checkwxapi

import (
	"fmt"
	"os"
	"strings"
)

type getMetarRes struct {
	Data    []string
	Results int
}

func (c *Client) RetrieveMetars(icaoCodes []string) map[string]string {
	var body getMetarRes
	err := c.get("metar/", icaoCodes, &body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := make(map[string]string, len(icaoCodes))

	for _, code := range icaoCodes {
		for _, message := range body.Data {
			if strings.HasPrefix(message, code) {
				result[code] = message
			}
		}
	}

	return result
}
