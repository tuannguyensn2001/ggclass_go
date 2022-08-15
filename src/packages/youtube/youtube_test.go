package youtubepkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIdFromLink(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		result, err := GetIdFromLink("https://www.youtube.com/watch?v=dShq3wqgzII")

		assert.Nil(t, err)
		assert.Equal(t, "dShq3wqgzII", result)
	})

	t.Run("case 2", func(t *testing.T) {
		result, err := GetIdFromLink("https://www.youtube.com/watch?v=_PQOELbHTIA&list=RD_PQOELbHTIA&start_radio=1")

		assert.Nil(t, err)
		assert.Equal(t, "_PQOELbHTIA", result)
	})
}

func TestGetLinkThumbnailFromId(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		result := GetLinkThumbnailFromId("_PQOELbHTIA")
		assert.Equal(t, "https://img.youtube.com/vi/_PQOELbHTIA/0.jpg", result)
	})
}
