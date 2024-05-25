package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/warpgr/bova_test/internal/controller/dto"
	"github.com/warpgr/bova_test/internal/controller/utils"
	"github.com/warpgr/bova_test/internal/service"

	"github.com/gin-gonic/gin"
)

type PriceList interface {
	// Register registers own handlers on provided router.
	Register(e *gin.RouterGroup)

	// GetPrices handles /list route request.
	GetPrices(ctx *gin.Context)

	// GetAllPrices handles / route request.
	GetAllPrices(ctx *gin.Context)
}

func NewPriceList(svc service.PriceList) PriceList {
	return &priceList{
		svc: svc,
	}
}

type priceList struct {
	svc service.PriceList
}

func (c *priceList) Register(e *gin.RouterGroup) {
	routes := e.Group("/ltp")
	routes.GET("/list", c.GetPrices)
	routes.GET("/", c.GetAllPrices)
}

func (c *priceList) GetPrices(ctx *gin.Context) {
	symbolsStr := ctx.Request.Header.Get("symbols")
	symbols := strings.Split(symbolsStr, ",")

	for _, symbol := range symbols {
		if !utils.ValidateSymbol(symbol) {
			ctx.JSON(http.StatusBadRequest,
				utils.BuildErrorResponse(errors.New("invalid symbol received"), 1))
			return
		}
	}

	priceList, err := c.svc.GetPrices(symbols)
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			utils.BuildErrorResponse(err, 2))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildOkResponse(dto.BuildPriceListDto(priceList), ""))
}

func (c *priceList) GetAllPrices(ctx *gin.Context) {
	priceList, err := c.svc.GetAllPrices()
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			utils.BuildErrorResponse(err, 2))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildOkResponse(dto.BuildPriceListDto(priceList), ""))
}
