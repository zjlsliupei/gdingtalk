package contacts

import (
	"fmt"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	user := NewUser("c11951e3b80f3ab2a9ab7929dc170f5e")
	fmt.Println(user.GetAllUsers())
}
