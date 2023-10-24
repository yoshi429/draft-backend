package registory

import (
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/logger"
	"gorm.io/gorm"
)

func NewTopPageRegistory(ctx *request.Context, l logger.Logger, master, reprica func() *gorm.DB) handler.TopPageHandler {
	articleRepository := persistence.NewArticlePersistence(master, reprica)
	articleUseCase := usecase.NewArticleUseCase(articleRepository, ctx.Cache())
	return handler.NewTopPageHandler(articleUseCase, ctx, l)
}
