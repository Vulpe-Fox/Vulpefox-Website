package resolvers

import (
	"context"
	"fmt"

	"github.com/Vulpe-Fox/Vulpefox-Website/graph/model"
)

func CreateAccountResolver(ctx context.Context, input *model.CreateAccountInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateAccount - createAccount"))
}

func LoginResolver(ctx context.Context, input *model.LoginInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}
