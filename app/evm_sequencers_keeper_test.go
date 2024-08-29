package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEVMWrappedSeqKeeper_generateValidator(t *testing.T) {
	k := EVMWrappedSeqKeeper{}
	v := k.generateValidator()
	require.NotPanics(t, func() {
		_ = v.GetOperator()
	})
}
