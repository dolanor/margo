package margo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageMetadata(t *testing.T) {
	r, err := os.Open("testdata/sample1.jpg")
	require.Nil(t, err)

	imd, err := ImageMetadata(r)
	require.Nil(t, err)
	comment := imd.Comment()
	assert.Equal(t, `"taken at basilica of chinese"`, comment)
}
