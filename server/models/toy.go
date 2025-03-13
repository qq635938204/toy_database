package models

type JsonCodeMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JsonToyInfo struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`        // 价格
	Stock       int64   `json:"stock"`        // 库存
	ImagePath   string  `json:"image_path"`   // 图片路径
	Description string  `json:"description"`  // 描述
	MBNum       string  `json:"mb_num"`       // Matchbox编号
	BuyDate     string  `json:"buy_date"`     // 购买日期
	Brand       string  `json:"brand"`        // 品牌
	ModelName   string  `json:"model_name"`   // 型号
	YearOnBase  string  `json:"year_on_base"` // 底座年份
}

type JsonToyInfoResp struct {
	JsonCodeMsg
	Data []JsonToyInfo `json:"data"`
}
