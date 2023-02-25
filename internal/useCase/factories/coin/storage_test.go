package coin

import (
	domain "coins/internal/domain/coin"
	"coins/internal/domain/coin/types/code"
	"coins/internal/domain/coin/types/icon"
	"coins/internal/domain/coin/types/name"
	mockStorage "coins/internal/repository/coin/database/mock"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mainCoin   *domain.Coin
	repository = new(mockStorage.Coin)
	factory    *Factory
)

func TestCreate(t *testing.T) {
	assertion := arrangeCreate(t)

	t.Run("Create coin", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.Create(ctx, mainCoin)

		assertion.NoError(err)
		assertion.Equal(mainCoin.Name, result.Name)
		assertion.Equal(mainCoin.Code, result.Code)
		assertion.Equal(mainCoin.Icon, result.Icon)
		assertion.NotZero(result, "ID")
	})
}

func TestUpdate(t *testing.T) {
	assertion := arrangeUpdate(t)

	t.Run("Update coin", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.Update(ctx, mainCoin)

		assertion.NoError(err)
		assertion.Equal(mainCoin, result)
	})
}

func TestDelete(t *testing.T) {
	assertion := arrangeDelete(t)

	t.Run("Delete coin", func(t *testing.T) {
		ctx := context.Empty()

		err := factory.Delete(ctx, mainCoin.ID)
		assertion.NoError(err)
	})
}

func TestUpsertCreate(t *testing.T) {
	assertion := arrangeUpsertCreate(t)

	t.Run("Upsert(create) coin", func(t *testing.T) {
		ctx := context.Empty()

		c := *mainCoin
		result, err := factory.Upsert(ctx, &c)

		assertion.NoError(err)
		assertion.Equal(mainCoin.Name, result[0].Name)
		assertion.Equal(mainCoin.Code, result[0].Code)
		assertion.Equal(mainCoin.Icon, result[0].Icon)
		assertion.NotZero(result[0], "ID")
	})
}

func TestUpsertUpdate(t *testing.T) {
	assertion := arrangeUpsertUpdate(t)

	t.Run("Upsert(create) coin", func(t *testing.T) {
		ctx := context.Empty()

		c := *mainCoin
		result, err := factory.Upsert(ctx, &c)

		assertion.NoError(err)
		assertion.Equal(mainCoin, result[0])
	})
}

func TestList(t *testing.T) {
	assertion := arrangeList(t)

	t.Run("List coins", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.List(ctx, queryParameter.QueryParameter{})

		assertion.NoError(err)
		assertion.Equal(mainCoin, result[0])
	})
}

func TestCount(t *testing.T) {
	assertion := arrangeCount(t)

	t.Run("Count coins", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.Count(ctx)

		assertion.NoError(err)
		assertion.Equal(uint64(1), result)
	})
}

func arrangeCreate(t *testing.T) *assert.Assertions {
	createCoin(nil)

	factory = New(repository, nil, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"CreateCoin",
			mock.Anything,
			mock.AnythingOfType("*coin.Coin"),
		).
		Return(func(ctx context.Context, coin *domain.Coin) *domain.Coin {
			assertion.Equal(mainCoin, coin)

			coin.ID = 1

			return coin
		}, func(ctx context.Context, coin *domain.Coin) error {
			return nil
		})

	return assertion
}

func arrangeUpdate(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createCoin(&ID)

	factory = New(repository, nil, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"UpdateCoin",
			mock.Anything,
			mock.AnythingOfType("*coin.Coin"),
		).
		Return(func(ctx context.Context, coin *domain.Coin) *domain.Coin {
			assertion.Equal(mainCoin, coin)

			return coin
		}, func(ctx context.Context, coin *domain.Coin) error {
			return nil
		})

	return assertion
}

func arrangeDelete(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createCoin(&ID)

	factory = New(repository, nil, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"DeleteCoin",
			mock.Anything,
			mock.AnythingOfType("uint"),
		).
		Return(func(ctx context.Context, ID uint) error {
			assertion.Equal(mainCoin.ID, ID)

			return nil
		})

	return assertion
}

func arrangeUpsertCreate(t *testing.T) *assert.Assertions {
	createCoin(nil)

	factory = New(repository, nil, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"UpsertCoins",
			mock.Anything,
			mock.Anything,
		).
		Return(func(ctx context.Context, coins ...*domain.Coin) []*domain.Coin {
			assertion.Equal(mainCoin, coins[0])

			coins[0].ID = 10

			return coins
		}, func(ctx context.Context, coins ...*domain.Coin) error {
			return nil
		})

	return assertion
}

func arrangeUpsertUpdate(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createCoin(&ID)

	factory = New(repository, nil, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"UpsertCoins",
			mock.Anything,
			mock.Anything,
		).
		Return(func(ctx context.Context, coins ...*domain.Coin) []*domain.Coin {
			assertion.Equal(mainCoin, coins[0])

			return coins
		}, func(ctx context.Context, coins ...*domain.Coin) error {
			return nil
		})

	return assertion
}

func arrangeList(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createCoin(&ID)

	factory = New(repository, nil, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"ListCoins",
			mock.Anything,
			mock.AnythingOfType("queryParameter.QueryParameter"),
		).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) []*domain.Coin {
			var coins []*domain.Coin
			coins = append(coins, mainCoin)

			return coins
		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})

	return assertion
}

func arrangeCount(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createCoin(&ID)

	factory = New(repository, nil, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"CountCoins",
			mock.Anything,
		).
		Return(func(ctx context.Context) uint64 {
			return 1
		}, func(ctx context.Context) error {
			return nil
		})

	return assertion
}

func createCoin(ID *uint) {
	coinName, _ := name.New("Bitcoin")
	coinCode, _ := code.New("BTC")
	coinIcon, _ := icon.New("/path/btc.png")
	mainCoin = domain.New(
		*coinName,
		*coinCode,
		coinIcon,
	)

	if ID != nil {
		mainCoin.ID = *ID
	}
}
