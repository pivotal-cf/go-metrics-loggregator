// This file was generated by counterfeiter
package gometricspcffakes

import (
	"sync"

	"github.com/rcrowley/go-metrics"
)

type FakeMeter struct {
	CountStub        func() int64
	countMutex       sync.RWMutex
	countArgsForCall []struct{}
	countReturns     struct {
		result1 int64
	}
	countReturnsOnCall map[int]struct {
		result1 int64
	}
	MarkStub        func(int64)
	markMutex       sync.RWMutex
	markArgsForCall []struct {
		arg1 int64
	}
	Rate1Stub        func() float64
	rate1Mutex       sync.RWMutex
	rate1ArgsForCall []struct{}
	rate1Returns     struct {
		result1 float64
	}
	rate1ReturnsOnCall map[int]struct {
		result1 float64
	}
	Rate5Stub        func() float64
	rate5Mutex       sync.RWMutex
	rate5ArgsForCall []struct{}
	rate5Returns     struct {
		result1 float64
	}
	rate5ReturnsOnCall map[int]struct {
		result1 float64
	}
	Rate15Stub        func() float64
	rate15Mutex       sync.RWMutex
	rate15ArgsForCall []struct{}
	rate15Returns     struct {
		result1 float64
	}
	rate15ReturnsOnCall map[int]struct {
		result1 float64
	}
	RateMeanStub        func() float64
	rateMeanMutex       sync.RWMutex
	rateMeanArgsForCall []struct{}
	rateMeanReturns     struct {
		result1 float64
	}
	rateMeanReturnsOnCall map[int]struct {
		result1 float64
	}
	SnapshotStub        func() metrics.Meter
	snapshotMutex       sync.RWMutex
	snapshotArgsForCall []struct{}
	snapshotReturns     struct {
		result1 metrics.Meter
	}
	snapshotReturnsOnCall map[int]struct {
		result1 metrics.Meter
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMeter) Count() int64 {
	fake.countMutex.Lock()
	ret, specificReturn := fake.countReturnsOnCall[len(fake.countArgsForCall)]
	fake.countArgsForCall = append(fake.countArgsForCall, struct{}{})
	fake.recordInvocation("Count", []interface{}{})
	fake.countMutex.Unlock()
	if fake.CountStub != nil {
		return fake.CountStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.countReturns.result1
}

func (fake *FakeMeter) CountCallCount() int {
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	return len(fake.countArgsForCall)
}

func (fake *FakeMeter) CountReturns(result1 int64) {
	fake.CountStub = nil
	fake.countReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMeter) CountReturnsOnCall(i int, result1 int64) {
	fake.CountStub = nil
	if fake.countReturnsOnCall == nil {
		fake.countReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.countReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMeter) Mark(arg1 int64) {
	fake.markMutex.Lock()
	fake.markArgsForCall = append(fake.markArgsForCall, struct {
		arg1 int64
	}{arg1})
	fake.recordInvocation("Mark", []interface{}{arg1})
	fake.markMutex.Unlock()
	if fake.MarkStub != nil {
		fake.MarkStub(arg1)
	}
}

func (fake *FakeMeter) MarkCallCount() int {
	fake.markMutex.RLock()
	defer fake.markMutex.RUnlock()
	return len(fake.markArgsForCall)
}

func (fake *FakeMeter) MarkArgsForCall(i int) int64 {
	fake.markMutex.RLock()
	defer fake.markMutex.RUnlock()
	return fake.markArgsForCall[i].arg1
}

func (fake *FakeMeter) Rate1() float64 {
	fake.rate1Mutex.Lock()
	ret, specificReturn := fake.rate1ReturnsOnCall[len(fake.rate1ArgsForCall)]
	fake.rate1ArgsForCall = append(fake.rate1ArgsForCall, struct{}{})
	fake.recordInvocation("Rate1", []interface{}{})
	fake.rate1Mutex.Unlock()
	if fake.Rate1Stub != nil {
		return fake.Rate1Stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rate1Returns.result1
}

func (fake *FakeMeter) Rate1CallCount() int {
	fake.rate1Mutex.RLock()
	defer fake.rate1Mutex.RUnlock()
	return len(fake.rate1ArgsForCall)
}

func (fake *FakeMeter) Rate1Returns(result1 float64) {
	fake.Rate1Stub = nil
	fake.rate1Returns = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) Rate1ReturnsOnCall(i int, result1 float64) {
	fake.Rate1Stub = nil
	if fake.rate1ReturnsOnCall == nil {
		fake.rate1ReturnsOnCall = make(map[int]struct {
			result1 float64
		})
	}
	fake.rate1ReturnsOnCall[i] = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) Rate5() float64 {
	fake.rate5Mutex.Lock()
	ret, specificReturn := fake.rate5ReturnsOnCall[len(fake.rate5ArgsForCall)]
	fake.rate5ArgsForCall = append(fake.rate5ArgsForCall, struct{}{})
	fake.recordInvocation("Rate5", []interface{}{})
	fake.rate5Mutex.Unlock()
	if fake.Rate5Stub != nil {
		return fake.Rate5Stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rate5Returns.result1
}

func (fake *FakeMeter) Rate5CallCount() int {
	fake.rate5Mutex.RLock()
	defer fake.rate5Mutex.RUnlock()
	return len(fake.rate5ArgsForCall)
}

func (fake *FakeMeter) Rate5Returns(result1 float64) {
	fake.Rate5Stub = nil
	fake.rate5Returns = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) Rate5ReturnsOnCall(i int, result1 float64) {
	fake.Rate5Stub = nil
	if fake.rate5ReturnsOnCall == nil {
		fake.rate5ReturnsOnCall = make(map[int]struct {
			result1 float64
		})
	}
	fake.rate5ReturnsOnCall[i] = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) Rate15() float64 {
	fake.rate15Mutex.Lock()
	ret, specificReturn := fake.rate15ReturnsOnCall[len(fake.rate15ArgsForCall)]
	fake.rate15ArgsForCall = append(fake.rate15ArgsForCall, struct{}{})
	fake.recordInvocation("Rate15", []interface{}{})
	fake.rate15Mutex.Unlock()
	if fake.Rate15Stub != nil {
		return fake.Rate15Stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rate15Returns.result1
}

func (fake *FakeMeter) Rate15CallCount() int {
	fake.rate15Mutex.RLock()
	defer fake.rate15Mutex.RUnlock()
	return len(fake.rate15ArgsForCall)
}

func (fake *FakeMeter) Rate15Returns(result1 float64) {
	fake.Rate15Stub = nil
	fake.rate15Returns = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) Rate15ReturnsOnCall(i int, result1 float64) {
	fake.Rate15Stub = nil
	if fake.rate15ReturnsOnCall == nil {
		fake.rate15ReturnsOnCall = make(map[int]struct {
			result1 float64
		})
	}
	fake.rate15ReturnsOnCall[i] = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) RateMean() float64 {
	fake.rateMeanMutex.Lock()
	ret, specificReturn := fake.rateMeanReturnsOnCall[len(fake.rateMeanArgsForCall)]
	fake.rateMeanArgsForCall = append(fake.rateMeanArgsForCall, struct{}{})
	fake.recordInvocation("RateMean", []interface{}{})
	fake.rateMeanMutex.Unlock()
	if fake.RateMeanStub != nil {
		return fake.RateMeanStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rateMeanReturns.result1
}

func (fake *FakeMeter) RateMeanCallCount() int {
	fake.rateMeanMutex.RLock()
	defer fake.rateMeanMutex.RUnlock()
	return len(fake.rateMeanArgsForCall)
}

func (fake *FakeMeter) RateMeanReturns(result1 float64) {
	fake.RateMeanStub = nil
	fake.rateMeanReturns = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) RateMeanReturnsOnCall(i int, result1 float64) {
	fake.RateMeanStub = nil
	if fake.rateMeanReturnsOnCall == nil {
		fake.rateMeanReturnsOnCall = make(map[int]struct {
			result1 float64
		})
	}
	fake.rateMeanReturnsOnCall[i] = struct {
		result1 float64
	}{result1}
}

func (fake *FakeMeter) Snapshot() metrics.Meter {
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

func (fake *FakeMeter) SnapshotCallCount() int {
	fake.snapshotMutex.RLock()
	defer fake.snapshotMutex.RUnlock()
	return len(fake.snapshotArgsForCall)
}

func (fake *FakeMeter) SnapshotReturns(result1 metrics.Meter) {
	fake.SnapshotStub = nil
	fake.snapshotReturns = struct {
		result1 metrics.Meter
	}{result1}
}

func (fake *FakeMeter) SnapshotReturnsOnCall(i int, result1 metrics.Meter) {
	fake.SnapshotStub = nil
	if fake.snapshotReturnsOnCall == nil {
		fake.snapshotReturnsOnCall = make(map[int]struct {
			result1 metrics.Meter
		})
	}
	fake.snapshotReturnsOnCall[i] = struct {
		result1 metrics.Meter
	}{result1}
}

func (fake *FakeMeter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	fake.markMutex.RLock()
	defer fake.markMutex.RUnlock()
	fake.rate1Mutex.RLock()
	defer fake.rate1Mutex.RUnlock()
	fake.rate5Mutex.RLock()
	defer fake.rate5Mutex.RUnlock()
	fake.rate15Mutex.RLock()
	defer fake.rate15Mutex.RUnlock()
	fake.rateMeanMutex.RLock()
	defer fake.rateMeanMutex.RUnlock()
	fake.snapshotMutex.RLock()
	defer fake.snapshotMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMeter) recordInvocation(key string, args []interface{}) {
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

var _ metrics.Meter = new(FakeMeter)