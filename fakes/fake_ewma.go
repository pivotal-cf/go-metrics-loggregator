// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/rcrowley/go-metrics"
)

type FakeEWMA struct {
	RateStub        func() float64
	rateMutex       sync.RWMutex
	rateArgsForCall []struct{}
	rateReturns     struct {
		result1 float64
	}
	rateReturnsOnCall map[int]struct {
		result1 float64
	}
	SnapshotStub        func() metrics.EWMA
	snapshotMutex       sync.RWMutex
	snapshotArgsForCall []struct{}
	snapshotReturns     struct {
		result1 metrics.EWMA
	}
	snapshotReturnsOnCall map[int]struct {
		result1 metrics.EWMA
	}
	TickStub          func()
	tickMutex         sync.RWMutex
	tickArgsForCall   []struct{}
	UpdateStub        func(int64)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 int64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEWMA) Rate() float64 {
	fake.rateMutex.Lock()
	ret, specificReturn := fake.rateReturnsOnCall[len(fake.rateArgsForCall)]
	fake.rateArgsForCall = append(fake.rateArgsForCall, struct{}{})
	fake.recordInvocation("Rate", []interface{}{})
	fake.rateMutex.Unlock()
	if fake.RateStub != nil {
		return fake.RateStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rateReturns.result1
}

func (fake *FakeEWMA) RateCallCount() int {
	fake.rateMutex.RLock()
	defer fake.rateMutex.RUnlock()
	return len(fake.rateArgsForCall)
}

func (fake *FakeEWMA) RateReturns(result1 float64) {
	fake.RateStub = nil
	fake.rateReturns = struct {
		result1 float64
	}{result1}
}

func (fake *FakeEWMA) RateReturnsOnCall(i int, result1 float64) {
	fake.RateStub = nil
	if fake.rateReturnsOnCall == nil {
		fake.rateReturnsOnCall = make(map[int]struct {
			result1 float64
		})
	}
	fake.rateReturnsOnCall[i] = struct {
		result1 float64
	}{result1}
}

func (fake *FakeEWMA) Snapshot() metrics.EWMA {
	fake.snapshotMutex.Lock()
	ret, specificReturn := fake.snapshotReturnsOnCall[len(fake.snapshotArgsForCall)]
	fake.snapshotArgsForCall = append(fake.snapshotArgsForCall, struct{}{})
	fake.recordInvocation("Snapshot", []interface{}{})
	fake.snapshotMutex.Unlock()
	if fake.SnapshotStub != nil {
		return fake.SnapshotStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.snapshotReturns.result1
}

func (fake *FakeEWMA) SnapshotCallCount() int {
	fake.snapshotMutex.RLock()
	defer fake.snapshotMutex.RUnlock()
	return len(fake.snapshotArgsForCall)
}

func (fake *FakeEWMA) SnapshotReturns(result1 metrics.EWMA) {
	fake.SnapshotStub = nil
	fake.snapshotReturns = struct {
		result1 metrics.EWMA
	}{result1}
}

func (fake *FakeEWMA) SnapshotReturnsOnCall(i int, result1 metrics.EWMA) {
	fake.SnapshotStub = nil
	if fake.snapshotReturnsOnCall == nil {
		fake.snapshotReturnsOnCall = make(map[int]struct {
			result1 metrics.EWMA
		})
	}
	fake.snapshotReturnsOnCall[i] = struct {
		result1 metrics.EWMA
	}{result1}
}

func (fake *FakeEWMA) Tick() {
	fake.tickMutex.Lock()
	fake.tickArgsForCall = append(fake.tickArgsForCall, struct{}{})
	fake.recordInvocation("Tick", []interface{}{})
	fake.tickMutex.Unlock()
	if fake.TickStub != nil {
		fake.TickStub()
	}
}

func (fake *FakeEWMA) TickCallCount() int {
	fake.tickMutex.RLock()
	defer fake.tickMutex.RUnlock()
	return len(fake.tickArgsForCall)
}

func (fake *FakeEWMA) Update(arg1 int64) {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 int64
	}{arg1})
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		fake.UpdateStub(arg1)
	}
}

func (fake *FakeEWMA) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeEWMA) UpdateArgsForCall(i int) int64 {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1
}

func (fake *FakeEWMA) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.rateMutex.RLock()
	defer fake.rateMutex.RUnlock()
	fake.snapshotMutex.RLock()
	defer fake.snapshotMutex.RUnlock()
	fake.tickMutex.RLock()
	defer fake.tickMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEWMA) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.EWMA = new(FakeEWMA)