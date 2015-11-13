// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/agent/script"
)

type FakeScript struct {
	TagStub        func() string
	tagMutex       sync.RWMutex
	tagArgsForCall []struct{}
	tagReturns     struct {
		result1 string
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 string
	}
	ExistsStub        func() bool
	existsMutex       sync.RWMutex
	existsArgsForCall []struct{}
	existsReturns     struct {
		result1 bool
	}
	RunStub        func() error
	runMutex       sync.RWMutex
	runArgsForCall []struct{}
	runReturns     struct {
		result1 error
	}
}

func (fake *FakeScript) Tag() string {
	fake.tagMutex.Lock()
	fake.tagArgsForCall = append(fake.tagArgsForCall, struct{}{})
	fake.tagMutex.Unlock()
	if fake.TagStub != nil {
		return fake.TagStub()
	} else {
		return fake.tagReturns.result1
	}
}

func (fake *FakeScript) TagCallCount() int {
	fake.tagMutex.RLock()
	defer fake.tagMutex.RUnlock()
	return len(fake.tagArgsForCall)
}

func (fake *FakeScript) TagReturns(result1 string) {
	fake.TagStub = nil
	fake.tagReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeScript) Path() string {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	} else {
		return fake.pathReturns.result1
	}
}

func (fake *FakeScript) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeScript) PathReturns(result1 string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeScript) Exists() bool {
	fake.existsMutex.Lock()
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct{}{})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub()
	} else {
		return fake.existsReturns.result1
	}
}

func (fake *FakeScript) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *FakeScript) ExistsReturns(result1 bool) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeScript) Run() error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub()
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeScript) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeScript) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

var _ script.Script = new(FakeScript)
