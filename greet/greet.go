package greet

import (
	"fmt"

	v1 "github.com/yuki/api/go/greet/v1"
)

func Greeter(user *v1.Greet) (result string) {
	userInfo := user.GetUser()
	if userInfo != nil {
		result = fmt.Sprintf(
			"Hello, %s %s",
			userInfo.Gender.String(),
			userInfo.GetLastName(),
		)
	}
	return
}
