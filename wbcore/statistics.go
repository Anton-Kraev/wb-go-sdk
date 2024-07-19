package wbcore

import (
	"errors"
	"fmt"

	"github.com/Anton-Kraev/wb-go-sdk/internal"
	"github.com/Anton-Kraev/wb-go-sdk/wbmodels"
)

// Statistics предоставляет интерфейс категории "Статистика" в WB API.
type Statistics interface {
	// WarehousesStocks реализует запрос https://openapi.wb.ru/statistics/api/ru/#tag/Statistika/paths/~1api~1v1~1supplier~1stocks/get.
	WarehousesStocks(wbmodels.WarehousesStocksRequest) (wbmodels.WarehousesStocksResponse, error)

	// SalesReport реализует запрос https://openapi.wb.ru/statistics/api/ru/#tag/Statistika/paths/~1api~1v5~1supplier~1reportDetailByPeriod/get.
	SalesReport(wbmodels.SalesReportRequest) wbmodels.SalesReportResponse
}

type statisticsImpl struct {
	token      string
	httpClient *internal.HttpClient
}

func NewStatistics(token string, httpClient *internal.HttpClient) Statistics {
	return statisticsImpl{token: token, httpClient: httpClient}
}

const (
	dateFromParam       = "dateFrom"
	warehousesStocksUrl = "https://statistics-api.wildberries.ru/api/v1/supplier/stocks"
)

var (
	errNotAuthorized   = errors.New("токен отсутствует, недействителен, удален или не применим к данному методу")
	errTooManyRequests = errors.New("превышено допустимое кол-во запросов в единицу времени")
	errWbServer        = errors.New("ошибка на стороне wildberries")
	errInternal        = errors.New("ошибка на нашей строне")
)

func (s statisticsImpl) WarehousesStocks(params wbmodels.WarehousesStocksRequest) (wbmodels.WarehousesStocksResponse, error) {
	req := s.httpClient.Request()

	var respModel wbmodels.WarehousesStocksResponse

	resp, err := req.
		SetQueryParam(dateFromParam, params.DateFrom).
		SetSuccessResult(&respModel).
		Get(warehousesStocksUrl)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errInternal, err)
	}

	if resp.IsSuccessState() {
		return respModel, nil
	}

	switch resp.StatusCode {
	case 401:
		return nil, errNotAuthorized
	case 429:
		return nil, errTooManyRequests
	case 500:
		return nil, errWbServer
	default:
		fmt.Println(resp.String())
		return nil, errInternal
	}
}

func (s statisticsImpl) SalesReport(params wbmodels.SalesReportRequest) wbmodels.SalesReportResponse {
	//TODO implement me
	panic("implement me")
}
