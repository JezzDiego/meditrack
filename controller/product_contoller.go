package controller

import (
	"fmt"
	"meditrack/model"
	"meditrack/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

//	@ID				GetAllProducts
//	@Tags			Produtos
//	@Description	Retorna todos os itens cadastrados.
//	@Produce		json
//	@Success		200			{object}	model.FullProduct	"Requisição bem sucedida."
//	@Failure		500			{string}	string				"Erro interno."
//	@Router			/v1/product	[get]
func (p *productController) GetAllProducts(c echo.Context) error {
	products, err := p.productUseCase.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}

//	@ID				GetProductById
//	@Tags			Produtos
//	@Description	Retorna um item com base no ID informado.
//	@Produce		json
//	@Param			id					path		int					true	"ID do item"
//	@Success		200					{object}	model.FullProduct	"Requisição bem sucedida."
//	@Failure		404					{string}	string				"Item não encontrado."
//	@Failure		500					{string}	string				"Erro interno."
//	@Router			/v1/product/{id} 	[get]
func (p *productController) GetProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Invalid ID"})
	}

	product, err := p.productUseCase.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, model.Response{Message: "Product not found"})
	}

	return c.JSON(http.StatusOK, product)
}

//	@ID				GetProductByGtin
//	@Tags			Produtos
//	@Description	Retorna um item com base no GTIN informado.
//	@Produce		json
//	@Param			gtin						path		string				true	"GTIN do item"
//	@Success		200							{object}	model.FullProduct	"Requisição bem sucedida."
//	@Failure		404							{string}	string				"Item não encontrado."
//	@Failure		500							{string}	string				"Erro interno."
//	@Router			/v1/product/gtin/{gtin} 	[get]
func (p *productController) GetProductByGtin(c echo.Context) error {
	gtin := c.Param("gtin")

	product, err := p.productUseCase.GetProductByGtin(gtin)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, model.Response{Message: "Product not found"})
	}

	return c.JSON(http.StatusOK, product)
}

//	@ID				CreateProduct
//	@Tags			Produtos
//	@Description	Cria um novo item.
//	@Accept			json
//	@Produce		json
//	@Param			product		body		model.FullProduct	true	"Item a ser criado."
//	@Success		201			{object}	model.FullProduct	"Item criado com sucesso."
//	@Failure		400			{string}	string				"Requisição inválida."
//	@Failure		500			{string}	string				"Erro interno."
//	@Router			/v1/product	  																											[post]
func (p *productController) CreateProduct(c echo.Context) error {
	var product model.FullProduct

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}

	createdProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, createdProduct)
}
