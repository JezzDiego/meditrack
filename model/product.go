package model

type Product struct {
	ID           int      `json:"id"`
	Description  string   `json:"description"`
	Gtin         int64    `json:"gtin"`
	Width        *float32 `json:"width"`
	Height       *float32 `json:"height"`
	Length       *float32 `json:"length"`
	NetWeight    *float32 `json:"net_weight"`
	GrossWeight  *float32 `json:"gross_weight"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	ReleaseDate  *string  `json:"release_date"`
	Price        *string  `json:"price"`
	AvgPrice     float32  `json:"avg_price"`
	MinPrice     float32  `json:"min_price"`
	MaxPrice     float32  `json:"max_price"`
	Origin       *string  `json:"origin"`
	BarcodeImage *string  `json:"barcode_image"`
	NCMCode      string   `json:"ncm_code"`
	BrandName    string   `json:"brand_name"`
	BrandPicture *string  `json:"brand_picture"`
}

type FullProduct struct {
	Product
	NCM NCM `json:"ncm"`
}
