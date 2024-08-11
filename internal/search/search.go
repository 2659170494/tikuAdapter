package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/model"
)

// Search 搜题接口
type Search interface {
	getHTTPClient() *resty.Client
	SearchAnswer(req model.SearchRequest) (answer [][]string, err error)
}
