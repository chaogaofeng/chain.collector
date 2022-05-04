package rest

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/glodnet/chain.go/restclient"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/gorm"
	"os"
)

type apiServer struct {
	logger log.Logger
	db     *gorm.DB
	client *restclient.RestClient
}

func NewApiServer(logger log.Logger, db *gorm.DB, client *restclient.RestClient) *apiServer {
	return &apiServer{
		logger: logger,
		db:     db,
		client: client,
	}
}

func (srv *apiServer) Start(addr string) {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	if gin.DebugMode == gin.Mode() {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	api := router.Group("api")
	api.GET("/blocks", srv.GetBlocks)
	api.GET("/blocks/latest", srv.GetBlockLatest)
	api.GET("/block/:height", srv.GetBlockByHeight)
	api.GET("/block/validatorset/:height", srv.GetBlockPreCommits)

	api.GET("/txs", srv.GetTxs)
	api.GET("/tx/:hash", srv.GetTx)


	if err := router.Run(addr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
