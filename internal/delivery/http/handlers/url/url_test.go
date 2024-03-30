package url

import (
	"bytes"
	url2 "coins/internal/delivery/http/responses/url"
	domainCoin "coins/internal/domain/coin"
	"coins/internal/domain/coin/types/code"
	"coins/internal/domain/coin/types/name"
	domainUrl "coins/internal/domain/url"
	"coins/internal/domain/url/types/externalId"
	"coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
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
	uc       = new(mockUseCase.Url)
	mainUrl  *domainUrl.Url
	mainCoin *domainCoin.Coin
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

	t.Run("Create url", func(t *testing.T) {
		handler.Create(ctx)

		var response url2.Response
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		assertion.Equal(http.StatusCreated, recorder.Code)
		assertion.Equal(strconv.Itoa(int(mainUrl.ID)), response.ID)
		assertion.Equal(mainUrl.Link.String(), response.Link)
		assertion.Equal(mainUrl.SocialMedia.String(), response.Type)
	})

	resetData()
}

func TestUpdate(t *testing.T) {
	assertion := arrangeUpdate(t)

	t.Run("Update url", func(t *testing.T) {
		handler.Update(ctx)

		var response url2.Response
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		assertion.Equal(http.StatusOK, recorder.Code)
		assertion.Equal(strconv.Itoa(int(mainUrl.ID)), response.ID)
		assertion.Equal(mainUrl.Link.String(), response.Link)
		assertion.Equal(mainUrl.SocialMedia.String(), response.Type)
	})

	resetData()
}

func TestDelete(t *testing.T) {
	assertion := arrangeDelete(t)

	t.Run("Delete url", func(t *testing.T) {
		handler.Delete(ctx)

		assertion.Equal(http.StatusNoContent, ctx.Writer.Status())
	})

	resetData()
}

func TestList(t *testing.T) {
	assertion := arrangeList(t)

	t.Run("List urls", func(t *testing.T) {
		handler.List(ctx)

		var response url2.List
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		assertion.Equal(http.StatusOK, recorder.Code)
		assertion.Equal(uint64(1), response.Total)
		assertion.Equal(uint64(1), response.Page)
		assertion.Equal(uint64(10), response.Limit)
		assertion.Equal(strconv.Itoa(int(mainUrl.ID)), response.Data[0].ID)
		assertion.Equal(mainUrl.Link.String(), response.Data[0].Link)
		assertion.Equal(mainUrl.SocialMedia.String(), response.Data[0].Type)
	})

	resetData()
}

func arrangeCreate(t *testing.T) *assert.Assertions {
	makeUrl(nil)
	setBody()
	setUrl(true, false)

	assertion := assert.New(t)

	uc.
		On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*url.Url"),
		).
		Return(func(c context.Context, url *domainUrl.Url) *domainUrl.Url {
			assertion.Equal(mainUrl.Link, url.Link)
			assertion.Equal(mainUrl.SocialMedia, url.SocialMedia)

			mainUrl.ID = 1

			return mainUrl
		}, func(c context.Context, url *domainUrl.Url) error {
			return nil
		})

	return assertion
}

func arrangeUpdate(t *testing.T) *assert.Assertions {
	urlID := uint(1)
	makeUrl(&urlID)
	setUrl(true, true)
	setBody()

	assertion := assert.New(t)

	uc.
		On(
			"Update",
			mock.Anything,
			mock.AnythingOfType("*url.Url"),
		).
		Return(func(c context.Context, url *domainUrl.Url) *domainUrl.Url {
			assertion.Equal(mainCoin.ID, url.CoinID)
			assertion.Equal(mainUrl.Link, url.Link)
			assertion.Equal(mainUrl.SocialMedia, url.SocialMedia)
			assertion.Equal(mainUrl.ID, url.ID)

			return mainUrl
		}, func(c context.Context, url *domainUrl.Url) error {
			return nil
		})

	return assertion
}

func arrangeDelete(t *testing.T) *assert.Assertions {
	urlId := uint(1)
	makeUrl(&urlId)
	setUrl(true, true)

	assertion := assert.New(t)

	uc.
		On(
			"Delete",
			mock.Anything,
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("uint"),
		).
		Return(func(ctx context.Context, urlId uint, coinId uint) error {
			assertion.Equal(mainUrl.ID, urlId)
			assertion.Equal(mainUrl.CoinID, coinId)

			return nil
		})

	return assertion
}

func arrangeList(t *testing.T) *assert.Assertions {
	urlId := uint(1)
	makeUrl(&urlId)
	setUrl(true, false)

	assertion := assert.New(t)

	uc.
		On(
			"List",
			mock.Anything,
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("queryParameter.QueryParameter"),
		).
		Return(func(ctx context.Context, coinId uint, params queryParameter.QueryParameter) []*domainUrl.Url {
			assertion.Equal(mainUrl.CoinID, coinId)

			var urls []*domainUrl.Url
			urls = append(urls, mainUrl)

			return urls
		}, func(ctx context.Context, coinId uint, params queryParameter.QueryParameter) error {
			return nil
		})

	uc.
		On(
			"Count",
			mock.Anything,
			mock.AnythingOfType("uint"),
		).
		Return(func(c context.Context, coinId uint) uint64 {
			assertion.Equal(mainUrl.CoinID, coinId)

			return 1
		}, func(c context.Context, coinId uint) error {
			return nil
		})

	return assertion
}

func makeUrl(ID *uint) {
	makeCoin()

	urlLink, _ := link.New("https://test.com")
	urlType, _ := socialMedia.New("twitter")
	urlExternalId, _ := externalId.New(1)
	mainUrl = domainUrl.New(
		urlExternalId,
		*urlLink,
		*urlType,
	)

	if ID != nil {
		mainUrl.ID = *ID
	}

	mainUrl.CoinID = mainCoin.ID
}

func makeCoin() {
	coinName, _ := name.New("Bitcoin")
	coinCode, _ := code.New("BTC")
	mainCoin = domainCoin.New(
		*coinName,
		*coinCode,
		nil,
	)

	mainCoin.ID = 1
}

func setBody() {
	marshal, err := json.Marshal(map[string]interface{}{
		"link": mainUrl.Link.String(),
		"Type": mainUrl.SocialMedia.String(),
	})
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(marshal))
}

func setUrl(setCoinId, setUrlId bool) {
	q := request.URL.Query()

	if setCoinId {
		coinId := strconv.Itoa(int(mainUrl.CoinID))
		q.Add("id", coinId)

		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "id",
			Value: coinId,
		})
	}

	if setUrlId {
		urlId := strconv.Itoa(int(mainUrl.ID))
		q.Add("urlId", urlId)

		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "urlId",
			Value: urlId,
		})
	}

	request.URL.RawQuery = q.Encode()

	ctx.Request = request
}
