// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package sync

import (
	"errors"

	"github.com/medibloc/go-medibloc/core"
)

// SyncService messages
const (
	SyncMetaRequest       = "meta_req"
	SyncMeta              = "meta"
	SyncBlockChunkRequest = "chunk_req"
	SyncBlockChunk        = "chunk"
)

// ErrAlreadyDownlaodActivated occurred sync is already activated
var (
	ErrAlreadyDownlaodActivated = errors.New("download manager is already activated")
)

//BlockManager is interface of core.blockmanager.
type BlockManager interface {
	Start()
	PushBlockData(block *core.BlockData) error
	BroadCast(block *core.BlockData) error
}

//ChainManager is interface of core.chainManager
type ChainManager interface {
	LIB() *core.Block
	SetLIB(block *core.Block) error
	MainTailBlock() *core.Block
	BlockByHeight(uint64) *core.Block
}
