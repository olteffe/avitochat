package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mError "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	mockRepo "github.com/olteffe/avitochat/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUseCase_CreateUserUseCase(t *testing.T) {
	t.Parallel()
	uuidUser1 := uuid.New().String()
	type args struct {
		user *models.Users
	}

	type userMockBehavior func(r *mockRepo.MockUser, args args)
	type mockBehavior func(r *mockRepo.MockUser, args args)

	tests := []struct {
		name     string
		input    args
		userMock userMockBehavior
		mock     mockBehavior
		wantErr  bool
		want     string
	}{
		{
			name: "Valid",
			input: args{
				user: &models.Users{
					Username: "User_1",
				},
			},
			userMock: func(r *mockRepo.MockUser, args args) {
				r.EXPECT().ExistenceUser(args.user).Return(nil)
			},
			mock: func(r *mockRepo.MockUser, args args) {
				r.EXPECT().CreateUserRepository(args.user).Return(uuidUser1, nil)
			},
			wantErr: false,
			want:    uuidUser1,
		},
		{
			name: "username is blank",
			input: args{
				user: &models.Users{
					Username: "",
				},
			},
			userMock: func(r *mockRepo.MockUser, args args) {},
			mock:     func(r *mockRepo.MockUser, args args) {},
			wantErr:  true,
		},
		{
			name: "username already exists",
			input: args{
				user: &models.Users{
					Username: "User_1",
				},
			},
			userMock: func(r *mockRepo.MockUser, args args) {
				r.EXPECT().ExistenceUser(args.user).Return(mError.ErrUserAlreadyUsed)
			},
			mock:    func(r *mockRepo.MockUser, args args) {},
			wantErr: true,
		},
		{
			name: "db error",
			input: args{
				user: &models.Users{
					Username: "User_1",
				},
			},
			userMock: func(r *mockRepo.MockUser, args args) {
				r.EXPECT().ExistenceUser(args.user).Return(nil)
			},
			mock: func(r *mockRepo.MockUser, args args) {
				r.EXPECT().CreateUserRepository(args.user).Return("", mError.ErrCantCreateUserDB)
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepo.NewMockUser(c)
			test.userMock(repo, test.input)
			test.mock(repo, test.input)
			s := &UserUseCase{repo: repo}

			got, err := s.CreateUserUseCase(test.input.user)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
