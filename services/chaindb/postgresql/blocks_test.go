// Copyright © 2021 Weald Technology Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package postgresql_test

import (
	"context"
	"os"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/chaind/services/chaindb"
	"github.com/wealdtech/chaind/services/chaindb/postgresql"
)

func TestIndeterminateBlocks(t *testing.T) {
	ctx := context.Background()
	s, err := postgresql.New(ctx,
		postgresql.WithLogLevel(zerolog.Disabled),
		postgresql.WithConnectionURL(os.Getenv("CHAINDB_DATABASE_URL")),
	)
	require.NoError(t, err)

	// Set a block without a canonical status.
	block := &chaindb.Block{
		Slot:          1,
		ProposerIndex: 2,
		Root: phase0.Root{
			0x00, 0x01, 0x02, 0x03, 0x04, 0x04, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x04, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		},
		Graffiti: []byte{
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		},
		RANDAOReveal: phase0.BLSSignature{
			0x10, 0x11, 0x12, 0x13, 0x14, 0x14, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x14, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x14, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x14, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x14, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x14, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		},
		BodyRoot: phase0.Root{
			0x20, 0x21, 0x22, 0x23, 0x24, 0x24, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
			0x20, 0x21, 0x22, 0x23, 0x24, 0x24, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		},
		ParentRoot: phase0.Root{
			0x30, 0x31, 0x32, 0x33, 0x34, 0x34, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
			0x30, 0x31, 0x32, 0x33, 0x34, 0x34, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
		},
		StateRoot: phase0.Root{
			0x40, 0x41, 0x42, 0x43, 0x44, 0x44, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
			0x40, 0x41, 0x42, 0x43, 0x44, 0x44, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
		},
		ETH1BlockHash: []byte{
			0x50, 0x51, 0x52, 0x53, 0x54, 0x54, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
			0x50, 0x51, 0x52, 0x53, 0x54, 0x54, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
		},
		ETH1DepositRoot: phase0.Root{
			0x60, 0x61, 0x62, 0x63, 0x64, 0x64, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
			0x60, 0x61, 0x62, 0x63, 0x64, 0x64, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
		},
	}

	ctx, cancel, err := s.BeginTx(ctx)
	require.NoError(t, err)
	defer cancel()

	// Set the block.
	err = s.SetBlock(ctx, block)
	require.NoError(t, err)
	// Fetch the block; ensure canonical is nil.
	dbBlock, err := s.BlockByRoot(ctx, block.Root)
	require.NoError(t, err)
	require.Nil(t, dbBlock.Canonical)

	// Update the block to be non-canonical.
	canonical := false
	block.Canonical = &canonical
	err = s.SetBlock(ctx, block)
	require.NoError(t, err)
	// Fetch the block; ensure canonical is false.
	dbBlock, err = s.BlockByRoot(ctx, block.Root)
	require.NoError(t, err)
	require.NotNil(t, dbBlock.Canonical)
	require.False(t, *dbBlock.Canonical)

	// Update the block to be canonical.
	canonical = true
	block.Canonical = &canonical
	err = s.SetBlock(ctx, block)
	require.NoError(t, err)
	// Fetch the block; ensure canonical is true.
	dbBlock, err = s.BlockByRoot(ctx, block.Root)
	require.NoError(t, err)
	require.NotNil(t, dbBlock.Canonical)
	require.True(t, *dbBlock.Canonical)
}
