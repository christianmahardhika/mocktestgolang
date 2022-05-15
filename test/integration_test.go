package test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/christianmahardhika/mocktestgolang/server"
	"github.com/stretchr/testify/suite"
)

type integrationTestSuite struct {
	suite.Suite
	dbConnString string
	dbName       string
	port         string
	serverHost   string
}

func RunIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &integrationTestSuite{})
}

func (s *integrationTestSuite) SetupTestServer(t *testing.T) {
	s.dbConnString = "mongodb://root:root@localhost:27017"
	dbName := "mocktestgolang_test"
	FiberApp := server.InitiateServer(s.dbConnString, dbName)
	s.port = "8080"
	s.serverHost = "localhost"
	server.StartApplication(FiberApp, s.port)
}

func (s *integrationTestSuite) Test_Integration_CreateTodo(t *testing.T) {
	requestPayload := `{
		"todo": {
			"title": "this is title"
			},
		"todo_detail": [
			{
				"item":   "item 1"
			},
			{
				"item":   "item 2"
			},
			{
				"item":   "item 3"
			}
		]
	}`
	req, err := http.NewRequest("POST", "http://"+s.serverHost+":"+s.port+"/todo", strings.NewReader(requestPayload))
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	response.Body.Close()
	s.Equal(http.StatusCreated, response.StatusCode)
}
