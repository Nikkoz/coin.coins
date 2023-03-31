package coin

import (
	"coins/internal/domain/coin"
	"coins/pkg/protobuf/coins"
	grpcCoin "github.com/Nikkoz/coin.sync/pkg/protobuf/coins"
)

func modelsToPb(models []*coin.Coin) []*coins.Coin {
	pbCoins := make([]*coins.Coin, len(models))
	for k, model := range models {
		pbCoins[k] = modelToPb(model, nil, false)
	}

	return pbCoins
}

func modelToPb(model *coin.Coin, c *grpcCoin.Coin, full bool) *coins.Coin {
	pb := &coins.Coin{
		Id:   uint64(model.ID),
		Name: model.Name.String(),
		Code: model.Code.String(),
		Icon: model.Icon.String(),
	}

	if full && c != nil {
		pb.Info = &coins.Info{
			Type:              coins.Info_Type(c.GetType()),
			IsActive:          c.GetIsActive(),
			HasSmartContracts: c.GetInfo().GetHasSmartContracts().Value,
			Platform:          c.GetInfo().GetPlatform().Value,
			DateStart:         c.GetInfo().GetDateStart(),
			MaxSupply:         uint64(c.GetInfo().GetMaxSupply().Value),
			KeyFeatures:       c.GetInfo().GetKeyFeatures().Value,
		}
	}

	return pb
}

//import (
//	"path/filepath"
//	"sync/pkg/files"
//	"sync/pkg/protobuf/coins"
//	"sync/pkg/types/context"
//	"sync/pkg/types/logger"
//	"sync/services/coins/internal/domain/coin"
//	"sync/services/coins/internal/domain/coin/types/alias"
//	"sync/services/coins/internal/domain/coin/types/code"
//	"sync/services/coins/internal/domain/coin/types/coinType"
//	"sync/services/coins/internal/domain/coin/types/marketId"
//	"sync/services/coins/internal/domain/coin/types/name"
//	"sync/services/coins/internal/domain/image"
//	"sync/services/coins/internal/domain/image/types/file"
//	imageName "sync/services/coins/internal/domain/image/types/name"
//	"sync/services/coins/internal/domain/image/types/path"
//	"sync/services/coins/internal/domain/info"
//	"sync/services/coins/internal/domain/url"
//	"sync/services/coins/internal/domain/url/types/link"
//	"sync/services/coins/internal/domain/url/types/urlType"
//	"time"
//)
//
//func ToModel(ctx context.Context, entity *coins.Coin) (*coin.Coin, error) {
//	err := entity.ValidateAll()
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	market, err := marketId.New(entity.GetMarketId())
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	coinName, err := name.New(entity.GetName())
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	coinCode, err := code.New(entity.GetCode())
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	coinAlias, err := alias.New(entity.GetAlias())
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	model := coin.New(
//		*market,
//		*coinName,
//		*coinCode,
//		*coinAlias,
//		coinType.CoinType(entity.GetType()),
//		entity.GetIsActive(),
//	)
//
//	coinInfo, err := getInfo(ctx, entity.GetInfo())
//	if err != nil {
//		return nil, err
//	}
//
//	model.CoinInfo = coinInfo
//
//	coinImage, err := getImage(ctx, entity.GetImage(), model.Code.String())
//	if err == nil {
//		model.Image = coinImage
//	}
//
//	CoinUrls := getCoinUrls(ctx, entity.GetLinks())
//	if len(CoinUrls) > 0 {
//		model.CoinUrls = CoinUrls
//	}
//
//	return model, nil
//}
//
//func ToListModel(ctx context.Context, request *coins.SaveCoinsRequest) []*coin.Coin {
//	var models []*coin.Coin
//
//	for _, v := range request.GetCoins() {
//		model, err := ToModel(ctx, v)
//		if err != nil {
//			continue
//		}
//
//		models = append(models, model)
//	}
//
//	return models
//}
//
//func getImage(ctx context.Context, logo, code string) (*image.Image, error) {
//	p, err := files.LoadImage(logo, code)
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	title, err := imageName.New(code)
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	fileName, err := file.New(filepath.Base(p))
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	filePath, err := path.New(p)
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	return image.New(
//		nil,
//		*title,
//		*fileName,
//		*filePath,
//		nil,
//		nil,
//	), nil
//}
//
//func getInfo(ctx context.Context, entity *coins.Info) (*info.Info, error) {
//	hasSmartContracts := entity.GetHasSmartContracts().GetValue()
//	platform := entity.GetPlatform().GetValue()
//	maxSupply := entity.GetMaxSupply().GetValue()
//	dateStart, err := time.Parse("2006-01-02", entity.GetDateStart())
//	if err != nil {
//		return nil, logger.ErrorWithContext(ctx, err)
//	}
//
//	return info.New(
//		nil,
//		&hasSmartContracts,
//		&platform,
//		dateStart,
//		false,
//		&maxSupply,
//		nil,
//		nil,
//		nil,
//		nil,
//	), nil
//}
//
//func getCoinUrls(ctx context.Context, entity *coins.Links) []*url.Url {
//	var coinUrls []*url.Url
//
//	for _, v := range entity.Links {
//		urlLink, err := link.New(v.Link)
//		if err != nil {
//			_ = logger.ErrorWithContext(ctx, err)
//
//			continue
//		}
//
//		ut, err := urlType.New(int32(v.Type))
//		if err != nil {
//			_ = logger.ErrorWithContext(ctx, err)
//
//			continue
//		}
//
//		newUrl := url.New(nil, *urlLink, *ut)
//
//		coinUrls = append(coinUrls, newUrl)
//	}
//
//	return coinUrls
//}
