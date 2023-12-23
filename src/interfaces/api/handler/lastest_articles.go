package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/pager"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LatestArticlesHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type latestArticlesHandler struct {
	*request.Context
	logger         logger.Logger
	articleUseCase usecase.ArticleUseCase
}

const (
	numberOfArticlePerPageAtLatestAritcles = 5
)

func NewLatestArticlesHandler(articleUseCase usecase.ArticleUseCase, c *request.Context, l logger.Logger) LatestArticlesHandler {
	return &latestArticlesHandler{
		Context:        c,
		logger:         l,
		articleUseCase: articleUseCase,
	}
}

type responseLatestAritcles struct {
	Articles []model.Article `json:"articles"`
	Pager    *pager.Pager    `json:"pager"`
}

func (h latestArticlesHandler) Get(w http.ResponseWriter, r *http.Request) error {
	var res responseLatestAritcles
	vars := mux.Vars(r)
	strPage := vars["page"]
	currentPage, err := strconv.Atoi(strPage)
	if err != nil {
		h.logger.Error("failed at convert string to integer.", zap.Error(err))
		currentPage = 1
	}

	limit := numberOfArticlePerPageAtLatestAritcles
	offset := limit * (currentPage - 1)

	cacheKey := cache.GetLatestArticleListKey(currentPage)
	err = h.Cache().GET(cacheKey, &res)
	if err == nil {
		return h.JSON(w, http.StatusOK, res)
	}

	err = h.articleUseCase.GetArticles(&res.Articles, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.logger.Warn("err no articles at latest Articles Handler")
			return h.JSON(w, http.StatusNotFound, err)
		}
		h.logger.Error("failed at get articles at latest articles.", zap.Error(err))
		return h.Error(w, http.StatusInternalServerError, err)
	}

	res.Pager, err = h.articleUseCase.GetPager(currentPage, limit)
	if err != nil {
		h.logger.Warn("failed at get pager at top page.", zap.Error(err))
		return h.Error(w, http.StatusInternalServerError, err)
	}

	err = h.Cache().SET(cacheKey, res)
	if err != nil {
		h.logger.Error("failed at set cache redis at top page.", zap.Error(err))
	}

	return h.JSON(w, http.StatusOK, res)

}