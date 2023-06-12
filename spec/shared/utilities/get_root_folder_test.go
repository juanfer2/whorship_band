package utilities

import (
	"regexp"
	"testing"

	"github.com/juanfer2/whorship_band/src/shared/utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetRootFolder(t *testing.T) {
	assert := assert.New(t)
	folder := utilities.GetRootFolder()

	match, _ := regexp.MatchString("wompi", folder)

	assert.Equal(match, true, "they should be truthy")
}
