package poker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	store := &StubPlayerStore{}
	cli := &CLI{store, in}
	cli.PlayPoker()
	assert.Equal(t, 1, len(store.winCalls))
	assert.Equal(t, "Chris", store.winCalls[0])

}
