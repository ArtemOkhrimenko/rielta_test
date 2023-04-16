package app_test

import (
	"github.com/stretchr/testify/assert"
	"rielta/intertal/app"
	"testing"
)

func TestApp_GetNumber(t *testing.T) {
	testCase := map[string]struct {
		wantCount int64
	}{
		"success": {wantCount: 0},
	}

	for name, tc := range testCase {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			module := app.New()

			count := module.GetNumber()
			assert.Equal(t, tc.wantCount, count)

		})
	}
}

func TestApp_GetUser(t *testing.T) {
	testCase := map[string]struct {
		count     int64
		req       int64
		wantCount int64
	}{
		"success":         {count: 4, req: 2, wantCount: 6},
		"negative_number": {count: 4, req: -2, wantCount: 2},
	}

	for name, tc := range testCase {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			module := app.New()
			app.Number = tc.count

			count := module.PutNumber(tc.req)
			assert.Equal(t, tc.wantCount, count)
		})
	}
}
