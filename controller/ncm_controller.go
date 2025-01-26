package controller

import (
	"meditrack/model"
	"meditrack/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NCMController struct {
	ncmUsecase usecase.NCMUsecase
}

func NewNCMController(ncmUsecase usecase.NCMUsecase) NCMController {
	return NCMController{
		ncmUsecase: ncmUsecase,
	}
}

// @ID				GetAllNCM
// @Tags			Nomenclatura Comum do Mercosul
// @Description	Retorna todos os NCMs cadastrados.
// @Produce		json
// @Success		200		{object}	model.NCM	"Requisição bem sucedida."
// @Failure		500		{string}	string		"Erro interno."
// @Router			/v1/ncm	[get]
func (n *NCMController) GetAllNCM(c echo.Context) error {
	ncm, err := n.ncmUsecase.GetAllNCM()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ncm)
}

// @ID				GetNCMByCode
// @Tags			Nomenclatura Comum do Mercosul
// @Description	Retorna um NCM com base no código informado.
// @Produce		json
// @Param			code			path		int			true	"Código do NCM"
// @Success		200				{object}	model.NCM	"Requisição bem sucedida."
// @Failure		404				{string}	string		"NCM não encontrado."
// @Failure		500				{string}	string		"Erro interno."
// @Router			/v1/ncm/{code}	[get]
func (n *NCMController) GetNCMByCode(c echo.Context) error {
	code := c.Param("code")

	if code == "" {
		return c.JSON(http.StatusBadRequest, "Code is required")
	}
	if _, err := strconv.Atoi(code); err != nil {
		return c.JSON(http.StatusBadRequest, "Code must be a number")
	}

	ncm, err := n.ncmUsecase.GetNCMByCode(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
	}

	if ncm == nil {
		return c.JSON(http.StatusNotFound, model.Response{Message: "NCM not found"})
	}

	return c.JSON(http.StatusOK, ncm)
}

// @ID				CreateNCM
// @Tags			Nomenclatura Comum do Mercosul
// @Description	Cria um novo NCM.
// @Accept			json
// @Produce		json
// @Param			ncm		body		model.NCM	true	"NCM a ser criado"
// @Success		201		{object}	model.NCM	"Requisição bem sucedida."
// @Failure		400		{string}	string		"Requisição inválida."
// @Failure		500		{string}	string		"Erro interno."
// @Router			/v1/ncm	[post]
func (n *NCMController) CreateNCM(c echo.Context) error {
	var ncm model.NCM
	if err := c.Bind(&ncm); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Invalid request"})
	}

	ncm, err := n.ncmUsecase.CreateNCM(ncm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, ncm)
}
