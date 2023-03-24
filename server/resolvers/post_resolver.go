package resolvers

import (
	"context"
	"fmt"

	"github.com/Vulpe-Fox/Vulpefox-Website/graph/model"
)

func CreateTextPostResolver(ctx context.Context, input *model.TextPostInput) (*model.Post, error) {
	// example:
	// authenticate user ->
	// 	isModerator(ctx) (boolean, error)
	// if(error != nil) return nil, error
	// if(!boolean) return nil, errors.new: not authorized
	panic(fmt.Errorf("not implemented: CreateTextPost - createTextPost"))
}

func CreateImagePostResolver(ctx context.Context, input *model.ImagePostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreateImagePost - createImagePost"))
}

func CreateSurveyPostResolver(ctx context.Context, input *model.SurveyPostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreateSurveyPost - createSurveyPost"))
}

func CreateCommentResolver(ctx context.Context, input *model.CommentInput) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}
