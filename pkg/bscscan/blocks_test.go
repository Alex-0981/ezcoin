package bscscan

import (
	"testing"
)

func TestGetBlockNumberByTimestamp(t *testing.T) {
	client := NewClient("QN28DPGEKFHNJ1JGMHKV42S2Q9SIQQ1NV8")
	blockNumber, err := client.GetBlockNumberByTimestamp(1715702400)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blockNumber)
}
