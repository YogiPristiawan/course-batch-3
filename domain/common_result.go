package domain

import "log"

type CommonResult struct {
	ResErrorCode    int    `json:"-"`
	ResErrorMessage string `json:"-"`
}

func (c *CommonResult) SetError(code int, message string) {
	c.ResErrorCode = code
	c.ResErrorMessage = message

	if code >= 500 {
		// send to slack, etc
		log.Printf("%d: %s", code, message)
	}
}
