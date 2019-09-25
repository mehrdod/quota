package controllers

import (
	"log"
	random "math/rand"
	"net/http"
	"strconv"
	"time"

	"alif/quota/app/models"

	"github.com/gin-gonic/gin"
)

// CreateQuota - creates new quota
func CreateQuota(c *gin.Context) {

	var (
		quota models.Quota
	)
	if c.ShouldBindJSON(&quota) != nil {
		replyError(c, _WRONG_PARAMS)
		return
	}
	quota.CreatedAt = time.Now()
	id := models.CreateQuota(&quota)

	c.JSON(http.StatusOK, gin.H{`id`: id})
}

// UpdateQuota - updates existing quota by id
func UpdateQuota(c *gin.Context) {
	var (
		quota models.Quota
		err   error
	)
	if c.ShouldBindJSON(&quota) != nil {
		replyError(c, _WRONG_PARAMS)
		return
	}
	quota.Id, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		replyError(c, _WRONG_PARAMS)
		return
	}
	err = models.UpdateQuota(&quota)
	if err != nil {
		log.Println(err)
		replyError(c, _INTERNAL_SERVER_ERR)
		return
	}
	c.JSON(http.StatusOK, gin.H{`status`: `ok`})

}

// RemoveQuota - removes quota by id
func RemoveQuota(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		replyError(c, _WRONG_PARAMS)
		return
	}
	err = models.RemoveQuota(id)
	if err != nil {
		log.Println(err)
		replyError(c, _INTERNAL_SERVER_ERR)
		return
	}
	c.JSON(http.StatusOK, gin.H{`status`: `ok`})
}

// GetAllQuotas - get all existing quotas
func GetAllQuotas(c *gin.Context) {

	c.JSON(http.StatusOK, models.GetAllQuotas())
}

// GetQuotasByCategory - get all quotas by category
func GetQuotaByCategory(c *gin.Context) {
	var (
		quotaCat []models.Quota
	)
	category := c.Param("category")
	for _, quota := range models.GetAllQuotas() {
		if quota.Category == category {
			quotaCat = append(quotaCat, quota)
		}
	}

	c.JSON(http.StatusOK, quotaCat)
}

// GetRandomQuota - gets random Quota
func GetRandomQuota(c *gin.Context) {
	quotaLen := models.GetQuotaLen()

	if quotaLen == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}

	i := random.Intn(quotaLen)

	c.JSON(http.StatusOK, models.GetQuotaByIndx(i))
}

func init() {
	random.Seed(time.Now().UTC().UnixNano())
}
