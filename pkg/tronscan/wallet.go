package tronscan

import (
	"errors"
	"strconv"
)

const (
	DIRECTION_BOTH = iota
	DIRECTION_IN
	DIRECTION_OUT
)

type TRC20TransferListParam struct {
	Start          int64
	Limit          int
	Address        string
	TRC20ID        string
	Direction      *int
	StartTimestamp int64
	EndTimestamp   int64
	DBVerison      int
	Reverse        *bool
}

type TRC20TransferListResponse struct {
	PageSize int                 `json:"page_size"`
	Code     int                 `json:"code"`
	Data     []TRC20TransferInfo `json:"data"`
}

type TRC20TransferInfo struct {
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	ApprovalAmount string `json:"approval_amount"`
	BlockTimestamp int64  `json:"block_timestamp"`
	Block          int64  `json:"block"`
	From           string `json:"from"`
	To             string `json:"to"`
	Hash           string `json:"hash"`
	Confirmed      int    `json:"confirmed"`
	Contract_Type  string `json:"contract_type"`
	ContractType   int    `json:"contractType"`
	Revert         int    `json:"revert"`
	ContractRet    string `json:"contract_ret"`
	EventType      string `json:"event_type"`
	IssueAddress   string `json:"issue_address"`
	Decimals       int    `json:"decimals"`
	TokenName      string `json:"token_name"`
	ID             string `json:"id"`
	Direction      int    `json:"direction"`
}

func (c *Client) GetTRC20TransferList(param TRC20TransferListParam) (*TRC20TransferListResponse, error) {
	direction := 1
	if param.Direction != nil {
		direction = *param.Direction
	}
	reverse := true
	if param.Reverse != nil {
		reverse = *param.Reverse
	}

	resp, err := c.client.R().SetQueryParams(map[string]string{
		"start":           strconv.FormatInt(param.Start, 10),
		"limit":           strconv.Itoa(param.Limit),
		"address":         param.Address,
		"trc20Id":         param.TRC20ID,
		"direction":       strconv.Itoa(direction),
		"start_timestamp": strconv.FormatInt(param.StartTimestamp, 10),
		"end_timestamp":   strconv.FormatInt(param.EndTimestamp, 10),
		"db_verison":      strconv.Itoa(param.DBVerison),
		"reverse":         strconv.FormatBool(reverse),
	}).SetResult(&TRC20TransferListResponse{}).Get("api/transfer/trxx")
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New(resp.String())
	}

	return resp.Result().(*TRC20TransferListResponse), nil
}
