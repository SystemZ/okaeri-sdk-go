package oksdk

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *NoProxyClient) Check(ctx context.Context, ip string) (result NoProxyResult, err error) {
	rawBody, httpResponse, err := c.Get(ctx, ip)
	if err != nil {
		return
	}
	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http error %d: %s", httpResponse.StatusCode, rawBody)
		return
	}
	err = json.Unmarshal(rawBody, &result)
	if err != nil {
		return
	}
	if c.LogEnabled {
		c.LogFunc(result)
	}
	return
}
