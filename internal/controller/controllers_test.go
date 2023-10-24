package controller

import (
	"backend/mocks"
	"backend/pkg/models"
	services "backend/services/articles"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAllArticle(t *testing.T) {
	testCases := []struct {
		name                    string
		mockDBAllArticlesReturn []models.Article
		expectedStatusCode      int
		expectedResponseBody    string
	}{
		{
			name: "Success Case",
			mockDBAllArticlesReturn: []models.Article{
				{ID: 1, Title: "Article 1", Content: "Content 1", Author: "Author 1"},
				{ID: 2, Title: "Article 2", Content: "Content 2", Author: "Author 2"},
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"Success","data":[{"id":1,"title":"Article 1","content":"Content 1","author":"Author 1"},{"id":2,"title":"Article 2","content":"Content 2","author":"Author 2"}]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create a mock object for your DBInterface
			mockDB := mocks.NewMockDBInterface(ctrl)
			// Set expectations for mockDB's AllArticles method
			mockDB.EXPECT().AllArticles().Return(tc.mockDBAllArticlesReturn, nil)

			// Create an instance of your Application with the mock dependencies
			app := &Controller{
				ArticleService: services.NewArticleService(mockDB),
			}

			// Create an HTTP request and response recorder for testing
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/articles", nil)

			// Call the function to be tested
			app.AllArticle(w, r)

			// Assertions for response status code and body
			if w.Code != tc.expectedStatusCode {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatusCode, w.Code)
			}
			assert.JSONEq(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
func TestAllArticle_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for your DBInterface
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Set expectations for mockDB's AllArticles method to return an error
	mockDB.EXPECT().AllArticles().Return(nil, errors.New("some error"))

	// Create an instance of your Application with the mock dependencies
	app := &Controller{
		ArticleService: services.NewArticleService(mockDB),
	}

	// Create an HTTP request and response recorder for testing
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/articles", nil)

	// Call the function to be tested
	app.AllArticle(w, r)

	// Assertions for response status code and body
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, w.Code)
	}

	// You can also check the response body here to ensure it matches your expected error response.
}

func TestInsertArticle1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for your ArticleService
	// mockArticleService := mocks.NewMockArticleServices(ctrl)

	// Create a mock object for your utility package
	// mockUtility := mocks.NewMockUtilityInterface(ctrl)
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Create an instance of your Application with the mock dependencies
	app := &Controller{
		ArticleService: services.NewArticleService(mockDB),
	}

	// Define a test table with different scenarios
	testCases := []struct {
		name             string
		sampleArticle    *models.Article
		requestBody      string
		expectedStatus   int
		expectedResponse string
		mockDBExpect     func(db *mocks.MockDBInterface)
	}{
		{
			name: "Successful Insert",
			sampleArticle: &models.Article{
				ID:      1,
				Title:   "Sample Article",
				Content: "Sample Content",
				Author:  "Sample Author",
			},
			requestBody:      `{"ID": 1, "Title": "Sample Article", "Content": "Sample Content", "Author": "Sample Author"}`,
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{"status":201,"message":"Success","data":{"id":1}}`,
			mockDBExpect: func(db *mocks.MockDBInterface) {
				db.EXPECT().CreateArticle(gomock.Any()).Return(1, nil)
			},
		},
		{
			name:             "Error Parsing JSON",
			sampleArticle:    nil,
			requestBody:      "{invalid-json}",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"status":400,"message":"Error parsing JSON request: ","data":null}`,
			mockDBExpect:     func(db *mocks.MockDBInterface) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set expectations for the CreateArticle method in your mockDB
			tc.mockDBExpect(mockDB)

			r, _ := http.NewRequest("POST", "/articles", bytes.NewBufferString(tc.requestBody))
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			// Call the function to be tested
			app.InsertArticle(w, r)

			// Assertions for response status code and body
			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatus, w.Code)
			}
			if w.Body.String() != tc.expectedResponse {
				t.Errorf("Expected response body:\n%s\n\nGot:\n%s", tc.expectedResponse, w.Body.String())
			}
		})
	}
}

