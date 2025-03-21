package entities_test

import (
	"testing"
	"time"

	"github.com/luizrgf2/go-video-encoder/internal/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := entities.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideNotIsUUID(t *testing.T) {
	video := entities.NewVideo()
	video.ID = "abc"
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := entities.NewVideo()
	video.ID = "9ec84c8a-dd80-4a5c-84d9-c44294d2ab55"
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Nil(t, err)
}
