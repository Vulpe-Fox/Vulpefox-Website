package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"

	"github.com/Vulpe-Fox/Vulpefox-Website/graph/model"
	"github.com/Vulpe-Fox/Vulpefox-Website/resolvers"
)

// CreateTextPost is the resolver for the createTextPost field.
func (r *mutationResolver) CreateTextPost(ctx context.Context, input *model.TextPostInput) (*model.Post, error) {
	return resolvers.CreateTextPostResolver(ctx, input)
}

// CreateImagePost is the resolver for the createImagePost field.
func (r *mutationResolver) CreateImagePost(ctx context.Context, input *model.ImagePostInput) (*model.Post, error) {
	return resolvers.CreateImagePostResolver(ctx, input)
}

// CreateSurveyPost is the resolver for the createSurveyPost field.
func (r *mutationResolver) CreateSurveyPost(ctx context.Context, input *model.SurveyPostInput) (*model.Post, error) {
	return resolvers.CreateSurveyPostResolver(ctx, input)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input *model.CommentInput) (*model.Comment, error) {
	return resolvers.CreateCommentResolver(ctx, input)
}

// QueryPosts is the resolver for the queryPosts field.
func (r *mutationResolver) QueryPosts(ctx context.Context, input *model.QueryInput) (*model.PostQuery, error) {
	return resolvers.QueryPostsResolver(ctx, input)
}

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input *model.CreateAccountInput) (*model.User, error) {
	return resolvers.CreateAccountResolver(ctx, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.LoginInput) (*model.User, error) {
	return resolvers.LoginResolver(ctx, input)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }