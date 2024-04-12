package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hariprasanth3031/go-blog-service/db"
	"github.com/hariprasanth3031/go-blog-service/tests"
	"github.com/hariprasanth3031/go-blog-service/tests/test_cases"
	"github.com/stretchr/testify/assert"
)

var (
	TestBlogService BlogService
	MockBlogsDb     db.BlogDb
	ctx             = context.Background()
)

func TestMain(m *testing.M) {
	MockBlogsDb = tests.MockBlogsDb
	TestBlogService = NewBlogService(MockBlogsDb)

	fmt.Println("----------Started test-cases for blog service----------")
	code := m.Run()
	fmt.Println("----------Completed test-cases for blog service----------")
	os.Exit(code)
}

func TestCreatePost(t *testing.T) {
	for _, test := range test_cases.CreatePostTestcases {
		t.Run(test.Case, func(t *testing.T) {

			mockCalls := tests.InitializeMockFunctions(test.MockFunctions)

			param1, param2 := TestBlogService.CreatePost(ctx, test.Input)
			//expectedResult := test.ExpectedOutput
			assert.Equal(t, param1, test.ExpectedOutputParam1, "actual result and expected result are not equal")
			assert.Equal(t, param2, test.ExpectedOutputParam2, "actual result and expected result are not equal")
			tests.UnsetMockCalls(mockCalls)
		})
	}
}

func TestDeletePost(t *testing.T) {
	for _, test := range test_cases.DeletePostTestcases {
		t.Run(test.Case, func(t *testing.T) {

			mockCalls := tests.InitializeMockFunctions(test.MockFunctions)

			param1 := TestBlogService.DeletePost(ctx, test.Input)
			assert.Equal(t, param1, test.ExpectedOutput, "actual result and expected result are not equal")
			tests.UnsetMockCalls(mockCalls)
		})
	}
}
