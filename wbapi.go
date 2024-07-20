package wbapi

import (
	"github.com/Anton-Kraev/wb-go-sdk/internal"
	"github.com/Anton-Kraev/wb-go-sdk/wbcore"
)

type WbApi interface {
	Content() wbcore.Content
	Statistics() wbcore.Statistics
}

type impl struct {
	content    wbcore.Content
	statistics wbcore.Statistics
}

func New() WbApi {
	reqClient := internal.NewReqClient()
	httpClient := internal.NewHttpClient(reqClient)

	return impl{
		content:    wbcore.NewContent(),
		statistics: wbcore.NewStatistics(httpClient),
	}
}

func (i impl) Content() wbcore.Content {
	return i.content
}

func (i impl) Statistics() wbcore.Statistics {
	return i.statistics
}
