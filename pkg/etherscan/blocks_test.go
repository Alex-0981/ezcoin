package etherscan

import (
	"testing"
)

func TestGetBlockNumberByTimestamp(t *testing.T) {
	client := NewClient("R6KBAY2FYSIJS125S33ZTTW3S9C8EZHI1S")
	blockNumber, err := client.GetBlockNumberByTimestamp(1715702400)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blockNumber)
}
