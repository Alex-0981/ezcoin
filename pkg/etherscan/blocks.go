package etherscan

import (
	"fmt"
	"strconv"
)

func (c *Client) GetBlockNumberByTimestamp(timestamp int64) (int, error) {
	resp, err := c.client.R().SetQueryParams(map[string]string{
		"module":    "block",
		"action":    "getblocknobytime",
		"timestamp": strconv.FormatInt(timestamp, 10),
		"closest":   "before",
	}).Get("/v2/api")
	if err != nil {
		return 0, err
	}

	fmt.Println(resp.String())
	return 0, nil
}
