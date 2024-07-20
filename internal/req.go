package internal

import (
	"time"

	"github.com/imroc/req/v3"
)

const reqTimeout = 5 * time.Second

func NewReqClient() *req.Client {
	return req.C().SetTimeout(reqTimeout)
}
