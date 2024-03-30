package Coin

//go:generate mockery --dir ./internal/repositories/coin/interfaces --name Storage --output ./internal/repositories/coin/database/mock --outpkg mockCoinStorage
//go:generate mockery --dir ./internal/repositories/url/interfaces --name Storage --output ./internal/repositories/url/database/mock --outpkg mockUrlStorage
//go:generate mockery --dir ./internal/useCases/interfaces --name Coin --output ./internal/useCases/mock/ --outpkg mockUseCase
//go:generate mockery --dir ./internal/useCases/interfaces --name Url --output ./internal/useCases/mock/ --outpkg mockUseCase
