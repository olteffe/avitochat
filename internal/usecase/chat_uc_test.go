package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mError "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	mockRepo "github.com/olteffe/avitochat/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChatUseCase_CreateChatUseCase(t *testing.T) {
	t.Parallel()
	uuidUser1 := uuid.New().String()
	uuidUser2 := uuid.New().String()
	uuidChat1 := uuid.New().String()
	type args struct {
		chat *models.Chats
	}

	type chatMockBehavior func(r *mockRepo.MockChat, args args)
	type mockBehavior func(r *mockRepo.MockChat, args args)

	tests := []struct {
		name     string
		input    args
		chatMock chatMockBehavior
		mock     mockBehavior
		wantErr  bool
		want     string
	}{
		{
			name: "Valid",
			input: args{
				chat: &models.Chats{
					Name: "Chat_1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			chatMock: func(r *mockRepo.MockChat, args args) {
				r.EXPECT().ExistenceChatName(args.chat).Return(nil)
			},
			mock: func(r *mockRepo.MockChat, args args) {
				r.EXPECT().CreateChatRepository(args.chat).Return(uuidChat1, nil)
			},
			wantErr: false,
			want:    uuidChat1,
		},
		{
			name: "skipped chat name",
			input: args{
				chat: &models.Chats{
					Name: "",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			chatMock: func(r *mockRepo.MockChat, args args) {},
			mock:     func(r *mockRepo.MockChat, args args) {},
			wantErr:  true,
		},
		{
			name: "users ID is invalid",
			input: args{
				chat: &models.Chats{
					Name: "Chat_1",
					Users: []string{
						uuidUser1,
						"",
					},
				},
			},
			chatMock: func(r *mockRepo.MockChat, args args) {},
			mock:     func(r *mockRepo.MockChat, args args) {},
			wantErr:  true,
		},
		{
			name: "user ID not found",
			input: args{
				chat: &models.Chats{
					Name: "Chat_1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			chatMock: func(r *mockRepo.MockChat, args args) {
				r.EXPECT().ExistenceChatName(args.chat).Return(mError.ErrUserInvalid)
			},
			mock:    func(r *mockRepo.MockChat, args args) {},
			wantErr: true,
		},
		{
			name: "chat name already used",
			input: args{
				chat: &models.Chats{
					Name: "Chat_1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			chatMock: func(r *mockRepo.MockChat, args args) {
				r.EXPECT().ExistenceChatName(args.chat).Return(mError.ErrChatInvalid)
			},
			mock:    func(r *mockRepo.MockChat, args args) {},
			wantErr: true,
		},
		{
			name: "DB error",
			input: args{
				chat: &models.Chats{
					Name: "Chat_1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
				},
			},
			chatMock: func(r *mockRepo.MockChat, args args) {
				r.EXPECT().ExistenceChatName(args.chat).Return(nil)
			},
			mock: func(r *mockRepo.MockChat, args args) {
				r.EXPECT().CreateChatRepository(args.chat).Return("", mError.ErrDB)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepo.NewMockChat(c)
			test.chatMock(repo, test.input)
			test.mock(repo, test.input)
			s := &ChatUseCase{repo: repo}

			got, err := s.CreateChatUseCase(test.input.chat)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestChatUseCase_GetChatUseCase(t *testing.T) {
	t.Parallel()
	uuidUser1 := uuid.New().String()
	uuidUser2 := uuid.New().String()
	uuidUser3 := uuid.New().String()
	uuidChat1 := uuid.New()
	uuidChat2 := uuid.New()
	time1 := time.Now().Round(time.Microsecond)
	time2 := time.Now().Add(time.Second * 5).Round(time.Microsecond)
	type args struct {
		user string
	}

	type userMockBehavior func(r *mockRepo.MockChat, userId string)
	type mockBehavior func(r *mockRepo.MockChat, args args)

	tests := []struct {
		name     string
		input    args
		userMock userMockBehavior
		mock     mockBehavior
		wantErr  bool
		want     []*models.Chats
	}{
		{
			name: "Valid",
			input: args{
				user: uuidUser1,
			},
			userMock: func(r *mockRepo.MockChat, userId string) {
				r.EXPECT().ExistenceUser(userId).Return(nil)
			},
			mock: func(r *mockRepo.MockChat, args args) {
				data := []*models.Chats{
					{
						ID:   uuidChat2,
						Name: "Chat 2",
						Users: []string{
							uuidUser1,
							uuidUser3,
						},
						CreatedAt: time2,
					},
					{
						ID:   uuidChat1,
						Name: "Chat 1",
						Users: []string{
							uuidUser1,
							uuidUser2,
						},
						CreatedAt: time1,
					},
				}
				r.EXPECT().GetChatRepository(args.user).Return(data, nil)
			},
			wantErr: false,
			want: []*models.Chats{
				{
					ID:   uuidChat2,
					Name: "Chat 2",
					Users: []string{
						uuidUser1,
						uuidUser3,
					},
					CreatedAt: time2,
				},
				{
					ID:   uuidChat1,
					Name: "Chat 1",
					Users: []string{
						uuidUser1,
						uuidUser2,
					},
					CreatedAt: time1,
				},
			},
		},
		{
			name: "user not found",
			input: args{
				user: uuidUser1,
			},
			userMock: func(r *mockRepo.MockChat, userId string) {
				r.EXPECT().ExistenceUser(userId).Return(mError.ErrUserIdInvalid)
			},
			mock:    func(r *mockRepo.MockChat, args args) {},
			wantErr: true,
		},
		{
			name: "user ID invalid",
			input: args{
				user: "uuidUser1",
			},
			userMock: func(r *mockRepo.MockChat, userId string) {},
			mock:     func(r *mockRepo.MockChat, args args) {},
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepo.NewMockChat(c)
			test.userMock(repo, test.input.user)
			test.mock(repo, test.input)
			s := &ChatUseCase{repo: repo}

			got, err := s.GetChatUseCase(test.input.user)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
