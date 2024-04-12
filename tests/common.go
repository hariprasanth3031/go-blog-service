package tests

import (
	"context"

	"github.com/hariprasanth3031/go-blog-service/tests/mock_db"
	"github.com/stretchr/testify/mock"
)

// Cache
var MockBlogsDb = new(mock_db.MockBlogsDb)

func InitializeMockFunctions(functions []func() *mock.Call) []*mock.Call {
	var mockCalls = make([]*mock.Call, 0)
	for _, function := range functions {
		mockCall := function()
		mockCalls = append(mockCalls, mockCall)
	}
	return mockCalls
}

func UnsetMockCalls(functions []*mock.Call) {
	for _, call := range functions {
		call.Unset()
	}
}

func InitializeContext() context.Context {
	var ctx = context.Background()
	ctx = context.WithValue(context.Background(), "email", "jay@tekion.com")
	ctx = context.WithValue(ctx, "accessemail", "jay@tekion.com")
	return ctx
}
