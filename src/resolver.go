package src

import (
	"context"
	"fmt"

	"github.com/novacloudcz/graphql-orm-example/gen"
	"github.com/novacloudcz/graphql-orm/events"

	"github.com/badoux/checkmail"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedResolver, input map[string]interface{}) (item *gen.User, err error) {
		if email, ok := input["email"].(string); ok {
			validationError := checkmail.ValidateFormat(email)
			if validationError != nil {
				err = fmt.Errorf("Email error: %s", validationError.Error())
				return
			}
		}
		return gen.CreateUserHandler(ctx, r, input)
	}

	return resolver
}

func (*QueryResolver) Hello(ctx context.Context) (string, error) {
	return "world", nil
}
