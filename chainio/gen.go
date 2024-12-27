package chainio

//go:generate mockgen -destination=./clients/mocks/fireblocks.go -package=mocks -mock_names=Client=MockFireblocksClient github.com/arithmic/eigensdk-go/chainio/clients/fireblocks Client
//go:generate mockgen -destination=./clients/mocks/wallet.go -package=mocks github.com/arithmic/eigensdk-go/chainio/clients/wallet Wallet
