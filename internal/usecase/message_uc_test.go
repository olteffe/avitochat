package usecase

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mErr "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	mockRepo "github.com/olteffe/avitochat/internal/repository/mocks"
)

func TestMessageUseCase_GetMessagesUseCase(t *testing.T) {
	t.Parallel()
	uuidChat1 := uuid.New().String()
	uuidUser1 := uuid.New().String()
	uuidUser2 := uuid.New().String()
	uuidMessage1 := uuid.New()
	uuidMessage2 := uuid.New()
	time1 := time.Now().Round(time.Microsecond)
	time2 := time.Now().Add(time.Second * 5).Round(time.Microsecond)
	type args struct {
		message *models.Messages
	}

	type mockBehavior func(r *mockRepo.MockMessage, args args)

	tests := []struct {
		name        string
		input       args
		messageMock mockBehavior
		mock        mockBehavior
		wantErr     bool
		want        []*models.Messages
	}{
		{
			name: "valid",
			input: args{
				message: &models.Messages{
					Chat: uuidChat1,
				},
			},
			messageMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceChat(args.message).Return(nil)
			},
			mock: func(r *mockRepo.MockMessage, args args) {
				data := []*models.Messages{
					{
						ID:        uuidMessage2,
						Chat:      uuidChat1,
						Author:    uuidUser2,
						Text:      "message2text",
						CreatedAt: time2,
					},
					{
						ID:        uuidMessage1,
						Chat:      uuidChat1,
						Author:    uuidUser1,
						Text:      "message1text",
						CreatedAt: time1,
					},
				}
				r.EXPECT().GetMessagesRepository(args.message).Return(data, nil)
			},
			wantErr: false,
			want: []*models.Messages{
				{
					ID:        uuidMessage2,
					Chat:      uuidChat1,
					Author:    uuidUser2,
					Text:      "message2text",
					CreatedAt: time2,
				},
				{
					ID:        uuidMessage1,
					Chat:      uuidChat1,
					Author:    uuidUser1,
					Text:      "message1text",
					CreatedAt: time1,
				},
			},
		},
		{
			name: "incorrect chat ID",
			input: args{
				message: &models.Messages{
					Chat: "uuidChat1",
				},
			},
			messageMock: func(r *mockRepo.MockMessage, args args) {},
			mock:        func(r *mockRepo.MockMessage, args args) {},
			wantErr:     true,
		},
		{
			name: "chat ID not found",
			input: args{
				message: &models.Messages{
					Chat: uuidChat1,
				},
			},
			messageMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceChat(args.message).Return(mErr.ErrUserOrChat)
			},
			mock:    func(r *mockRepo.MockMessage, args args) {},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepo.NewMockMessage(c)
			test.messageMock(repo, test.input)
			test.mock(repo, test.input)
			s := &MessageUseCase{repo: repo}

			got, err := s.GetMessagesUseCase(test.input.message)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestMessageUseCase_SendMessageUseCase(t *testing.T) {
	t.Parallel()
	uuidChat1 := uuid.New().String()
	uuidUser1 := uuid.New().String()
	uuidMessage1 := uuid.New().String()
	type args struct {
		message *models.Messages
	}

	type mockBehavior func(r *mockRepo.MockMessage, args args)

	tests := []struct {
		name       string
		input      args
		chatMock   mockBehavior
		authorMock mockBehavior
		mock       mockBehavior
		wantErr    bool
		want       string
	}{
		{
			name: "valid",
			input: args{
				message: &models.Messages{
					Chat:   uuidChat1,
					Author: uuidUser1,
					Text:   "chat1_message1",
				},
			},
			chatMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceChat(args.message).Return(nil)
			},
			authorMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceAuthor(args.message).Return(nil)
			},
			mock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().SendMessageRepository(args.message).Return(uuidMessage1, nil)
			},
			wantErr: false,
			want:    uuidMessage1,
		},
		{
			name: "wrong chat ID",
			input: args{
				message: &models.Messages{
					Chat:   "uuidChat1",
					Author: uuidUser1,
					Text:   "chat1_message1",
				},
			},
			chatMock:   func(r *mockRepo.MockMessage, args args) {},
			authorMock: func(r *mockRepo.MockMessage, args args) {},
			mock:       func(r *mockRepo.MockMessage, args args) {},
			wantErr:    true,
		},
		{
			name: "wrong author ID",
			input: args{
				message: &models.Messages{
					Chat:   uuidChat1,
					Author: "uuidUser1",
					Text:   "chat1_message1",
				},
			},
			chatMock:   func(r *mockRepo.MockMessage, args args) {},
			authorMock: func(r *mockRepo.MockMessage, args args) {},
			mock:       func(r *mockRepo.MockMessage, args args) {},
			wantErr:    true,
		},
		{
			name: "chat ID not found",
			input: args{
				message: &models.Messages{
					Chat:   uuidChat1,
					Author: uuidUser1,
					Text:   "chat1_message1",
				},
			},
			chatMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceChat(args.message).Return(mErr.ErrUserOrChat)
			},
			authorMock: func(r *mockRepo.MockMessage, args args) {},
			mock:       func(r *mockRepo.MockMessage, args args) {},
			wantErr:    true,
		},
		{
			name: "user ID not found",
			input: args{
				message: &models.Messages{
					Chat:   uuidChat1,
					Author: uuidUser1,
					Text:   "chat1_message1",
				},
			},
			chatMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceChat(args.message).Return(nil)
			},
			authorMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceAuthor(args.message).Return(mErr.ErrUserOrChat)
			},
			mock:    func(r *mockRepo.MockMessage, args args) {},
			wantErr: true,
		},
		{
			name: "db error",
			input: args{
				message: &models.Messages{
					Chat:   uuidChat1,
					Author: uuidUser1,
					Text:   "chat1_message1",
				},
			},
			chatMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceChat(args.message).Return(nil)
			},
			authorMock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().ExistenceAuthor(args.message).Return(nil)
			},
			mock: func(r *mockRepo.MockMessage, args args) {
				r.EXPECT().SendMessageRepository(args.message).Return("", mErr.ErrDB)
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepo.NewMockMessage(c)
			test.chatMock(repo, test.input)
			test.authorMock(repo, test.input)
			test.mock(repo, test.input)
			s := &MessageUseCase{repo: repo}

			got, err := s.SendMessageUseCase(test.input.message)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
