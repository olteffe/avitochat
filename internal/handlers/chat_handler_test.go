package handlers

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mError "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/usecase"
	"github.com/olteffe/avitochat/internal/usecase/mocks"
)

func TestHandler_CreateChatHandler(t *testing.T) {
	uuidUser1 := uuid.New().String()
	uuidUser2 := uuid.New().String()
	uuidChat1 := uuid.New().String()
	type args struct {
		chat *models.Chats
	}
	type mockBehavior func(r *mock_usecase.MockChat, args args)
	tests := []struct {
		name                 string
		inputBody            string
		inputChat            args
		mock                 mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Valid",
			inputBody: fmt.Sprintf(`{"name": "chat_1", "users": ["%s", "%s"]}`, uuidUser1, uuidUser2),
			inputChat: args{
				&models.Chats{
					Name: "chat_1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().CreateChatUseCase(args.chat).Return(uuidChat1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: fmt.Sprintf(`{"id": "%s"}`, uuidChat1),
		},
		{
			name:      "invalid input data",
			inputBody: fmt.Sprintf(`{"name": "chat_1", "users": ["DROP TABLE IF EXISTS users CASCADE", "%s", "%s"]}`, uuidUser1, uuidUser2),
			inputChat: args{
				&models.Chats{
					Name: "chat_1",
					Users: []string{
						"DROP TABLE IF EXISTS users CASCADE",
						uuidUser1,
						uuidUser2,
					},
				},
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().CreateChatUseCase(args.chat).Return("", mError.ErrUserInvalid)
			},
			expectedStatusCode:   400,
			expectedResponseBody: fmt.Sprintf(`{"message":"Bad request"}`),
		},
		{
			name:      "invalid input data: skipped chat name",
			inputBody: fmt.Sprintf(`{"name": "", "users": ["%s", "%s"]}`, uuidUser1, uuidUser2),
			inputChat: args{
				&models.Chats{
					Name: "",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().CreateChatUseCase(args.chat).Return("", mError.ErrChatInvalid)
			},
			expectedStatusCode:   400,
			expectedResponseBody: fmt.Sprintf(`{"message":"Bad request"}`),
		},
		{
			name:      "invalid input data: one user is given",
			inputBody: fmt.Sprintf(`{"name": "chat_1", "users": ["%s"]}`, uuidUser1),
			inputChat: args{
				&models.Chats{
					Name: "chat_1",
					Users: []string{
						uuidUser1,
					},
				},
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().CreateChatUseCase(args.chat).Return("", mError.ErrCountUsers)
			},
			expectedStatusCode:   400,
			expectedResponseBody: fmt.Sprintf(`{"message":"Bad request"}`),
		},
		{
			name:      "the chat name has already been used",
			inputBody: fmt.Sprintf(`{"name": "chat_1", "users": ["%s", "%s"]}`, uuidUser1, uuidUser2),
			inputChat: args{
				&models.Chats{
					Name: "chat_1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().CreateChatUseCase(args.chat).Return("", mError.ErrChatInvalid)
			},
			expectedStatusCode:   400,
			expectedResponseBody: fmt.Sprintf(`{"message": "Bad request"}`),
		},
		{
			name:      "Database error",
			inputBody: fmt.Sprintf(`{"name": "chat_1", "users": ["%s", "%s"]}`, uuidUser1, uuidUser2),
			inputChat: args{
				&models.Chats{
					Name: "chat_1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().CreateChatUseCase(args.chat).Return("", mError.ErrDB)
			},
			expectedStatusCode:   500,
			expectedResponseBody: fmt.Sprintf(`{"message": "Internal server error"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_usecase.NewMockChat(c)
			test.mock(repo, test.inputChat)

			useCase := &usecase.UseCase{Chat: repo}
			handler := Handler{useCase}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/chats/add",
				bytes.NewBufferString(test.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			app.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetChatHandler(t *testing.T) {
	uuidUser1 := uuid.New().String()
	uuidUser2 := uuid.New().String()
	uuidChat1 := uuid.New()
	uuidChat2 := uuid.New()
	time1 := time.Now().Round(time.Microsecond)
	time2 := time.Now().Add(time.Second * 5).Round(time.Microsecond)
	type args struct {
		user string
	}
	type mockBehavior func(r *mock_usecase.MockChat, args args)
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            args
		mock                 mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Valid",
			inputBody: fmt.Sprintf(`{"user": "%s"}`, uuidUser1),
			inputUser: args{
				user: uuidUser1,
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				data := []*models.Chats{
					{
						ID:   uuidChat1,
						Name: "chat_1",
						Users: []string{
							uuidUser1,
							uuidUser2,
						},
						CreatedAt: time2,
					},
					{
						ID:   uuidChat2,
						Name: "chat_2",
						Users: []string{
							uuidUser1,
							uuidUser2,
						},
						CreatedAt: time1,
					},
				}
				r.EXPECT().GetChatUseCase(args.user).Return(data, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: fmt.Sprintf(`[{"id": "%s", "name": "chat_1", "users": ["%s", "%s"], 
				"created_at": "%s"}, {"id": "%s", "name": "chat_2", "users": ["%s", "%s"], "created_at": "%s"}]`,
				uuidChat1, uuidUser1, uuidUser2, time2.Format(time.RFC3339Nano), uuidChat2, uuidUser1, uuidUser2,
				time1.Format(time.RFC3339Nano)),
		},
		{
			name:      "input data is invalid",
			inputBody: fmt.Sprintf(`{"user": ""}`),
			inputUser: args{
				user: "",
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().GetChatUseCase(args.user).Return(nil, mError.ErrChatInvalid)
			},
			expectedStatusCode:   400,
			expectedResponseBody: fmt.Sprintf(`{"message": "Bad request"}`),
		},
		{
			name:      "user ID not found in database",
			inputBody: fmt.Sprintf(`{"user": "%s"}`, uuidUser1),
			inputUser: args{
				user: uuidUser1,
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().GetChatUseCase(args.user).Return(nil, mError.ErrUserIdInvalid)
			},
			expectedStatusCode:   404,
			expectedResponseBody: fmt.Sprintf(`{"message": "User not found"}`),
		},
		{
			name:      "Database error",
			inputBody: fmt.Sprintf(`{"user": "%s"}`, uuidUser1),
			inputUser: args{
				user: uuidUser1,
			},
			mock: func(r *mock_usecase.MockChat, args args) {
				r.EXPECT().GetChatUseCase(args.user).Return(nil, mError.ErrDB)
			},
			expectedStatusCode:   500,
			expectedResponseBody: fmt.Sprintf(`{"message": "Internal server error"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_usecase.NewMockChat(c)
			test.mock(repo, test.inputUser)

			useCase := &usecase.UseCase{Chat: repo}
			handler := Handler{useCase}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/chats/get",
				bytes.NewBufferString(test.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			app.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
