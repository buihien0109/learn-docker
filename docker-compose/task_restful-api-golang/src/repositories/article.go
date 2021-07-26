package repositories

import (
	"context"
	"example/src/models"
	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	base baseRepository
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewArticleRepository(dbCtx *gorm.DB) models.ArticleRepository {
	return &mysqlArticleRepository{base: baseRepository{dbCtx: dbCtx}}
}

func (m *mysqlArticleRepository) Fetch(ctx context.Context) (res []models.Article, err error) {
	var articles []models.Article
	if err = m.base.dbCtx.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}
func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int64) (models.Article, error) {
	var article models.Article
	if err := m.base.dbCtx.Where("id = ?", id).First(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (m *mysqlArticleRepository) GetByTitle(ctx context.Context, title string) (models.Article, error) {
	var article models.Article
	if err := m.base.dbCtx.Where("title = ?", title).First(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (m *mysqlArticleRepository) Store(ctx context.Context, article *models.Article) (err error) {
	if err = m.base.dbCtx.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlArticleRepository) Delete(ctx context.Context, id int64) (err error) {
	if err = m.base.dbCtx.Where("id = ?", id).Delete(&models.Article{}).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlArticleRepository) Update(ctx context.Context, ar *models.Article) (err error) {
	if err = m.base.dbCtx.Save(ar).Error; err != nil {
		return err
	}
	return nil
}
