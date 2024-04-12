package resources

import (
	"github.com/hariprasanth3031/go-blog-service/models"
	"github.com/stretchr/testify/mock"
)

type TestCreatePost struct {
	Case                 string
	Input                models.Blog
	MockFunctions        []func() *mock.Call
	ExpectedOutputParam1 models.Blog
	ExpectedOutputParam2 error
}

type TestDeletePost struct {
	Case           string
	Input          uint64
	MockFunctions  []func() *mock.Call
	ExpectedOutput error
}

type TestUpdatePost struct {
	Case           string
	Input          models.Blog
	MockFunctions  []func() *mock.Call
	ExpectedOutput error
}

type TestGetPost struct {
	Case                 string
	Input                models.Blog
	MockFunctions        []func() *mock.Call
	ExpectedOutputParam1 models.Blog
	ExpectedOutputParam2 error
}
