package handlers

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mError "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/usecase"
	"github.com/olteffe/avitochat/internal/usecase/mocks"
)

func TestUserHandler_Create(t *testing.T) {
	type args struct {
		user *models.Users
	}
	type mockBehavior func(r *mock_usecase.MockUser, args args)
	tests := []struct {
		name                 string
		inputBody            string
		input                args
		mock                 mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Valid",
			inputBody: `{"username": "user_1"}`,
			input: args{
				&models.Users{
					Username: "user_1",
				},
			},
			mock: func(r *mock_usecase.MockUser, args args) {
				r.EXPECT().CreateUserUseCase(args.user).Return("845ac772-cb49-433c-a871-0a98af34f7fb", nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: fmt.Sprintf(`{"id":"845ac772-cb49-433c-a871-0a98af34f7fb"}` + "\n"),
		},
		{
			name:      "Bad request",
			inputBody: `{"username": ""}`,
			input: args{
				&models.Users{
					Username: "",
				},
			},
			mock: func(r *mock_usecase.MockUser, args args) {
				r.EXPECT().CreateUserUseCase(args.user).Return("", mError.ErrUserInvalid)
			},
			expectedStatusCode:   400,
			expectedResponseBody: fmt.Sprintf(`{"message":"Bad request"}` + "\n"),
		},
		{
			name:      "The repository is not available",
			inputBody: `{"username": "user_1"}`,
			input: args{
				&models.Users{
					Username: "user_1",
				},
			},
			mock: func(r *mock_usecase.MockUser, args args) {
				r.EXPECT().CreateUserUseCase(args.user).Return("", mError.ErrDB)
			},
			expectedStatusCode:   500,
			expectedResponseBody: fmt.Sprintf(`{"message":"Internal server error"}` + "\n"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_usecase.NewMockUser(c)
			test.mock(repo, test.input)

			useCase := &usecase.UseCase{User: repo}
			handler := Handler{useCase}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/users/add",
				bytes.NewBufferString(test.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			app.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
