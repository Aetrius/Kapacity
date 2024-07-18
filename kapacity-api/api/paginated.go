package api

import (
	"fmt"
	output "kapacity-api/internal/prettyconsole"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getPaginatedData(c *gin.Context, db gorm.DB, model interface{}, pageQuery, pageSizeQuery string, additionalParams map[string]interface{}, sortBy string, order string) {
	pageStr := pageQuery
	pageSizeStr := pageSizeQuery

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid '%s' value", pageQuery)})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid '%s' value", pageSizeQuery)})
		return
	}

	modelType := reflect.TypeOf(model).Elem()
	data := reflect.New(reflect.SliceOf(modelType)).Interface()

	var query *gorm.DB

	//add .Debug() query to the gorm queries if the debug option is enabled in env vars
	//check if sort by and order is applied
	if sortBy != "" && order != "" {
		//sort by desc and order
		orderClause := strings.ToUpper(order) + " " + sortBy
		if strings.ToLower(order) == "desc" {
			query = db.Model(model).Order(orderClause)
		} else if strings.ToLower(order) == "asc" {
			query = db.Model(model).Order(orderClause)
		} else {
			query = db.Model(model).Order(orderClause)
		}
	} else {

		query = db.Model(model)
	}

	for key, value := range additionalParams {
		query = query.Debug().Where(key, value)
	}

	// Calculate the total count of records based on the additionalParams
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	// Calculate totalPages based on totalCount and pageSize
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	nextPage := 0
	if totalPages > page {
		nextPage = page + 1
	}

	var nextURL string
	if nextPage > 0 {
		nextURL = fmt.Sprintf("%s?page=%d&pageSize=%d", c.FullPath(), nextPage, pageSize)
	}

	if err := query.Scopes(paginate(pageStr, pageSizeStr)).Find(data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_pages": totalPages,
		"next_url":    nextURL,
		"page":        page,
		"page_size":   pageSize,
		"data":        data,
	})
}

func getPaginatedDataWithQuery(c *gin.Context, db *gorm.DB, model interface{}, pageQuery, pageSizeQuery string, additionalParams map[string]interface{}, sortBy string, order string) {
	pageStr := pageQuery
	pageSizeStr := pageSizeQuery

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid '%s' value", pageQuery)})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid '%s' value", pageSizeQuery)})
		return
	}

	var query *gorm.DB

	//add .Debug() query to the gorm queries if the debug option is enabled in env vars
	//check if sort by and order is applied
	if sortBy != "" && order != "" {
		//sort by desc and order
		orderClause := strings.ToUpper(order) + " " + sortBy
		if strings.ToLower(order) == "desc" {
			query = db.Model(model).Order(orderClause)
			output.DEBUG(fmt.Sprintf("Descending"))
		} else if strings.ToLower(order) == "asc" {
			query = db.Model(model).Order(orderClause)
			output.DEBUG(fmt.Sprintf("Ascending"))
		} else {
			query = db.Model(model).Order(orderClause)
			output.DEBUG(fmt.Sprintf("Default"))
		}
	} else {

		query = db.Model(model)
	}

	for key, value := range additionalParams {
		query = query.Where(key, value)
	}

	// Clone the query for count without pagination
	var countQuery *gorm.DB
	if countQuery = db.Model(model); countQuery == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create count query"})
		return
	}

	for key, value := range additionalParams {
		countQuery = countQuery.Where(key, value)
	}

	var totalCount int64
	if err := countQuery.Debug().Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	nextPage := 0
	if totalPages > page {
		nextPage = page + 1
	}

	var nextURL string
	if nextPage > 0 {
		nextURL = fmt.Sprintf("%s?page=%d&pageSize=%d", c.FullPath(), nextPage, pageSize)
	}

	results := reflect.New(reflect.SliceOf(reflect.TypeOf(model))).Interface()

	if err := query.Debug().Scopes(paginate(pageStr, pageSizeStr)).Find(results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_pages": totalPages,
		"next_url":    nextURL,
		"page":        page,
		"page_size":   pageSize,
		"data":        results,
	})
}

func paginate(page, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Convert page and pageSize strings to integers
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return db
		}

		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			return db
		}

		offset := (pageInt - 1) * pageSizeInt
		return db.Offset(offset).Limit(pageSizeInt)
	}
}
