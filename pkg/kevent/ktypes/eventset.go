/*
 * Copyright 2021-2022 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ktypes

import (
	"github.com/bits-and-blooms/bitset"
	"golang.org/x/sys/windows"
)

// EventsetMasks allows efficient testing
// of a group of bitsets containing event
// hook identifiers. For each provider
// is represented by a GUID, a dedicated
// bitset is defined.
type EventsetMasks struct {
	masks [ProvidersCount]bitset.BitSet
}

// Set puts a new event type into the bitset.
func (e *EventsetMasks) Set(ktype Ktype) {
	i := e.bitsetIndex(ktype.GUID())
	if i < 0 {
		panic("invalid bitset index")
	}
	e.masks[i].Set(uint(ktype.HookID()))
}

// Test checks if the given provider GUID and
// hook identifier are present in the bitset.
func (e *EventsetMasks) Test(guid windows.GUID, hookID uint16) bool {
	i := e.bitsetIndex(guid)
	if i < 0 {
		return false
	}
	return e.masks[i].Test(uint(hookID))
}

// Clear clears the bitset for a given provider GUID.
func (e *EventsetMasks) Clear(guid windows.GUID) {
	i := e.bitsetIndex(guid)
	if i < 0 {
		panic("invalid bitset index")
	}
	e.masks[i].ClearAll()
}

func (e *EventsetMasks) bitsetIndex(guid windows.GUID) int {
	switch guid {
	case ProcessEventGUID:
		return 0
	case ThreadEventGUID:
		return 1
	case ImageEventGUID:
		return 2
	case FileEventGUID:
		return 3
	case RegistryEventGUID:
		return 4
	case NetworkEventGUID:
		return 5
	case HandleEventGUID:
		return 6
	case MemEventGUID:
		return 7
	case AuditAPIEventGUID:
		return 8
	case DNSEventGUID:
		return 9
	}
	return -1
}
