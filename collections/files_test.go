package collections_test

import (
	"testing"

	"github.com/cadigun/goeasyclient/collections"
	"github.com/stretchr/testify/assert"
)

func Test_CopyFileFromURL(t *testing.T) {
	testCases := map[string]struct {
		url      string
		filename string
		expected error
	}{
		"test_case_1": {
			url:      "https://pixabay.com/vectors/watermelon-smiley-emoji-face-happy-5877895/",
			filename: "watermelon.png",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := collections.CopyFileFromURL(tc.url, tc.filename)
			assert.Equal(t, tc.expected, err)
		})
	}
}
