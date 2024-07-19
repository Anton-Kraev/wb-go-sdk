package wbcore

import "github.com/Anton-Kraev/wb-go-sdk/wbmodels"

// Statistics предоставляет интерфейс категории "Статистика" в WB API.
type Statistics interface {
	// WarehousesStocks реализует запрос https://openapi.wb.ru/statistics/api/ru/#tag/Statistika/paths/~1api~1v1~1supplier~1stocks/get.
	WarehousesStocks(wbmodels.WarehousesStocksRequest) wbmodels.WarehousesStocksResponse

	// SalesReport реализует запрос https://openapi.wb.ru/statistics/api/ru/#tag/Statistika/paths/~1api~1v5~1supplier~1reportDetailByPeriod/get.
	SalesReport(wbmodels.SalesReportRequest) wbmodels.SalesReportResponse
}

type statisticsImpl struct {
	token string
}

func NewStatistics(token string) Statistics {
	return statisticsImpl{token: token}
}

func (s statisticsImpl) WarehousesStocks(req wbmodels.WarehousesStocksRequest) wbmodels.WarehousesStocksResponse {
	//TODO implement me
	panic("implement me")
}

func (s statisticsImpl) SalesReport(req wbmodels.SalesReportRequest) wbmodels.SalesReportResponse {
	//TODO implement me
	panic("implement me")
}
