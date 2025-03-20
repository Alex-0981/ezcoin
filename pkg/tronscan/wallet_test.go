package tronscan

import (
	"testing"
	"time"
)

func TestGetTRC20TransferList(t *testing.T) {
	param := TRC20TransferListParam{
		Address:        "TZG5eG6ZJ8DTMjW19cZRujVjZQStNcgePi",
		Limit:          50,
		StartTimestamp: time.Now().AddDate(0, -1, 0).UnixMilli(),
		EndTimestamp:   time.Now().UnixMilli(),
	}

	client := NewClient("2964698c-5362-4b25-8ee4-36c842292583")
	resp, err := client.GetTRC20TransferList(param)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	t.Log(resp)
}
