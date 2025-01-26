package entities

import (
	"encoding/json"
	"fmt"
	"meditrack/database"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Medicine struct {
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

type FullMedicine struct {
	Medicine
	NCM NCM `json:"ncm"`
}

//	@ID				GetAllMedicines
//	@Tags			v1
//	@Description	Retorna todos os itens.
//	@Produce		json
//	@Success		200			{object}	FullMedicine[]	"Requisição bem sucedida."
//	@Failure		404			{string}	string			"Item não encontrado."
//	@Failure		500			{string}	string			"Erro interno."
//	@Router			/medicine 	[get]
func GetAllMedicines(c echo.Context) error {
	db := database.DBConn()
	defer db.Close()

	query := `SELECT * FROM medicine INNER JOIN ncm ON medicine.ncm_code = ncm.code`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	defer rows.Close()

	var medicinesArr []FullMedicine

	for rows.Next() {
		var medicine FullMedicine
		var ncm NCM

		if err := rows.Scan(
			&medicine.ID,
			&medicine.Description,
			&medicine.Gtin,
			&medicine.Width,
			&medicine.Height,
			&medicine.Length,
			&medicine.NetWeight,
			&medicine.GrossWeight,
			&medicine.CreatedAt,
			&medicine.UpdatedAt,
			&medicine.ReleaseDate,
			&medicine.Price,
			&medicine.AvgPrice,
			&medicine.MinPrice,
			&medicine.MaxPrice,
			&medicine.Origin,
			&medicine.BarcodeImage,
			&medicine.NCMCode,
			&medicine.BrandName,
			&medicine.BrandPicture,

			// NCM
			&ncm.Code,
			&ncm.Description,
			&ncm.FullDescription,
			&ncm.Ex,
		); err != nil {
			fmt.Println("Item not found:", err)
			return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
		}

		medicine.NCM = ncm

		medicinesArr = append(medicinesArr, medicine)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, medicinesArr)
}

//	@ID				GetMedicineById
//	@Tags			v1
//	@Description	Retorna um item com base no id informado.
//	@Produce		json
//	@Success		200				{object}	FullMedicine	"Requisição bem sucedida."
//	@Failure		404				{string}	string			"Item não encontrado."
//	@Failure		500				{string}	string			"Erro interno."
//	@Router			/medicine/{id} 	[get]
func GetMedicineById(c echo.Context) error {
	db := database.DBConn()
	defer db.Close()

	id := c.Param("id")

	query := `SELECT * FROM medicine INNER JOIN ncm ON medicine.ncm_code = ncm.code WHERE medicine.id = ?`

	rows := db.QueryRow(query, id)

	var medicine FullMedicine
	var ncm NCM

	if err := rows.Scan(
		&medicine.ID,
		&medicine.Description,
		&medicine.Gtin,
		&medicine.Width,
		&medicine.Height,
		&medicine.Length,
		&medicine.NetWeight,
		&medicine.GrossWeight,
		&medicine.CreatedAt,
		&medicine.UpdatedAt,
		&medicine.ReleaseDate,
		&medicine.Price,
		&medicine.AvgPrice,
		&medicine.MinPrice,
		&medicine.MaxPrice,
		&medicine.Origin,
		&medicine.BarcodeImage,
		&medicine.NCMCode,
		&medicine.BrandName,
		&medicine.BrandPicture,

		// NCM
		&ncm.Code,
		&ncm.Description,
		&ncm.FullDescription,
		&ncm.Ex,
	); err != nil {
		fmt.Println("Item not found:", err)
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	medicine.NCM = ncm

	return c.JSON(http.StatusOK, medicine)
}

//	@ID				GetMedicineByGtin
//	@Tags			v1
//	@Description	Retorna um item com base no gtin informado.
//	@Produce		json
//	@Success		200						{object}	FullMedicine	"Requisição bem sucedida."
//	@Failure		404						{string}	string			"Item não encontrado."
//	@Failure		500						{string}	string			"Erro interno."
//	@Router			/medicine/gtin/{gtin} 	[get]
func (h *OuterAPIHandler) GetMedicineByGtin(c echo.Context) error {
	db := database.DBConn()
	defer db.Close()

	gtin := c.Param("gtin")

	query := `SELECT * FROM medicine INNER JOIN ncm ON medicine.ncm_code = ncm.code WHERE medicine.gtin = ?`

	row := db.QueryRow(query, gtin)

	var medicine FullMedicine
	var ncm NCM

	if err := row.Scan(
		&medicine.ID,
		&medicine.Description,
		&medicine.Gtin,
		&medicine.Width,
		&medicine.Height,
		&medicine.Length,
		&medicine.NetWeight,
		&medicine.GrossWeight,
		&medicine.CreatedAt,
		&medicine.UpdatedAt,
		&medicine.ReleaseDate,
		&medicine.Price,
		&medicine.AvgPrice,
		&medicine.MinPrice,
		&medicine.MaxPrice,
		&medicine.Origin,
		&medicine.BarcodeImage,
		&medicine.NCMCode,
		&medicine.BrandName,
		&medicine.BrandPicture,

		// NCM
		&ncm.Code,
		&ncm.Description,
		&ncm.FullDescription,
		&ncm.Ex,
	); err != nil {
		fmt.Println("medicine with gtin", gtin, "not found in database:", err)

		client := &http.Client{}

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/gtins/%s.json", h.OuterAPIURL, gtin), nil)

		if err != nil {
			fmt.Println("failed to set up the request:", err)
			return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		}

		req.Header.Add(h.OuterAPIAuthHeader, h.OuterAPIToken)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "Cosmos-API-Request")

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("item not found in outer api:", err)
			return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
		}

		var outerAPIResponse FullMedicine

		if err := json.NewDecoder(resp.Body).Decode(&outerAPIResponse); err != nil {
			fmt.Println("failed to decode response:", err)
			return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		}

		PostNCMWithParams(outerAPIResponse.NCM)

		PostMedicineWithParams(Medicine{
			Description:  outerAPIResponse.Description,
			Gtin:         outerAPIResponse.Gtin,
			Width:        outerAPIResponse.Width,
			Height:       outerAPIResponse.Height,
			Length:       outerAPIResponse.Length,
			NetWeight:    outerAPIResponse.NetWeight,
			GrossWeight:  outerAPIResponse.GrossWeight,
			ReleaseDate:  outerAPIResponse.ReleaseDate,
			Price:        outerAPIResponse.Price,
			AvgPrice:     outerAPIResponse.AvgPrice,
			MinPrice:     outerAPIResponse.MinPrice,
			MaxPrice:     outerAPIResponse.MaxPrice,
			Origin:       outerAPIResponse.Origin,
			BarcodeImage: outerAPIResponse.BarcodeImage,
			NCMCode:      outerAPIResponse.NCM.Code,
			BrandName:    outerAPIResponse.BrandName,
			BrandPicture: outerAPIResponse.BrandPicture,
		})

		outerAPIResponse.ID = -1

		return c.JSON(http.StatusOK, outerAPIResponse)
	}

	medicine.NCM = ncm

	return c.JSON(http.StatusOK, medicine)
}

func PostMedicine(c echo.Context) error {
	db := database.DBConn()
	defer db.Close()

	medicine := new(Medicine)
	if err := c.Bind(medicine); err != nil {
		fmt.Fprintf(os.Stderr, "failed to bind: %v\n", err)

		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	query := `
	INSERT INTO medicine
	(description, gtin, width, height, length, net_weight, gross_weight,
	release_date, price, avg_price, min_price, max_price, origin, barcode_image, ncm_code, brand_name, brand_picture)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		medicine.Description,
		medicine.Gtin,
		medicine.Width,
		medicine.Height,
		medicine.Length,
		medicine.NetWeight,
		medicine.GrossWeight,
		medicine.ReleaseDate,
		medicine.Price,
		medicine.AvgPrice,
		medicine.MinPrice,
		medicine.MaxPrice,
		medicine.Origin,
		medicine.BarcodeImage,
		medicine.NCMCode,
		medicine.BrandName,
		medicine.BrandPicture,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, medicine)
}

func PostMedicineWithParams(medicine Medicine) {
	db := database.DBConn()
	defer db.Close()

	query := `
	INSERT INTO medicine
	(description, gtin, width, height, length, net_weight, gross_weight,
	release_date, price, avg_price, min_price, max_price, origin, barcode_image, ncm_code, brand_name, brand_picture)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		medicine.Description,
		medicine.Gtin,
		medicine.Width,
		medicine.Height,
		medicine.Length,
		medicine.NetWeight,
		medicine.GrossWeight,
		medicine.ReleaseDate,
		medicine.Price,
		medicine.AvgPrice,
		medicine.MinPrice,
		medicine.MaxPrice,
		medicine.Origin,
		medicine.BarcodeImage,
		medicine.NCMCode,
		medicine.BrandName,
		medicine.BrandPicture,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
	}

}
