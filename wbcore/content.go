package wbcore

import "github.com/Anton-Kraev/wb-go-sdk/wbmodels"

// Content предоставляет интерфейс категории "Контент" в WB API.
type Content interface {
	// ProductCardLimits реализует запрос https://openapi.wb.ru/content/api/ru/#tag/Prosmotr/paths/~1content~1v2~1cards~1limits/get.
	ProductCardLimits(wbmodels.ProductCardLimitsRequest) wbmodels.ProductCardLimitsResponse
}

type contentImpl struct {
	token string
}

func NewContent(token string) Content {
	return contentImpl{token: token}
}

func (c contentImpl) ProductCardLimits(req wbmodels.ProductCardLimitsRequest) wbmodels.ProductCardLimitsResponse {
	//TODO implement me
	panic("implement me")
}
