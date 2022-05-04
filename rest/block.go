package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/glodnet/chain.collector/schema"
	"net/http"
	"strconv"
)

// @Summary list
// @Description get blocks
// @Tags block
// @Accept json
// @Produce json
// @Param page query int false "page num" Default(1)
// @Param size query int false "page size" Default(10)
// @Success 200 {object} []schema.Block "success"
// @Router /api/blocks [get]
func (s *apiServer) GetBlocks(c *gin.Context) {
	page := PageRequest{
		PageNum:  1,
		PageSize: 10,
	}
	if err := c.BindQuery(&page); err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeInValidParam, err))
		return
	}

	items, total, err := schema.QueryBlocks(s.db, page.PageNum, page.PageSize)
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeUnKnown, err))
		return
	}

	pageTotal := total / int64(page.PageSize)
	if total%int64(page.PageSize) != 0 {
		pageTotal += 1
	}
	c.JSON(http.StatusOK, NewResponse(CodeSuccess, &PageItems{
		PageRequest: page,
		PageResponse: PageResponse{
			Total:     total,
			PageTotal: pageTotal,
		},
		Items: items,
	}))
	return
}

// @Summary block validatorset
// @Description get  block validatorset
// @Tags block
// @Accept  json
// @Produce  json
// @Param   page    query   int false    "page num" Default(1)
// @Param   size   query   int false    "page size" Default(10)
// @Param   height   path   int true    "block height"
// @Success 200 {object} []schema.PreCommit	"success"
// @Router /api/block/validatorset/{height} [get]
func (s *apiServer) GetBlockPreCommits(c *gin.Context) {
	page := PageRequest{
		PageNum:  1,
		PageSize: 10,
	}
	if err := c.BindQuery(&page); err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeInValidParam, err))
		return
	}

	height, err := strconv.ParseInt(c.Param("height"), 10, 0)
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeInValidParam, err))
		return
	}

	items, total, err := schema.QueryPreCommitsByHeight(s.db, height, page.PageNum, page.PageSize)
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeUnKnown, err))
		return
	}

	pageTotal := total / int64(page.PageSize)
	if total%int64(page.PageSize) != 0 {
		pageTotal += 1
	}
	c.JSON(http.StatusOK, NewResponse(CodeSuccess, &PageItems{
		PageRequest: page,
		PageResponse: PageResponse{
			Total:     total,
			PageTotal: pageTotal,
		},
		Items: items,
	}))
	return
}

type LatestHeightRespond struct {
	BlockHeightLcd int64 `json:"block_height_lcd"`
	BlockHeightDB  int64 `json:"block_height_db"`
}

// @Summary latest
// @Description get block latest
// @Tags block
// @Accept json
// @Produce json
// @Success 200 {object} LatestHeightRespond	"success"
// @Router /api/blocks/latest [get]
func (s *apiServer) GetBlockLatest(c *gin.Context) {
	item, err := schema.QueryBlockLatestHeight(s.db)
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeUnKnown, err))
		return
	}

	resp, err := s.client.BlockLatest()
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeUnKnown, err))
		return
	}

	c.JSON(http.StatusOK, NewResponse(CodeSuccess, LatestHeightRespond{
		BlockHeightLcd: resp.Block.Header.Height,
		BlockHeightDB:  item.Height,
	}))
	return
}

// @Summary detail
// @Description get block info
// @Tags block
// @Accept json
// @Produce json
// @Param height path int true "block height"
// @Success 200 {object} schema.Block "success"
// @Router /api/block/{height} [get]
func (s *apiServer) GetBlockByHeight(c *gin.Context) {
	height, err := strconv.ParseInt(c.Query("height"), 10, 0)
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeInValidParam, err))
		return
	}

	item, err := schema.QueryBlockByHeight(s.db, height)
	if err != nil {
		s.logger.Error(c.Request.URL.Path, "error", err)
		c.JSON(http.StatusOK, NewResponse(CodeUnKnown, err))
		return
	}

	c.JSON(http.StatusOK, NewResponse(CodeSuccess, item))
	return
}
