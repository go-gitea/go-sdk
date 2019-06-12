package gitea

import (
	"fmt"
	"os"
	"testing"
)

func TestGiteaUserCreate(t *testing.T) {
	client := NewClient(os.Getenv("GITEA_HOST"), os.Getenv("GITEA_SECRET"))
	user, err := client.AdminCreateUser(CreateUserOption{
		Email:      "random@example.com",
		Username:   "random",
		Password:   "password",
		SendNotify: false,
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", user)
}
