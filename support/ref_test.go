package support

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
}

func TestGetRef(t *testing.T) {

	ts := &TestStruct{}
	ref := GetRef(ts)
	assert.Equal(t, "github.com/AiRISTAFlowInc/flow-studio-core/support", ref)
}
