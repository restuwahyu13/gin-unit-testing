package main

import (
	"net/http"
	"testing"

	"go-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router = SetupRouter()

func TestUnitTestingExampleGET(t *testing.T) {

	rr, req := utils.HttpTestRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, <-request)

	response := utils.Parse(rr.Body.Bytes())

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "successfuly", response.Status)
	assert.Equal(t, "get request successfuly", response.Message)
	assert.Equal(t, nil, response.Data)
}

func TestUnitTestingExamplePOST(t *testing.T) {

	payload := gin.H{
		"name": "restu wahyu saputra",
		"age":  25,
	}

	rr, req := utils.HttpTestRequest("POST", "/", utils.Strigify(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, <-request)

	response := utils.Parse(rr.Body.Bytes())

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "successfuly", response.Status)
	assert.Equal(t, "post request successfuly", response.Message)
}
