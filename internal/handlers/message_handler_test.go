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
	mockUsecase "github.com/olteffe/avitochat/internal/usecase/mocks"
)

func TestHandler_GetMessagesHandler(t *testing.T) {
	type args struct {
		message *models.Messages
	}
	type mockBehavior func(r *mockUsecase.MockMessage, args args)
	chat1Uuid := uuid.New().String()
	message1Uuid := uuid.New()
	message2Uuid := uuid.New()
	user1Uuid := uuid.New().String()
	user2Uuid := uuid.New().String()
	time2message := time.Now().Round(time.Microsecond)
	time1message := time.Now().Add(time.Second * 5).Round(time.Microsecond)

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
			inputBody: fmt.Sprintf(`{"chat": "%s"}`, chat1Uuid),
			input: args{
				&models.Messages{
					Chat: chat1Uuid,
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				data := []*models.Messages{
					{
						ID:        message1Uuid,
						Chat:      chat1Uuid,
						Author:    user1Uuid,
						Text:      "first message",
						CreatedAt: time1message,
					},
					{
						ID:        message2Uuid,
						Chat:      chat1Uuid,
						Author:    user2Uuid,
						Text:      "second message",
						CreatedAt: time2message,
					},
				}
				r.EXPECT().GetMessagesUseCase(args.message).Return(data, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: fmt.Sprintf(
				`[{"id": "%s", "chat": "%s", "author": "%s", "text": "first message", "created_at": "%s"}, 
				{"id": "%s", "chat": "%s", "author": "%s", "text": "second message", "created_at": "%s"}]`,
				message1Uuid, chat1Uuid, user1Uuid, time1message.Format(time.RFC3339Nano),
				message2Uuid, chat1Uuid, user2Uuid, time2message.Format(time.RFC3339Nano)),
		},
		{
			name:      "invalid input data",
			inputBody: `{"chat": ""}`,
			input: args{
				&models.Messages{
					Chat: "",
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				r.EXPECT().GetMessagesUseCase(args.message).Return(nil, mError.ErrChatIdInvalid)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message": "Bad request"}`,
		},
		{
			name:      "chat id not found in the database",
			inputBody: fmt.Sprintf(`{"chat": "%s"}`, chat1Uuid),
			input: args{
				&models.Messages{
					Chat: chat1Uuid,
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				r.EXPECT().GetMessagesUseCase(args.message).Return(nil, mError.ErrUserOrChat)
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message": "Chat not found"}`,
		},
		{
			name:      "database error",
			inputBody: fmt.Sprintf(`{"chat": "%s"}`, chat1Uuid),
			input: args{
				&models.Messages{
					Chat: chat1Uuid,
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				r.EXPECT().GetMessagesUseCase(args.message).Return(nil, mError.ErrDB)
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message": "Internal server error"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockUsecase.NewMockMessage(c)
			test.mock(repo, test.input)

			useCase := &usecase.UseCase{Message: repo}
			handler := Handler{useCase}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/messages/get",
				bytes.NewBufferString(test.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			app.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_SendMessageHandler(t *testing.T) {
	type args struct {
		message *models.Messages
	}
	type mockBehavior func(r *mockUsecase.MockMessage, args args)
	chat1Uuid := uuid.New().String()
	message1Uuid := uuid.New().String()
	user1Uuid := uuid.New().String()

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
			inputBody: fmt.Sprintf(`{"chat": "%s", "author": "%s", "text": "first message"}`, chat1Uuid, user1Uuid),
			input: args{
				message: &models.Messages{
					Chat:   chat1Uuid,
					Author: user1Uuid,
					Text:   "first message",
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				r.EXPECT().SendMessageUseCase(args.message).Return(message1Uuid, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: fmt.Sprintf(`{"id": "%s"}`, message1Uuid),
		},
		{
			name:      "input data is invalid",
			inputBody: fmt.Sprintf(`{"chat": "", "author": "%s", "text": "first message"}`, user1Uuid),
			input: args{
				message: &models.Messages{
					Chat:   "",
					Author: user1Uuid,
					Text:   "first message",
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				r.EXPECT().SendMessageUseCase(args.message).Return("", mError.ErrChatIdInvalid)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message": "Bad request"}`,
		},
		{
			name:      "chat/user ID not in db",
			inputBody: fmt.Sprintf(`{"chat": "%s", "author": "%s", "text": "first message"}`, chat1Uuid, user1Uuid),
			input: args{
				message: &models.Messages{
					Chat:   chat1Uuid,
					Author: user1Uuid,
					Text:   "first message",
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				r.EXPECT().SendMessageUseCase(args.message).Return("", mError.ErrUserOrChat)
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message": "User or chat not found"}`,
		},
		{
			name:      "db error",
			inputBody: fmt.Sprintf(`{"chat": "%s", "author": "%s", "text": "first message"}`, chat1Uuid, user1Uuid),
			input: args{
				message: &models.Messages{
					Chat:   chat1Uuid,
					Author: user1Uuid,
					Text:   "first message",
				},
			},
			mock: func(r *mockUsecase.MockMessage, args args) {
				r.EXPECT().SendMessageUseCase(args.message).Return("", mError.ErrDB)
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message": "Internal server error"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockUsecase.NewMockMessage(c)
			test.mock(repo, test.input)

			useCase := &usecase.UseCase{Message: repo}
			handler := Handler{useCase}

			app := echo.New()
			handler.Init(app)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/messages/add",
				bytes.NewBufferString(test.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			app.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