func TestInsertArticle_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for your DBInterface
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Create an instance of your Application with the mock dependencies
	app := &Controller{
		ArticleService: services.NewArticleService(mockDB),
	}

	// Create a sample article for your test
	sampleArticle := &models.Article{
		ID:      1,
		Title:   "Sample Article",
		Content: "Sample Content",
		Author:  "Sample Author",
	}

	// Set expectations for the CreateArticle method in your mockDB to return an error
	mockDB.EXPECT().CreateArticle(gomock.Any()).Return(0, errors.New("some error"))

	// Create an HTTP request with the sample article as the JSON body
	body, err := json.Marshal(sampleArticle)
	if err != nil {
		t.Fatal("Failed to marshal JSON:", err)
	}

	r, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	r.Header.Set("Content-Type", "application/json")

	// Create an HTTP response recorder for testing
	w := httptest.NewRecorder()

	// Call the function to be tested
	app.InsertArticle(w, r)

	// Assertions for response status code
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, w.Code)
	}

}

func TestGetArticle_Error(t *testing.T) {
	testCases := []struct {
		name                   string
		mockArticleService     func(ctrl *gomock.Controller) services.ArticleServices
		mockUtility            func(ctrl *gomock.Controller) UtilityInterface
		mockDBGetArticleReturn *models.Article
		expectedStatusCode     int
		expectedResponseBody   string
	}{

		{
			name: "Error In Retrieving Article",
			mockArticleService: func(ctrl *gomock.Controller) services.ArticleServices {
				mock := mocks.NewMockArticleServices(ctrl)
				mock.EXPECT().GetArticleByID(3).Return(nil, errors.New("some error"))
				return mock
			},
			mockUtility: func(ctrl *gomock.Controller) UtilityInterface {
				mock := mocks.NewMockUtilityInterface(ctrl)
				mock.EXPECT().WriteJSON(gomock.Any(), http.StatusInternalServerError, gomock.Any()).Return(nil)
				return mock
			},
			mockDBGetArticleReturn: nil,
			expectedStatusCode:     http.StatusBadRequest,
			expectedResponseBody:   `{"status":400,"message":"Error parsing article ID: strconv.Atoi: parsing \"\": invalid syntax","data":null}`, // Fix expected response body
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create the mocks for the article service, utility, and database
			//	mockArticleService := tc.mockArticleService(ctrl)
			// mockUtility := tc.mockUtility(ctrl)
			// mockDB := mocks.NewMockDBInterface(ctrl)

			// Create a mock object for your DBInterface
			mockDB := mocks.NewMockDBInterface(ctrl)
			// Set expectations for mockDB's AllArticles method
			//mockDB.EXPECT().AllArticles().Return(tc.mockDBAllArticlesReturn, nil)

			// Create an instance of your Application with the mock dependencies
			app := &Controller{
				ArticleService: services.NewArticleService(mockDB),
			}

			r, _ := http.NewRequest("GET", "/articles/1", nil)
			w := httptest.NewRecorder()

			app.GetArticle(w, r)

			// Check the expected status code
			if w.Code != tc.expectedStatusCode {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatusCode, w.Code)
			}

			// Check the expected response body
			if w.Body.String() != tc.expectedResponseBody {
				t.Errorf("Expected response body:\n%s\n\nGot:\n%s", tc.expectedResponseBody, w.Body.String())
			}
		})
	}
}

func TestHealthCheck(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Success Case",
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"Success","data":"Go articles up and running"}`,
		},
	}

	app := &Controller{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create an HTTP request
			r, _ := http.NewRequest("GET", "/healthCheck", nil)

			// Create an HTTP response recorder for testing
			w := httptest.NewRecorder()

			// Call the function to be tested
			app.HealthCheck(w, r)

			// Check the expected status code
			if w.Code != tc.expectedStatusCode {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatusCode, w.Code)
			}

			// Check the expected response body
			assert.JSONEq(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
