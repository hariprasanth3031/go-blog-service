package test_cases

import (
	"github.com/hariprasanth3031/go-blog-service/models"
	"github.com/hariprasanth3031/go-blog-service/resources"
	"github.com/stretchr/testify/mock"

	//"net/http"
	"github.com/hariprasanth3031/go-blog-service/tests"
)

var CreatePostTestcases = []resources.TestCreatePost{
	{
		Case:  "Create Post",
		Input: models.Blog{},
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("CreatePost", mock.Anything, mock.Anything).Return(models.Blog{}, nil)
			},
		},
		ExpectedOutputParam1: models.Blog{},
		ExpectedOutputParam2: nil,
	},
}

var GetPostTestcases = []resources.TestCreatePost{
	{
		Case:  "Create Post",
		Input: models.Blog{},
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("CreatePost", mock.Anything, mock.Anything).Return(models.Blog{}, nil)
			},
		},
		ExpectedOutputParam1: models.Blog{},
		ExpectedOutputParam2: nil,
	},
}

var UpdatePostTestcases = []resources.TestUpdatePost{
	{
		Case:  "Create Post",
		Input: models.Blog{},
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("UpdatePost", mock.Anything, mock.Anything).Return(models.Blog{}, nil)
			},
		},
		ExpectedOutput: nil,
	},
}

var DeletePostTestcases = []resources.TestDeletePost{
	{
		Case:  "Delete Post",
		Input: 1,
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("DeletePost", mock.Anything, uint64(1)).Return(nil)
			},
		},
		ExpectedOutput: nil,
	},
}
