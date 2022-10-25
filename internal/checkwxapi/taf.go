package checkwxapi

import (
	"fmt"
	"os"
	"strings"
)

type getTafRes struct {
	Data    []string
	Results int
}

func (c *Client) RetrieveTafs(icaoCodes []string) map[string]string {
	var body getTafRes
	err := c.get("taf/", icaoCodes, &body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := make(map[string]string, len(icaoCodes))

	for _, code := range icaoCodes {
		prefix := fmt.Sprintf("TAF %s", code)
		prefixAmd := fmt.Sprintf("TAF AMD %s", code)

		for _, message := range body.Data {
			if strings.HasPrefix(message, prefix) || strings.HasPrefix(message, prefixAmd) {
				result[code] = message
			}
		}
	}

	return result
}
