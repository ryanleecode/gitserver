package git_test

import (
	"fmt"
	"testing"

	"github.com/drdgvhbh/gitserver/internal/git"
	"github.com/stretchr/testify/assert"

	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func TestCommitSummary(t *testing.T) {
	assert := assert.New(t)

	summary := "This is a summary"
	depCommit := object.Commit{
		Message: fmt.Sprintf("%s\n\nThis is a description", summary),
	}

	commit := git.GitCommit{
		Wrapee: &depCommit,
	}

	assert.EqualValues(summary, commit.Summary())
}

func TestCommitHash(t *testing.T) {
	assert := assert.New(t)

	hash := "55245d63089b55144010e408895a4350e163b49e"
	depCommit := object.Commit{
		Hash: plumbing.NewHash(hash),
	}

	commit := git.GitCommit{
		Wrapee: &depCommit,
	}

	assert.EqualValues(hash, commit.Hash())
}
