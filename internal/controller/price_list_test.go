package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/warpgr/bova_test/internal/controller"
	"github.com/warpgr/bova_test/internal/controller/dto"
	"github.com/warpgr/bova_test/internal/controller/utils"
	"github.com/warpgr/bova_test/internal/repository"
	"github.com/warpgr/bova_test/internal/service"
	"github.com/warpgr/bova_test/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetPrices(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/ltp/list", nil)
	req.Header.Set("symbols", "BTC/USDT,ETH/USDT")

	c.Request = req

	st := store.NewKVMapStorage[string, float64](2)
	require.NoError(t, st.Store("BTCUSDT", 65000.2))

	ctr := controller.NewPriceList(
		service.NewPriceList(
			repository.NewPriceList(st)))

	ctr.GetPrices(c)

	require.Equal(t, http.StatusOK, w.Code)
	price := utils.Response[dto.PriceListDto]{}
	err := json.NewDecoder(w.Body).Decode(&price)
	require.NoError(t, err)

	require.Equal(t, "BTCUSDT", price.Data.Ltp[0].Pair)
	require.Equal(t, strconv.FormatFloat(65000.2, 'f', 2, 64), price.Data.Ltp[0].Amount)

}

func TestGetAllPrices(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/ltp/list", nil)
	req.Header.Set("symbols", "BTC/USDT,ETH/USDT")

	c.Request = req

	st := store.NewKVMapStorage[string, float64](2)
	require.NoError(t, st.Store("BTCUSDT", 65000.2))
	require.NoError(t, st.Store("ETHUSDT", 35000.2))

	ctr := controller.NewPriceList(
		service.NewPriceList(
			repository.NewPriceList(st)))

	ctr.GetAllPrices(c)

	require.Equal(t, http.StatusOK, w.Code)
	price := utils.Response[dto.PriceListDto]{}
	err := json.NewDecoder(w.Body).Decode(&price)
	require.NoError(t, err)

	for _, priceData := range price.Data.Ltp {
		if priceData.Pair == "BTCUSDT" {
			require.Equal(t, strconv.FormatFloat(65000.2, 'f', 2, 64), priceData.Amount)
		}
		if priceData.Pair == "ETHUSDT" {
			require.Equal(t, strconv.FormatFloat(35000.2, 'f', 2, 64), priceData.Amount)
		}
	}
}
