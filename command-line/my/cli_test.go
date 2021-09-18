package poker_test

import (
	"strings"
	"testing"

	"github.com/elchead/poker"
	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	store := &poker.StubPlayerStore{}
	cli := &poker.CLI{store, in}
	cli.PlayPoker()
	assert.Equal(t, 1, len(store.winCalls))
	assert.Equal(t, "Chris", store.winCalls[0])

}
