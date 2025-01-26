package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"meditrack/handlers"
	"meditrack/model"
	"net/http"
)

type ProductRepository struct {
	connection *sql.DB
	outerAPI   *handlers.OuterAPIHandler
}

func NewProductRepository(db *sql.DB, outerAPI *handlers.OuterAPIHandler) ProductRepository {
	return ProductRepository{
		connection: db,
		outerAPI:   outerAPI,
	}
}

func (p *ProductRepository) GetAllProducts() ([]model.FullProduct, error) {
	query := `SELECT * FROM medicine INNER JOIN ncm ON medicine.ncm_code = ncm.code`

	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.FullProduct{}, err
	}
	defer rows.Close()

	var productList []model.FullProduct

	for rows.Next() {
		var product model.FullProduct
		var ncm model.NCM

		if err := rows.Scan(
			&product.ID,
			&product.Description,
			&product.Gtin,
			&product.Width,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.GrossWeight,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.ReleaseDate,
			&product.Price,
			&product.AvgPrice,
			&product.MinPrice,
			&product.MaxPrice,
			&product.Origin,
			&product.BarcodeImage,
			&product.NCMCode,
			&product.BrandName,
			&product.BrandPicture,

			// NCM
			&ncm.Code,
			&ncm.Description,
			&ncm.FullDescription,
			&ncm.Ex,
		); err != nil {
			fmt.Println("Item not found:", err)
			return []model.FullProduct{}, err
		}

		product.NCM = ncm

		productList = append(productList, product)
	}

	return productList, nil
}

func (p *ProductRepository) GetProductById(id int) (*model.FullProduct, error) {
	query := `SELECT * FROM medicine INNER JOIN ncm ON medicine.ncm_code = ncm.code WHERE medicine.id = ?`

	row := p.connection.QueryRow(query, id)

	var product model.FullProduct
	var ncm model.NCM

	if err := row.Scan(
		&product.ID,
		&product.Description,
		&product.Gtin,
		&product.Width,
		&product.Height,
		&product.Length,
		&product.NetWeight,
		&product.GrossWeight,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.ReleaseDate,
		&product.Price,
		&product.AvgPrice,
		&product.MinPrice,
		&product.MaxPrice,
		&product.Origin,
		&product.BarcodeImage,
		&product.NCMCode,
		&product.BrandName,
		&product.BrandPicture,

		// NCM
		&ncm.Code,
		&ncm.Description,
		&ncm.FullDescription,
		&ncm.Ex,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	product.NCM = ncm

	return &product, nil
}

func (p *ProductRepository) GetProductByGtin(gtin string) (*model.FullProduct, error) {
	query := `SELECT * FROM medicine INNER JOIN ncm ON medicine.ncm_code = ncm.code WHERE medicine.gtin = ?`

	row := p.connection.QueryRow(query, gtin)

	var product model.FullProduct
	var ncm model.NCM

	if err := row.Scan(
		&product.ID,
		&product.Description,
		&product.Gtin,
		&product.Width,
		&product.Height,
		&product.Length,
		&product.NetWeight,
		&product.GrossWeight,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.ReleaseDate,
		&product.Price,
		&product.AvgPrice,
		&product.MinPrice,
		&product.MaxPrice,
		&product.Origin,
		&product.BarcodeImage,
		&product.NCMCode,
		&product.BrandName,
		&product.BrandPicture,

		// NCM
		&ncm.Code,
		&ncm.Description,
		&ncm.FullDescription,
		&ncm.Ex,
	); err != nil {
		if err == sql.ErrNoRows {
			client := &http.Client{}
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/gtins/%s.json", p.outerAPI.OuterAPIURL, gtin), nil)
			if err != nil {
				return nil, err
			}

			req.Header.Add(p.outerAPI.OuterAPIAuthHeader, p.outerAPI.OuterAPIToken)
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("User-Agent", "Cosmos-API-Request")

			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}

			var outerAPIResponse model.FullProduct

			if err := json.NewDecoder(resp.Body).Decode(&outerAPIResponse); err != nil {
				return nil, err
			}

			ncmQuery := `INSERT INTO ncm (code, description, full_description, ex) VALUES (?, ?, ?, ?)`

			_, err = p.connection.Exec(ncmQuery,
				outerAPIResponse.NCM.Code,
				outerAPIResponse.NCM.Description,
				outerAPIResponse.NCM.FullDescription,
				outerAPIResponse.NCM.Ex,
			)
			if err != nil {
				fmt.Println(err)
			}

			productQuery := `
				INSERT INTO medicine
				(description, gtin, width, height, length, net_weight, gross_weight,
				release_date, price, avg_price, min_price, max_price, origin, barcode_image, ncm_code, brand_name, brand_picture)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

			result, err := p.connection.Exec(productQuery,
				outerAPIResponse.Description,
				outerAPIResponse.Gtin,
				outerAPIResponse.Width,
				outerAPIResponse.Height,
				outerAPIResponse.Length,
				outerAPIResponse.NetWeight,
				outerAPIResponse.GrossWeight,
				outerAPIResponse.ReleaseDate,
				outerAPIResponse.Price,
				outerAPIResponse.AvgPrice,
				outerAPIResponse.MinPrice,
				outerAPIResponse.MaxPrice,
				outerAPIResponse.Origin,
				outerAPIResponse.BarcodeImage,
				outerAPIResponse.NCM.Code,
				outerAPIResponse.BrandName,
				outerAPIResponse.BrandPicture,

				// NCM
				outerAPIResponse.NCM.Code,
				outerAPIResponse.NCM.Description,
				outerAPIResponse.NCM.FullDescription,
				outerAPIResponse.NCM.Ex,
			)
			if err != nil {
				return nil, err
			}

			id, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}

			product = outerAPIResponse
			product.ID = int(id)

			return &product, nil
		}

		return nil, err
	}

	product.NCM = ncm

	return &product, nil
}

func (p *ProductRepository) CreateProduct(product model.FullProduct) (int, error) {
	productQuery := `
	INSERT INTO medicine
	(description, gtin, width, height, length, net_weight, gross_weight,
	release_date, price, avg_price, min_price, max_price, origin, barcode_image, ncm_code, brand_name, brand_picture)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	ncmQuery := `INSERT INTO ncm (code, description, full_description, ex) VALUES (?, ?, ?, ?)`

	_, err := p.connection.Exec(ncmQuery,
		product.NCM.Code,
		product.NCM.Description,
		product.NCM.FullDescription,
		product.NCM.Ex,
	)
	if err != nil {
		fmt.Println(err)
	}

	result, err := p.connection.Exec(productQuery,
		product.Description,
		product.Gtin,
		product.Width,
		product.Height,
		product.Length,
		product.NetWeight,
		product.GrossWeight,
		product.ReleaseDate,
		product.Price,
		product.AvgPrice,
		product.MinPrice,
		product.MaxPrice,
		product.Origin,
		product.BarcodeImage,
		product.NCMCode,
		product.BrandName,
		product.BrandPicture,
	)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return int(id), nil
}
