package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	store := &StubPlayerStore{}
	cli := &CLI{store}
	cli.PlayPoker()
	assert.Equal(t, 1, len(store.winCalls))

}
