package Coin

//go:generate mockery --dir ./internal/useCase/adapters/storage --name Coin --output ./internal/repository/coin/database/mock --outpkg mockCoinStorage
//go:generate mockery --dir ./internal/useCase/adapters/storage --name Url --output ./internal/repository/url/database/mock --outpkg mockUrlStorage
//go:generate mockery --dir ./internal/useCase/interfaces --name Coin --output ./internal/useCase/mock/ --outpkg mockUseCase
//go:generate mockery --dir ./internal/useCase/interfaces --name Url --output ./internal/useCase/mock/ --outpkg mockUseCase
