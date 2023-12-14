//go:build e2e
// +build e2e

package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-unit-test/internal/generated/mocks"
	delivery "github.com/martinyonatann/go-unit-test/internal/users/delivery/http"
	"github.com/martinyonatann/go-unit-test/internal/users/dtos"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_handlers_CreateHandler(t *testing.T) {
	echo := echo.New()
	defer echo.Shutdown(context.Background())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		uc = mocks.NewMockUseCases(ctrl)
	)

	t.Run("positive_CreateHandler", func(t *testing.T) {
		var request = dtos.Users{
			Name:     faker.Name(),
			Password: faker.Password(),
		}

		requestJSON, err := json.Marshal(request)
		require.NoError(t, err)

		uc.EXPECT().CreateUser(context.Background(), request).Return(nil)
		handlers := delivery.NewHandlers(uc)
		delivery.MapRoutes(echo, handlers)

		doRequest := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(requestJSON)))
		doRequest.Header.Set("Content-Type", "application/json")
		doResponse := httptest.NewRecorder()

		c := echo.NewContext(doRequest, doResponse)

		require.NoError(t, handlers.CreateHandler(c))
		require.Equal(t, doResponse.Code, http.StatusOK)
	})
}

func Test_handlers_DetailHandler(t *testing.T) {
	echo := echo.New()
	defer echo.Shutdown(context.Background())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		uc = mocks.NewMockUseCases(ctrl)
	)

	t.Run("positive_DetailHandler", func(t *testing.T) {
		var (
			now = time.Now()
			id  = uuid.NewString()
		)

		var expectedResponse = dtos.Users{
			ID:        id,
			Name:      faker.Name(),
			Password:  faker.Password(),
			CreatedAt: now,
		}

		uc.EXPECT().GetUserDetail(gomock.Any(), gomock.Any()).Return(expectedResponse, nil)
		handlers := delivery.NewHandlers(uc)
		delivery.MapRoutes(echo, handlers)

		fmt.Println("generated path : ", fmt.Sprintf("/users/%s", id))

		// Create a request
		doRequest := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
		doResponse := httptest.NewRecorder()

		// Create an Echo context
		c := echo.NewContext(doRequest, doResponse)

		// Call the handler
		require.NoError(t, handlers.DetailHandler(c))
		require.Equal(t, http.StatusOK, doResponse.Code)

		// Unmarshal the response body into a struct
		var response dtos.Users
		err := json.Unmarshal(doResponse.Body.Bytes(), &response)
		require.NoError(t, err)

		// Assert the response
		require.Equal(t, expectedResponse.ID, response.ID)
		require.Equal(t, expectedResponse.Name, response.Name)
		require.Equal(t, expectedResponse.Password, response.Password)
	})

}
