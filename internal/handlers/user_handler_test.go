package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mErr "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/usecase/mocks"
)

func TestCreateUser(t *testing.T) {
	testUser := &models.Users{
		ID:       uuid.New(),
		Username: "user_1",
	}

	tests := []struct {
		testName     string
		expectations func(ctx context.Context, svc *mocks.User)
		inputBody    string
		responseBody string
		err          error
		code         int
	}{
		{
			testName: "valid",
			expectations: func(ctx context.Context, svc *mocks.User) {
				svc.On("CreateUserUseCase", testUser).Return(testUser.ID, nil)
			},
			inputBody: `{"username": "user_1"}`,
			code:      http.StatusCreated,
		},
		{
			testName:     "missing parameter",
			expectations: func(ctx context.Context, svc *mocks.User) {},
			inputBody:    `{}`,
			err:          errors.New("invalid username"),
			code:         http.StatusBadRequest,
		},
		{
			testName:     "bad request",
			expectations: func(ctx context.Context, svc *mocks.User) {},
			inputBody:    `{some"}`,
			err:          errors.New("invalid username"),
			code:         http.StatusBadRequest,
		},
		{
			testName: "service error",
			expectations: func(ctx context.Context, svc *mocks.User) {
				svc.On("CreateUserUseCase", ctx, testUser).Return(nil, mErr.ErrDB)
			},
			inputBody: `{"username": "user_1"}`,
			err:       errors.New("can't create user: database error"),
			code:      http.StatusInternalServerError,
		},
		{
			testName: "already used",
			expectations: func(ctx context.Context, svc *mocks.User) {
				svc.On("CreateUserUseCase", ctx, testUser).Return(nil, mErr.ErrDB)
			},
			inputBody: `{"username": "user_1"}`,
			err:       errors.New("username already used"),
			code:      http.StatusConflict,
		},
	}

	for _, test := range tests {
		t.Logf("running %v", test.testName)

		// initialize the echo context to use for the test
		e := echo.New()
		r, err := http.NewRequest(echo.POST, "/users/add", strings.NewReader(test.inputBody))
		if err != nil {
			t.Fatal("could not create request")
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)
		svc := &mocks.User{}
		test.expectations(ctx.Request().Context(), svc)
		assert.Equal(t, w.Code, test.code)
		assert.Equal(t, w.Body.String(), test.expectations)
	}
}
