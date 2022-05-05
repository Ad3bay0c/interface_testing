package handlers

import (
	"fmt"
	"github.com/Ad3bay0c/interface_testing/database/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPostgresDB := mocks.NewMockDB(ctrl)

	h := CreateHandler(mockPostgresDB)
	mockPostgresDB.EXPECT().ListUsers().Return([]string{"Joseph", "Biden"}, nil)
	resp, err := h.ListUsers()
	if err != nil {
		t.Fatal(err)
	}
	if len(resp) < 1 {
		t.Fatalf("expected %v, got %v", 1, len(resp))
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedDB := mocks.NewMockDB(ctrl)

	h := CreateHandler(mockedDB)
	t.Run("return_error", func(t *testing.T) {
		mockedDB.EXPECT().CreateUser("Cool boy").
			Return("", fmt.Errorf("an error occured"))


		user, err := h.CreateUser("Cool boy")
		if err == nil {
			t.Fatal(err)
		}
		if len(user) > 0 {
			t.Fatal("Expected an empty value")
		}
	})
	t.Run("successful", func(t *testing.T) {
		mockedDB.EXPECT().CreateUser("Cool boy").
			Return("John ", nil)

		user, err := h.CreateUser("Cool boy")
		if err != nil {
			t.Fatal(err)
		}
		if len(user) == 0 {
			t.Fatalf("Expected %v, got %v", len(user), 0)
		}
	})
}
