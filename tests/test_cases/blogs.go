package test_cases

import (
	"errors"

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
				return tests.MockBlogsDb.On("CreatePost", mock.Anything, mock.Anything).Return(&models.Blog{}, nil)
			},
		},
		ExpectedOutputParam1: &models.Blog{},
		ExpectedOutputParam2: nil,
	},
}

var GetPostTestcases = []resources.TestGetPost{
	{
		Case:  "Get Post(Success)",
		Input: 1,
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				blog := models.Blog{

					Title:   "Mock Title",
					Content: "Mock Content",
					Author:  "Mock Author",
					Tags:    "tag2",
				}
				return tests.MockBlogsDb.On("GetPost", mock.Anything, mock.Anything).Return(&blog, nil)
			},
		},
		ExpectedOutputParam1: &models.Blog{
			Title:   "Mock Title",
			Content: "Mock Content",
			Author:  "Mock Author",
			Tags:    "tag2",
		},
		ExpectedOutputParam2: nil,
	},
	{
		Case:  "Get Post(Error)",
		Input: 1,
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("CreatePost", mock.Anything, mock.Anything).Return(&models.Blog{}, errors.New("error"))
			},
		},
		ExpectedOutputParam1: &models.Blog{},
		ExpectedOutputParam2: errors.New("error"),
	},
}

var UpdatePostTestcases = []resources.TestUpdatePost{
	{
		Case:  "Update Post(Success)",
		Input: models.Blog{},
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("UpdatePost", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		ExpectedOutput: nil,
	},
	{
		Case:  "Update Post(Failure)",
		Input: models.Blog{},
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("UpdatePost", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
			},
		},
		ExpectedOutput: errors.New("error"),
	},
}

var DeletePostTestcases = []resources.TestDeletePost{
	{
		Case:  "Delete Post(Success)",
		Input: 1,
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("DeletePost", mock.Anything, uint64(1)).Return(nil)
			},
		},
		ExpectedOutput: nil,
	},
	{
		Case:  "Delete Post(Failure)",
		Input: 1,
		MockFunctions: []func() *mock.Call{

			func() *mock.Call {
				return tests.MockBlogsDb.On("DeletePost", mock.Anything, uint64(1)).Return(nil)
			},
		},
		ExpectedOutput: nil,
	},
}
