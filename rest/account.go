package rest

import (
	"github.com/gin-gonic/gin"
)

// @Summary detail
// @Description get detail of account
// @Tags account
// @Accept json
// @Produce json
// @Param  address path string true "account address"
// @Success 200 {object} vo.AccountVo	"success"
// @Router /api/account/{address} [get]
func (s *apiServer) GetAccount(c *gin.Context) {
}
