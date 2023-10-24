package cache

import (
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/repository"
)

type articlesCacheAdaptor struct {
	client RedisClient
}

func NewArticlesCacheAdaptor(c RedisClient) repository.ArticlesCacheRepository {
	return &articlesCacheAdaptor{
		client: c,
	}
}

func (c *articlesCacheAdaptor) GetByCategory(articles *[]model.Article, slug string) error {
	return c.client.GET(GetArticlesByCategoryKey(slug), articles)
}

func (c *articlesCacheAdaptor) SetByCategory(articles *[]model.Article, slug string) error {
	return c.client.SET(GetArticlesByCategoryKey(slug), articles)
}

func (c *articlesCacheAdaptor) GetByTag(articles *[]model.Article, slug string) error {
	return c.client.GET(GetArticlesByTagKey(slug), articles)
}

func (c *articlesCacheAdaptor) SetByTag(articles *[]model.Article, slug string) error {
	return c.client.SET(GetArticlesByTagKey(slug), articles)
}

func (c *articlesCacheAdaptor) GetLastest(articles *[]model.Article, pageNumber int) error {
	return c.client.GET(GetLatestArticleListKey(pageNumber), articles)
}

func (c *articlesCacheAdaptor) SetLastest(articles *[]model.Article, pageNumber int) error {
	return c.client.SET(GetLatestArticleListKey(pageNumber), articles)
}
