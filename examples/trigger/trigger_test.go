package sample

import (
	"testing"

	"github.com/AiRISTAFlowInc/flow-studio-core/support"
	"github.com/AiRISTAFlowInc/flow-studio-core/trigger"
	"github.com/stretchr/testify/assert"
)

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}
