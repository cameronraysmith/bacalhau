package computenode

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/filecoin-project/bacalhau/pkg/computenode"
	"github.com/filecoin-project/bacalhau/pkg/executor"
	_ "github.com/filecoin-project/bacalhau/pkg/logger"
	"github.com/stretchr/testify/assert"
)

// a simple sanity test of the RunJob with docker executor
func TestRunJob(t *testing.T) {

	EXAMPLE_TEXT := "hello"
	computeNode, ipfsStack, cm := SetupTestDockerIpfs(t, computenode.NewDefaultComputeNodeConfig())
	defer cm.Cleanup()

	cid, err := ipfsStack.AddTextToNodes(1, []byte(EXAMPLE_TEXT))
	assert.NoError(t, err)

	result, err := computeNode.RunJob(context.Background(), &executor.Job{
		ID:   "test",
		Spec: GetJobSpec(cid),
	})
	assert.NoError(t, err)

	stdoutPath := fmt.Sprintf("%s/stdout", result)
	assert.DirExists(t, result, "The job result folder exists")
	assert.FileExists(t, stdoutPath, "The stdout file exists")

	dat, err := os.ReadFile(stdoutPath)
	assert.NoError(t, err)
	assert.Equal(t, EXAMPLE_TEXT, string(dat), "The stdout file contained the correct result from the job")

}

func TestEmptySpec(t *testing.T) {

	computeNode, _, cm := SetupTestDockerIpfs(t, computenode.NewDefaultComputeNodeConfig())
	defer cm.Cleanup()

	// it seems when we return an error so quickly we need to sleep a little bit
	// otherwise we don't cleanup
	// TODO: work out why
	time.Sleep(time.Millisecond * 10)
	_, err := computeNode.RunJob(context.Background(), &executor.Job{
		ID:   "test",
		Spec: nil,
	})
	assert.Error(t, err)
}