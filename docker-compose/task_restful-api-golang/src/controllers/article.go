package controllers

import (
	"example/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"time"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleController  represent the httpHandler for article
type ArticleController struct {
	base     BaseController
	AUseCase models.ArticleService
}

// NewArticleController will initialize the articles/ resources endpoint
func NewArticleController(e *gin.Engine, us models.ArticleService) {
	handler := &ArticleController{
		AUseCase: us,
	}

	grp1 := e.Group("/v1")
	{
		grp1.GET("/articles", handler.FetchArticle)
		grp1.POST("/articles", handler.Store)
		grp1.GET("/articles/:id", handler.GetByID)
		grp1.DELETE("/articles/:id", handler.Delete)
	}
}

// FetchArticle will fetch the article based on given params
func (a *ArticleController) FetchArticle(c *gin.Context) {
	startTime := time.Now().Unix()

	listAr, err := a.AUseCase.Fetch(c.Request.Context())
	if err != nil {
		log.WithFields(log.Fields{
			"status_code":   400,
			"start_time":    startTime,
			"error_details": err.Error(),
		}).Error("Failed to get all article!")

		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	log.WithFields(log.Fields{
		"status_code": http.StatusOK,
		"start_time":  startTime,
	}).Info(listAr)

	c.JSON(http.StatusOK, listAr)
}

// GetByID will get article by given id
func (a *ArticleController) GetByID(c *gin.Context) {
	startTime := time.Now().Unix()

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"status_code":   400,
			"start_time":    startTime,
			"error_details": err.Error(),
		}).Error("Failed to parse article id from param!")

		c.JSON(http.StatusNotFound, ErrNotFound.Error())
	}

	id := int64(idP)
	art, err := a.AUseCase.GetByID(c.Request.Context(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"status_code":   400,
			"start_time":    startTime,
			"error_details": err.Error(),
		}).Error("Failed to get article by id!")

		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	log.WithFields(log.Fields{
		"status_code": http.StatusOK,
		"start_time":  startTime,
	}).Info(art)

	c.JSON(http.StatusOK, art)
}

func isRequestValid(m *models.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body
func (a *ArticleController) Store(c *gin.Context) {
	startTime := time.Now().Unix()

	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		log.WithFields(log.Fields{
			"status_code":     400,
			"start_time":      startTime,
			"error_details":   err.Error(),
			"request_content": c.Request.Body,
		}).Error("Failed to parse article content!")

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var ok bool
	if ok, err = isRequestValid(&article); !ok {
		log.WithFields(log.Fields{
			"status_code":     400,
			"start_time":      startTime,
			"error_details":   err.Error(),
			"request_content": c.Request.Body,
		}).Error("Article is invalid!")

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = a.AUseCase.Store(c.Request.Context(), &article)
	if err != nil {
		log.WithFields(log.Fields{
			"status_code":     400,
			"start_time":      startTime,
			"error_details":   err.Error(),
			"request_content": c.Request.Body,
		}).Error("Failed to create article object!")

		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	log.WithFields(log.Fields{
		"status_code":     http.StatusCreated,
		"start_time":      startTime,
		"request_content": c.Request.Body,
	}).Info(article)

	c.JSON(http.StatusCreated, article)
}

// Delete will delete article by given param
func (a *ArticleController) Delete(c *gin.Context) {
	startTime := time.Now().Unix()

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"status_code":   400,
			"start_time":    startTime,
			"error_details": err.Error(),
		}).Error("Failed to parse article id from param!")

		c.JSON(http.StatusNotFound, ErrNotFound.Error())
		return
	}

	id := int64(idP)
	err = a.AUseCase.Delete(c.Request.Context(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"status_code":   400,
			"start_time":    startTime,
			"error_details": err.Error(),
		}).Error("Failed to delete article by id!")

		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	log.WithFields(log.Fields{
		"status_code":     http.StatusOK,
		"start_time":      startTime,
		"request_content": c.Request.Body,
	}).Info(fmt.Sprintf("Delete successfully article with id=%d", id))

	c.AbortWithStatus(http.StatusNoContent)
}
