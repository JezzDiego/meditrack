package entities

import (
	"fmt"
	"meditrack/database"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type NCM struct {
	Code            string  `json:"code"`
	Description     string  `json:"description"`
	FullDescription string  `json:"full_description"`
	Ex              *string `json:"ex"`
}

func GetAllNCMs(c echo.Context) error {
	db := database.DBConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ncm")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	defer rows.Close()

	var ncmArr []NCM

	for rows.Next() {
		var ncm NCM

		if err := rows.Scan(&ncm.Code, &ncm.Description, &ncm.FullDescription, &ncm.Ex); err != nil {
			fmt.Println("Error scanning row:", err)
			return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		}

		ncmArr = append(ncmArr, ncm)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ncmArr)
}

func GetNCMById(c echo.Context) error {
	db := database.DBConn()
	defer db.Close()

	code := c.Param("code")

	rows := db.QueryRow("SELECT * FROM ncm WHERE code = ?", code)

	var ncm NCM

	if err := rows.Scan(&ncm.Code, &ncm.Description, &ncm.FullDescription, &ncm.Ex); err != nil {
		fmt.Println("Error scanning row:", err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ncm)
}

func PostNCM(c echo.Context) error {
	db := database.DBConn()
	defer db.Close()

	ncm := new(NCM)
	if err := c.Bind(ncm); err != nil {
		fmt.Println("Error binding NCM:", err)
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	_, err := db.Exec("INSERT INTO ncm (code, description, full_description, ex) VALUES (?, ?, ?, ?)", ncm.Code, ncm.Description, ncm.FullDescription, ncm.Ex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, ncm)
}

func PostNCMWithParams(ncm NCM) {
	db := database.DBConn()
	defer db.Close()

	_, err := db.Exec("INSERT INTO ncm (code, description, full_description, ex) VALUES (?, ?, ?, ?)", ncm.Code, ncm.Description, ncm.FullDescription, ncm.Ex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
	}
}
