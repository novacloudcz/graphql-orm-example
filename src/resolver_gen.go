package src

import (
	"context"

	"github.com/novacloudcz/graphql-orm-example/gen"
)

func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type CompanyResultTypeResolver struct {
	*gen.GeneratedCompanyResultTypeResolver
}

func (r *Resolver) CompanyResultType() gen.CompanyResultTypeResolver {
	return &CompanyResultTypeResolver{&gen.GeneratedCompanyResultTypeResolver{r.GeneratedResolver}}
}

type CompanyResolver struct {
	*gen.GeneratedCompanyResolver
}

func (r *Resolver) Company() gen.CompanyResolver {
	return &CompanyResolver{&gen.GeneratedCompanyResolver{r.GeneratedResolver}}
}

type UserResultTypeResolver struct {
	*gen.GeneratedUserResultTypeResolver
}

func (r *Resolver) UserResultType() gen.UserResultTypeResolver {
	return &UserResultTypeResolver{&gen.GeneratedUserResultTypeResolver{r.GeneratedResolver}}
}

type UserResolver struct {
	*gen.GeneratedUserResolver
}

func (r *Resolver) User() gen.UserResolver {
	return &UserResolver{&gen.GeneratedUserResolver{r.GeneratedResolver}}
}

type TaskResultTypeResolver struct {
	*gen.GeneratedTaskResultTypeResolver
}

func (r *Resolver) TaskResultType() gen.TaskResultTypeResolver {
	return &TaskResultTypeResolver{&gen.GeneratedTaskResultTypeResolver{r.GeneratedResolver}}
}

type TaskResolver struct {
	*gen.GeneratedTaskResolver
}

func (r *Resolver) Task() gen.TaskResolver {
	return &TaskResolver{&gen.GeneratedTaskResolver{r.GeneratedResolver}}
}
