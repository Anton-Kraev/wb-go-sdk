package wbmodels

type WarehousesStocksRequest struct {
	DateFrom string
}

type WarehousesStocksResponse []struct {
	LastChangeDate  string  `json:"lastChangeDate"`
	WarehouseName   string  `json:"warehouseName"`
	SupplierArticle string  `json:"supplierArticle"`
	NmID            int     `json:"nmId"`
	Barcode         string  `json:"barcode"`
	Quantity        int     `json:"quantity"`
	InWayToClient   int     `json:"inWayToClient"`
	InWayFromClient int     `json:"inWayFromClient"`
	QuantityFull    int     `json:"quantityFull"`
	Category        string  `json:"category"`
	Subject         string  `json:"subject"`
	Brand           string  `json:"brand"`
	TechSize        string  `json:"techSize"`
	Price           float64 `json:"price"`
	Discount        float64 `json:"discount"`
	IsSupply        bool    `json:"isSupply"`
	IsRealization   bool    `json:"isRealization"`
	SCCode          string  `json:"scCode"`
}
