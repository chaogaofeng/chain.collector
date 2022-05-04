package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/glodnet/chain.collector/schema"
	"net/http"
)

// @Summary list
// @Description get txs list
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   page    query   int false    "page num" Default(1)
// @Param   size   query   int false    "page size" Default(10)
// @Param   height   query   int64 false    "height"
// @Param   txType   query   string false    "txType"
// @Param   status   query   string false    "status" Enums(success,fail)
// @Param   address   query   string false    "address"
// @Param   beginTime   query  int64 false    "beginTime"
// @Param   endTime   query   int64 false    "endTime"
// @Success 200 {object} []schema.Transaction	"success"
// @Router /api/txs [get]
func (s *apiServer) GetTxs(c *gin.Context) {

}

// @Summary tx detail
// @Description get txs detail
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   hash    path   string true  "tx hash"
// @Success 200 {object} schema.Transaction	"success"
// @Router /api/tx/{hash} [get]
func (s *apiServer) GetTx(c *gin.Context) {
	hash := c.Param("hash")

	item, err := schema.QueryTxByHash(s.db, hash)
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeUnKnown, err))
		return
	}

	c.JSON(http.StatusOK, NewResponse(CodeSuccess, item))
	return
}

// @Summary txsByAddress
// @Description txsByAddress
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   address  path   string true    "address"
// @Param   page   path   int64 true    "pagenum"
// @Param   size   path   int64 true    "pagesize"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.PageVo	"success"
// @Router /api/txsByAddress/{address}/{page}/{size} [get]
func registerQueryTxsByAccount(c *gin.Context) {

}
