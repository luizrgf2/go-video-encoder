package entities_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/luizrgf2/go-video-encoder/internal/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := entities.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	job, err := entities.NewJob(video, "path", "converted")
	require.NotNil(t, job)
	require.Nil(t, err)
	require.Equal(t, video.ID, job.VideoId)
	require.Equal(t, "converted", job.Status)
	require.Equal(t, "path", job.OutputBucketPath)
	require.NotEmpty(t, job.ID)
}
