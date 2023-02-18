package broker

import (
	"coins/internal/domain/coin"
	coinCode "coins/internal/domain/coin/types/code"
	coinIcon "coins/internal/domain/coin/types/icon"
	coinName "coins/internal/domain/coin/types/name"
	coinUrl "coins/internal/domain/url"
	urlExternalId "coins/internal/domain/url/types/externalId"
	urlLink "coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
	"coins/internal/repository/coin/broker/entities"
)

func toModels(data []entities.Coin) ([]*coin.Coin, error) {
	coins := make([]*coin.Coin, 0, len(data))

	for _, c := range data {
		newCoin, err := coinFromEntity(c)
		if err != nil {
			return nil, err
		}

		var urls []*coinUrl.Url
		if len(c.Urls) > 0 {
			for _, u := range c.Urls {
				url, err := urlFromEntity(u)
				if err != nil {
					return nil, err
				}

				urls = append(urls, url)
			}
		}

		newCoin.CoinUrls = urls
		coins = append(coins, newCoin)
	}

	return coins, nil
}

func coinFromEntity(entity entities.Coin) (*coin.Coin, error) {
	name, err := coinName.New(entity.Name)
	if err != nil {
		return nil, err
	}

	code, err := coinCode.New(entity.Code)
	if err != nil {
		return nil, err
	}

	icon, err := coinIcon.New(entity.Icon)
	if err != nil {
		return nil, err
	}

	newCoin := coin.New(*name, *code, icon)
	newCoin.ID = uint(entity.Id)

	return newCoin, nil
}

func urlFromEntity(entity entities.Url) (*coinUrl.Url, error) {
	externalId, err := urlExternalId.New(uint(entity.Id))
	if err != nil {
		return nil, err
	}

	link, err := urlLink.New(entity.Link)
	if err != nil {
		return nil, err
	}

	social, err := socialMedia.New(entity.Type)
	if err != nil {
		return nil, err
	}

	return coinUrl.New(*externalId, *link, *social), nil
}
