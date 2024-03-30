package coin

import (
	"bytes"
	coinResponse "coins/internal/delivery/http/responses/coin"
	domain "coins/internal/domain/coin"
	"coins/internal/domain/coin/types/code"
	"coins/internal/domain/coin/types/name"
	mockUseCase "coins/internal/useCases/mock"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
)

var (
	recorder *httptest.ResponseRecorder
	ctx      *gin.Context
	request  *http.Request
	handler  *Handler
	uc       = new(mockUseCase.Coin)
	mainCoin *domain.Coin
)

func TestMain(m *testing.M) {
	resetData()

	handler = New(uc)

	os.Exit(m.Run())
}

func resetData() {
	recorder = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(recorder)

	request = &http.Request{
		URL: &url.URL{},
	}

	ctx.Request = request
}

func TestCreate(t *testing.T) {
	assertion := arrangeCreate(t)

	t.Run("Create coin", func(t *testing.T) {
		handler.Create(ctx)

		var response coinResponse.Response
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		assertion.Equal(http.StatusCreated, recorder.Code)
		assertion.Equal(strconv.Itoa(int(mainCoin.ID)), response.ID)
		assertion.Equal(mainCoin.Name.String(), response.Short.Name)
		assertion.Equal(mainCoin.Code.String(), response.Short.Code)
	})

	resetData()
}

func TestUpdate(t *testing.T) {
	assertion := arrangeUpdate(t)

	t.Run("Update coin", func(t *testing.T) {
		handler.Update(ctx)

		var response coinResponse.Response
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		assertion.Equal(http.StatusOK, recorder.Code)
		assertion.Equal(strconv.Itoa(int(mainCoin.ID)), response.ID)
		assertion.Equal(mainCoin.Name.String(), response.Short.Name)
		assertion.Equal(mainCoin.Code.String(), response.Short.Code)
	})

	resetData()
}

func TestDelete(t *testing.T) {
	assertion := arrangeDelete(t)

	t.Run("Delete coin", func(t *testing.T) {
		handler.Delete(ctx)

		assertion.Equal(http.StatusNoContent, ctx.Writer.Status())
	})

	resetData()
}

func TestList(t *testing.T) {
	assertion := arrangeList(t)

	t.Run("List coins", func(t *testing.T) {
		handler.List(ctx)

		var response coinResponse.List
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		assertion.Equal(http.StatusOK, recorder.Code)
		assertion.Equal(uint64(1), response.Total)
		assertion.Equal(uint64(1), response.Page)
		assertion.Equal(uint64(10), response.Limit)
		assertion.Equal(strconv.Itoa(int(mainCoin.ID)), response.Data[0].ID)
		assertion.Equal(mainCoin.Name.String(), response.Data[0].Name)
		assertion.Equal(mainCoin.Code.String(), response.Data[0].Code)
	})

	resetData()
}

func arrangeCreate(t *testing.T) *assert.Assertions {
	makeCoin(nil)
	setBody()

	assertion := assert.New(t)

	uc.
		On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*coin.Coin"),
		).
		Return(func(c context.Context, coin *domain.Coin) *domain.Coin {
			assertion.Equal(mainCoin.Name, coin.Name)
			assertion.Equal(mainCoin.Code, coin.Code)

			mainCoin.ID = 1

			return mainCoin
		}, func(c context.Context, coin *domain.Coin) error {
			return nil
		})

	return assertion
}

func arrangeUpdate(t *testing.T) *assert.Assertions {
	id := uint(1)
	makeCoin(&id)
	setUrl()
	setBody()

	assertion := assert.New(t)

	uc.
		On(
			"Update",
			mock.Anything,
			mock.AnythingOfType("*coin.Coin"),
		).
		Return(func(c context.Context, coin *domain.Coin) *domain.Coin {
			assertion.Equal(mainCoin.Name, coin.Name)
			assertion.Equal(mainCoin.Code, coin.Code)
			assertion.Equal(mainCoin.ID, coin.ID)

			return mainCoin
		}, func(c context.Context, coin *domain.Coin) error {
			return nil
		})

	return assertion
}

func arrangeDelete(t *testing.T) *assert.Assertions {
	id := uint(1)
	makeCoin(&id)
	setUrl()

	assertion := assert.New(t)

	uc.
		On(
			"Delete",
			mock.Anything,
			mock.AnythingOfType("uint"),
		).
		Return(func(ctx context.Context, coinId uint) error {
			assertion.Equal(mainCoin.ID, coinId)

			return nil
		})

	return assertion
}

func arrangeList(t *testing.T) *assert.Assertions {
	id := uint(1)
	makeCoin(&id)

	assertion := assert.New(t)

	uc.
		On(
			"List",
			mock.Anything,
			mock.AnythingOfType("queryParameter.QueryParameter"),
		).
		Return(func(ctx context.Context, params queryParameter.QueryParameter) []*domain.Coin {
			var coins []*domain.Coin
			coins = append(coins, mainCoin)

			return coins
		}, func(ctx context.Context, params queryParameter.QueryParameter) error {
			return nil
		})

	uc.
		On(
			"Count",
			mock.Anything,
		).
		Return(func(c context.Context) uint64 {
			return 1
		}, func(c context.Context) error {
			return nil
		})

	return assertion
}

func setBody() {
	marshal, err := json.Marshal(mainCoin)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(marshal))
}

func setUrl() {
	id := strconv.Itoa(int(mainCoin.ID))
	q := request.URL.Query()
	q.Add("id", id)

	request.URL.RawQuery = q.Encode()

	ctx.Request = request

	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id,
		},
	}
}

func makeCoin(ID *uint) {
	coinName, _ := name.New("Bitcoin")
	coinCode, _ := code.New("BTC")
	mainCoin = domain.New(
		*coinName,
		*coinCode,
		nil,
	)

	if ID != nil {
		mainCoin.ID = *ID
	}
}
