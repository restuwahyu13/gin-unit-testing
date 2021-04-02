package main

import (
	"net/http"
	"testing"

	"go-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUnitTestingExampleGET(t *testing.T) {

	router := SetupRouter()

	rr := utils.HttpTestRequest(router, "GET", "/", nil)

	response := utils.Parse(rr.Body.Bytes())

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "successfuly", response.Status)
	assert.Equal(t, "get request successfuly", response.Message)
	assert.Equal(t, nil, response.Data)
}

func TestUnitTestingExamplePOST(t *testing.T) {

	router := SetupRouter()

	payload := gin.H{
		"name": "restu wahyu saputra",
		"age":  25,
	}

	rr := utils.HttpTestRequest(router, "POST", "/", utils.Strigify(payload))

	response := utils.Parse(rr.Body.Bytes())

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "successfuly", response.Status)
	assert.Equal(t, "post request successfuly", response.Message)
}
