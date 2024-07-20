package wbapi

import (
	"github.com/Anton-Kraev/wb-go-sdk/internal"
	"github.com/Anton-Kraev/wb-go-sdk/wbcore"
)

type WbApi interface {
	Content() wbcore.Content
	Statistics() wbcore.Statistics
}

type Impl struct {
	token      string
	content    wbcore.Content
	statistics wbcore.Statistics
}

func New(token string) WbApi {
	reqClient := internal.NewReqClient(token)
	httpClient := internal.NewHttpClient(reqClient)

	return Impl{
		token:      token,
		content:    wbcore.NewContent(token),
		statistics: wbcore.NewStatistics(token, httpClient),
	}
}

func (i Impl) Content() wbcore.Content {
	return i.content
}

func (i Impl) Statistics() wbcore.Statistics {
	return i.statistics
}
