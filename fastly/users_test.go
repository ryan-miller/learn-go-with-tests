package fastly

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

const testUserName string = "Test User"
const testUserLogin string = "test-user-go@golang.com"
const testUserRole string = "user"
const keepUserLogin string = "ryan.l.miller@icloud.com"
const keepUserID string = "29lmRy49CoXwO00ATOzOvj"

func TestUserAPI(t *testing.T) {

	t.Run("Create user and delete user", func(t *testing.T) {
		u := User{}
		u.Name = testUserName
		u.Login = testUserLogin
		u.Role = testUserRole

		createdUser := CreateUser(&u)

		if strings.Compare(createdUser.Login, u.Login) != 0 {
			fmt.Printf("%s was not created.", testUserName)
		}

		deletedUserStatusCode := DeleteUser(createdUser.ID)
		want := http.StatusOK

		if deletedUserStatusCode != want {
			fmt.Printf("want %d, got %d", want, deletedUserStatusCode)
		}
	})

	t.Run("Get current user", func(t *testing.T) {
		got := GetCurrentUser()

		if strings.Compare(got.ID, keepUserID) != 0 {
			t.Errorf("got %s instead of '%s'", got.ID, keepUserID)
		}
	})

}
