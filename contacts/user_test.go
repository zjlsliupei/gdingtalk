package contacts

import (
	"fmt"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	user := User{AccessToken: "2530c6e4f7f93d44a9495ace12c80d7c"}
	fmt.Println(user.GetAllUsers())
}
