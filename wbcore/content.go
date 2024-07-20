package wbcore

import "github.com/Anton-Kraev/wb-go-sdk/wbmodels"

// Content предоставляет интерфейс категории "Контент" в WB API.
type Content interface {
	// ProductCardLimits реализует запрос https://openapi.wb.ru/content/api/ru/#tag/Prosmotr/paths/~1content~1v2~1cards~1limits/get.
	ProductCardLimits(token string, req wbmodels.ProductCardLimitsRequest) wbmodels.ProductCardLimitsResponse
}

type contentImpl struct {
}

func NewContent() Content {
	return contentImpl{}
}

func (c contentImpl) ProductCardLimits(token string, req wbmodels.ProductCardLimitsRequest) wbmodels.ProductCardLimitsResponse {
	//TODO implement me
	panic("implement me")
}
