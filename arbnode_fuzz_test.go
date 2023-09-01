package arbnode

import (
	"context"
	"math/big"
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_GetBatchCount(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}

		var ctx context.Context
		var blockNumber big.Int

		err = tp.Fill(&ctx)
		if err != nil {
			return
		}

		err = tp.Fill(&blockNumber)
		if err != nil {
			return
		}

		client := // Initialize your arbutil.L1Interface here
		addr := // Initialize your common.Address here
		seqInbox, _ := NewSequencerInbox(client, addr, 0)

		// Call the function being fuzzed
		seqInbox.GetBatchCount(ctx, &blockNumber)
	})
}
