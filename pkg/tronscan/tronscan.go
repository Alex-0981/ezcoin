package tronscan

import (
	"context"

	"go.uber.org/zap"
	"resty.dev/v3"
)

type Client struct {
	client *resty.Client
	apikey string
}

func NewClient(apikey string) *Client {
	return &Client{
		client: resty.New().
			SetBaseURL("https://apilist.tronscanapi.com/").
			SetQueryParam("apikey", apikey).
			SetDebug(true).
			SetDebugLogFormatter(resty.DebugLogJSONFormatter).
			SetTrace(true),
		apikey: apikey,
	}
}

type log struct {
	logger *zap.Logger
}

func (l *log) Debugf(format string, v ...interface{}) {
	// l.logger.Debug(format, v...)
}

func (l *log) Errorf(format string, v ...interface{}) {
	// l.logger.Debug(format, v...)
}

func (l *log) Warnf(format string, v ...interface{}) {
	// l.logger.Debug(format, v...)
}

func (l *log) Debug(v ...interface{}) {
	// l.logger.Debug(v...)
	// l.logger.Warn()
}

func NewClientWithCtx(ctx context.Context, apikey string) *Client {
	logger, _ := zap.NewDevelopment()
	logger = logger.With(zap.String("apikey", apikey))
	l := &log{
		logger: logger,
	}
	return &Client{
		client: resty.New().
			SetBaseURL("https://apilist.tronscanapi.com/").
			SetQueryParam("apikey", apikey).
			SetDebug(true).
			SetContext(ctx).
			SetLogger(l).
			SetDebugLogFormatter(resty.DebugLogJSONFormatter).
			SetTrace(true),
		apikey: apikey,
	}
}

func (c *Client) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
