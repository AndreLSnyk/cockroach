// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package engine

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/storage/engine/enginepb"
)

// NewInMem allocates and returns a new, opened in-memory engine. The caller
// must call the engine's Close method when the engine is no longer needed.
//
// FIXME(tschottdorf): make the signature similar to NewRocksDB (require a cfg).
func NewInMem(engine enginepb.EngineType, attrs roachpb.Attributes, cacheSize int64) Engine {
	switch engine {
	case enginepb.EngineTypePebble:
		return newPebbleInMem(attrs, cacheSize)
	case enginepb.EngineTypeRocksDB:
		return newRocksDBInMem(attrs, cacheSize)
	}
	panic(fmt.Sprintf("unknown engine type: %d", engine))
}

// NewDefaultInMem allocates and returns a new, opened in-memory engine with
// the default configuration. The caller must call the engine's Close method
// when the engine is no longer needed.
func NewDefaultInMem() Engine {
	return NewInMem(TestStorageEngine, roachpb.Attributes{}, 1<<20 /* 1 MB */)
}
