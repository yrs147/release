// Code generated by counterfeiter. DO NOT EDIT.
package objectfakes

import (
	"sync"

	"k8s.io/release/pkg/object"
)

type FakeStore struct {
	CopyBucketToBucketStub        func(string, string) error
	copyBucketToBucketMutex       sync.RWMutex
	copyBucketToBucketArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyBucketToBucketReturns struct {
		result1 error
	}
	copyBucketToBucketReturnsOnCall map[int]struct {
		result1 error
	}
	CopyToLocalStub        func(string, string) error
	copyToLocalMutex       sync.RWMutex
	copyToLocalArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyToLocalReturns struct {
		result1 error
	}
	copyToLocalReturnsOnCall map[int]struct {
		result1 error
	}
	CopyToRemoteStub        func(string, string) error
	copyToRemoteMutex       sync.RWMutex
	copyToRemoteArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyToRemoteReturns struct {
		result1 error
	}
	copyToRemoteReturnsOnCall map[int]struct {
		result1 error
	}
	GetMarkerPathStub        func(string, string) (string, error)
	getMarkerPathMutex       sync.RWMutex
	getMarkerPathArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getMarkerPathReturns struct {
		result1 string
		result2 error
	}
	getMarkerPathReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetReleasePathStub        func(string, string, string, bool) (string, error)
	getReleasePathMutex       sync.RWMutex
	getReleasePathArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}
	getReleasePathReturns struct {
		result1 string
		result2 error
	}
	getReleasePathReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	IsPathNormalizedStub        func(string) bool
	isPathNormalizedMutex       sync.RWMutex
	isPathNormalizedArgsForCall []struct {
		arg1 string
	}
	isPathNormalizedReturns struct {
		result1 bool
	}
	isPathNormalizedReturnsOnCall map[int]struct {
		result1 bool
	}
	NormalizePathStub        func(...string) (string, error)
	normalizePathMutex       sync.RWMutex
	normalizePathArgsForCall []struct {
		arg1 []string
	}
	normalizePathReturns struct {
		result1 string
		result2 error
	}
	normalizePathReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	PathExistsStub        func(string) (bool, error)
	pathExistsMutex       sync.RWMutex
	pathExistsArgsForCall []struct {
		arg1 string
	}
	pathExistsReturns struct {
		result1 bool
		result2 error
	}
	pathExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RsyncRecursiveStub        func(string, string) error
	rsyncRecursiveMutex       sync.RWMutex
	rsyncRecursiveArgsForCall []struct {
		arg1 string
		arg2 string
	}
	rsyncRecursiveReturns struct {
		result1 error
	}
	rsyncRecursiveReturnsOnCall map[int]struct {
		result1 error
	}
	SetOptionsStub        func(...object.OptFn)
	setOptionsMutex       sync.RWMutex
	setOptionsArgsForCall []struct {
		arg1 []object.OptFn
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) CopyBucketToBucket(arg1 string, arg2 string) error {
	fake.copyBucketToBucketMutex.Lock()
	ret, specificReturn := fake.copyBucketToBucketReturnsOnCall[len(fake.copyBucketToBucketArgsForCall)]
	fake.copyBucketToBucketArgsForCall = append(fake.copyBucketToBucketArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CopyBucketToBucketStub
	fakeReturns := fake.copyBucketToBucketReturns
	fake.recordInvocation("CopyBucketToBucket", []interface{}{arg1, arg2})
	fake.copyBucketToBucketMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) CopyBucketToBucketCallCount() int {
	fake.copyBucketToBucketMutex.RLock()
	defer fake.copyBucketToBucketMutex.RUnlock()
	return len(fake.copyBucketToBucketArgsForCall)
}

func (fake *FakeStore) CopyBucketToBucketCalls(stub func(string, string) error) {
	fake.copyBucketToBucketMutex.Lock()
	defer fake.copyBucketToBucketMutex.Unlock()
	fake.CopyBucketToBucketStub = stub
}

func (fake *FakeStore) CopyBucketToBucketArgsForCall(i int) (string, string) {
	fake.copyBucketToBucketMutex.RLock()
	defer fake.copyBucketToBucketMutex.RUnlock()
	argsForCall := fake.copyBucketToBucketArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) CopyBucketToBucketReturns(result1 error) {
	fake.copyBucketToBucketMutex.Lock()
	defer fake.copyBucketToBucketMutex.Unlock()
	fake.CopyBucketToBucketStub = nil
	fake.copyBucketToBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) CopyBucketToBucketReturnsOnCall(i int, result1 error) {
	fake.copyBucketToBucketMutex.Lock()
	defer fake.copyBucketToBucketMutex.Unlock()
	fake.CopyBucketToBucketStub = nil
	if fake.copyBucketToBucketReturnsOnCall == nil {
		fake.copyBucketToBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBucketToBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) CopyToLocal(arg1 string, arg2 string) error {
	fake.copyToLocalMutex.Lock()
	ret, specificReturn := fake.copyToLocalReturnsOnCall[len(fake.copyToLocalArgsForCall)]
	fake.copyToLocalArgsForCall = append(fake.copyToLocalArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CopyToLocalStub
	fakeReturns := fake.copyToLocalReturns
	fake.recordInvocation("CopyToLocal", []interface{}{arg1, arg2})
	fake.copyToLocalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) CopyToLocalCallCount() int {
	fake.copyToLocalMutex.RLock()
	defer fake.copyToLocalMutex.RUnlock()
	return len(fake.copyToLocalArgsForCall)
}

func (fake *FakeStore) CopyToLocalCalls(stub func(string, string) error) {
	fake.copyToLocalMutex.Lock()
	defer fake.copyToLocalMutex.Unlock()
	fake.CopyToLocalStub = stub
}

func (fake *FakeStore) CopyToLocalArgsForCall(i int) (string, string) {
	fake.copyToLocalMutex.RLock()
	defer fake.copyToLocalMutex.RUnlock()
	argsForCall := fake.copyToLocalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) CopyToLocalReturns(result1 error) {
	fake.copyToLocalMutex.Lock()
	defer fake.copyToLocalMutex.Unlock()
	fake.CopyToLocalStub = nil
	fake.copyToLocalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) CopyToLocalReturnsOnCall(i int, result1 error) {
	fake.copyToLocalMutex.Lock()
	defer fake.copyToLocalMutex.Unlock()
	fake.CopyToLocalStub = nil
	if fake.copyToLocalReturnsOnCall == nil {
		fake.copyToLocalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyToLocalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) CopyToRemote(arg1 string, arg2 string) error {
	fake.copyToRemoteMutex.Lock()
	ret, specificReturn := fake.copyToRemoteReturnsOnCall[len(fake.copyToRemoteArgsForCall)]
	fake.copyToRemoteArgsForCall = append(fake.copyToRemoteArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CopyToRemoteStub
	fakeReturns := fake.copyToRemoteReturns
	fake.recordInvocation("CopyToRemote", []interface{}{arg1, arg2})
	fake.copyToRemoteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) CopyToRemoteCallCount() int {
	fake.copyToRemoteMutex.RLock()
	defer fake.copyToRemoteMutex.RUnlock()
	return len(fake.copyToRemoteArgsForCall)
}

func (fake *FakeStore) CopyToRemoteCalls(stub func(string, string) error) {
	fake.copyToRemoteMutex.Lock()
	defer fake.copyToRemoteMutex.Unlock()
	fake.CopyToRemoteStub = stub
}

func (fake *FakeStore) CopyToRemoteArgsForCall(i int) (string, string) {
	fake.copyToRemoteMutex.RLock()
	defer fake.copyToRemoteMutex.RUnlock()
	argsForCall := fake.copyToRemoteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) CopyToRemoteReturns(result1 error) {
	fake.copyToRemoteMutex.Lock()
	defer fake.copyToRemoteMutex.Unlock()
	fake.CopyToRemoteStub = nil
	fake.copyToRemoteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) CopyToRemoteReturnsOnCall(i int, result1 error) {
	fake.copyToRemoteMutex.Lock()
	defer fake.copyToRemoteMutex.Unlock()
	fake.CopyToRemoteStub = nil
	if fake.copyToRemoteReturnsOnCall == nil {
		fake.copyToRemoteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyToRemoteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetMarkerPath(arg1 string, arg2 string) (string, error) {
	fake.getMarkerPathMutex.Lock()
	ret, specificReturn := fake.getMarkerPathReturnsOnCall[len(fake.getMarkerPathArgsForCall)]
	fake.getMarkerPathArgsForCall = append(fake.getMarkerPathArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetMarkerPathStub
	fakeReturns := fake.getMarkerPathReturns
	fake.recordInvocation("GetMarkerPath", []interface{}{arg1, arg2})
	fake.getMarkerPathMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetMarkerPathCallCount() int {
	fake.getMarkerPathMutex.RLock()
	defer fake.getMarkerPathMutex.RUnlock()
	return len(fake.getMarkerPathArgsForCall)
}

func (fake *FakeStore) GetMarkerPathCalls(stub func(string, string) (string, error)) {
	fake.getMarkerPathMutex.Lock()
	defer fake.getMarkerPathMutex.Unlock()
	fake.GetMarkerPathStub = stub
}

func (fake *FakeStore) GetMarkerPathArgsForCall(i int) (string, string) {
	fake.getMarkerPathMutex.RLock()
	defer fake.getMarkerPathMutex.RUnlock()
	argsForCall := fake.getMarkerPathArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) GetMarkerPathReturns(result1 string, result2 error) {
	fake.getMarkerPathMutex.Lock()
	defer fake.getMarkerPathMutex.Unlock()
	fake.GetMarkerPathStub = nil
	fake.getMarkerPathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetMarkerPathReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMarkerPathMutex.Lock()
	defer fake.getMarkerPathMutex.Unlock()
	fake.GetMarkerPathStub = nil
	if fake.getMarkerPathReturnsOnCall == nil {
		fake.getMarkerPathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getMarkerPathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetReleasePath(arg1 string, arg2 string, arg3 string, arg4 bool) (string, error) {
	fake.getReleasePathMutex.Lock()
	ret, specificReturn := fake.getReleasePathReturnsOnCall[len(fake.getReleasePathArgsForCall)]
	fake.getReleasePathArgsForCall = append(fake.getReleasePathArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetReleasePathStub
	fakeReturns := fake.getReleasePathReturns
	fake.recordInvocation("GetReleasePath", []interface{}{arg1, arg2, arg3, arg4})
	fake.getReleasePathMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetReleasePathCallCount() int {
	fake.getReleasePathMutex.RLock()
	defer fake.getReleasePathMutex.RUnlock()
	return len(fake.getReleasePathArgsForCall)
}

func (fake *FakeStore) GetReleasePathCalls(stub func(string, string, string, bool) (string, error)) {
	fake.getReleasePathMutex.Lock()
	defer fake.getReleasePathMutex.Unlock()
	fake.GetReleasePathStub = stub
}

func (fake *FakeStore) GetReleasePathArgsForCall(i int) (string, string, string, bool) {
	fake.getReleasePathMutex.RLock()
	defer fake.getReleasePathMutex.RUnlock()
	argsForCall := fake.getReleasePathArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeStore) GetReleasePathReturns(result1 string, result2 error) {
	fake.getReleasePathMutex.Lock()
	defer fake.getReleasePathMutex.Unlock()
	fake.GetReleasePathStub = nil
	fake.getReleasePathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetReleasePathReturnsOnCall(i int, result1 string, result2 error) {
	fake.getReleasePathMutex.Lock()
	defer fake.getReleasePathMutex.Unlock()
	fake.GetReleasePathStub = nil
	if fake.getReleasePathReturnsOnCall == nil {
		fake.getReleasePathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getReleasePathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) IsPathNormalized(arg1 string) bool {
	fake.isPathNormalizedMutex.Lock()
	ret, specificReturn := fake.isPathNormalizedReturnsOnCall[len(fake.isPathNormalizedArgsForCall)]
	fake.isPathNormalizedArgsForCall = append(fake.isPathNormalizedArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IsPathNormalizedStub
	fakeReturns := fake.isPathNormalizedReturns
	fake.recordInvocation("IsPathNormalized", []interface{}{arg1})
	fake.isPathNormalizedMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) IsPathNormalizedCallCount() int {
	fake.isPathNormalizedMutex.RLock()
	defer fake.isPathNormalizedMutex.RUnlock()
	return len(fake.isPathNormalizedArgsForCall)
}

func (fake *FakeStore) IsPathNormalizedCalls(stub func(string) bool) {
	fake.isPathNormalizedMutex.Lock()
	defer fake.isPathNormalizedMutex.Unlock()
	fake.IsPathNormalizedStub = stub
}

func (fake *FakeStore) IsPathNormalizedArgsForCall(i int) string {
	fake.isPathNormalizedMutex.RLock()
	defer fake.isPathNormalizedMutex.RUnlock()
	argsForCall := fake.isPathNormalizedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) IsPathNormalizedReturns(result1 bool) {
	fake.isPathNormalizedMutex.Lock()
	defer fake.isPathNormalizedMutex.Unlock()
	fake.IsPathNormalizedStub = nil
	fake.isPathNormalizedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStore) IsPathNormalizedReturnsOnCall(i int, result1 bool) {
	fake.isPathNormalizedMutex.Lock()
	defer fake.isPathNormalizedMutex.Unlock()
	fake.IsPathNormalizedStub = nil
	if fake.isPathNormalizedReturnsOnCall == nil {
		fake.isPathNormalizedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isPathNormalizedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStore) NormalizePath(arg1 ...string) (string, error) {
	fake.normalizePathMutex.Lock()
	ret, specificReturn := fake.normalizePathReturnsOnCall[len(fake.normalizePathArgsForCall)]
	fake.normalizePathArgsForCall = append(fake.normalizePathArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.NormalizePathStub
	fakeReturns := fake.normalizePathReturns
	fake.recordInvocation("NormalizePath", []interface{}{arg1})
	fake.normalizePathMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) NormalizePathCallCount() int {
	fake.normalizePathMutex.RLock()
	defer fake.normalizePathMutex.RUnlock()
	return len(fake.normalizePathArgsForCall)
}

func (fake *FakeStore) NormalizePathCalls(stub func(...string) (string, error)) {
	fake.normalizePathMutex.Lock()
	defer fake.normalizePathMutex.Unlock()
	fake.NormalizePathStub = stub
}

func (fake *FakeStore) NormalizePathArgsForCall(i int) []string {
	fake.normalizePathMutex.RLock()
	defer fake.normalizePathMutex.RUnlock()
	argsForCall := fake.normalizePathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) NormalizePathReturns(result1 string, result2 error) {
	fake.normalizePathMutex.Lock()
	defer fake.normalizePathMutex.Unlock()
	fake.NormalizePathStub = nil
	fake.normalizePathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) NormalizePathReturnsOnCall(i int, result1 string, result2 error) {
	fake.normalizePathMutex.Lock()
	defer fake.normalizePathMutex.Unlock()
	fake.NormalizePathStub = nil
	if fake.normalizePathReturnsOnCall == nil {
		fake.normalizePathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.normalizePathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) PathExists(arg1 string) (bool, error) {
	fake.pathExistsMutex.Lock()
	ret, specificReturn := fake.pathExistsReturnsOnCall[len(fake.pathExistsArgsForCall)]
	fake.pathExistsArgsForCall = append(fake.pathExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.PathExistsStub
	fakeReturns := fake.pathExistsReturns
	fake.recordInvocation("PathExists", []interface{}{arg1})
	fake.pathExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) PathExistsCallCount() int {
	fake.pathExistsMutex.RLock()
	defer fake.pathExistsMutex.RUnlock()
	return len(fake.pathExistsArgsForCall)
}

func (fake *FakeStore) PathExistsCalls(stub func(string) (bool, error)) {
	fake.pathExistsMutex.Lock()
	defer fake.pathExistsMutex.Unlock()
	fake.PathExistsStub = stub
}

func (fake *FakeStore) PathExistsArgsForCall(i int) string {
	fake.pathExistsMutex.RLock()
	defer fake.pathExistsMutex.RUnlock()
	argsForCall := fake.pathExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) PathExistsReturns(result1 bool, result2 error) {
	fake.pathExistsMutex.Lock()
	defer fake.pathExistsMutex.Unlock()
	fake.PathExistsStub = nil
	fake.pathExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) PathExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.pathExistsMutex.Lock()
	defer fake.pathExistsMutex.Unlock()
	fake.PathExistsStub = nil
	if fake.pathExistsReturnsOnCall == nil {
		fake.pathExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.pathExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) RsyncRecursive(arg1 string, arg2 string) error {
	fake.rsyncRecursiveMutex.Lock()
	ret, specificReturn := fake.rsyncRecursiveReturnsOnCall[len(fake.rsyncRecursiveArgsForCall)]
	fake.rsyncRecursiveArgsForCall = append(fake.rsyncRecursiveArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.RsyncRecursiveStub
	fakeReturns := fake.rsyncRecursiveReturns
	fake.recordInvocation("RsyncRecursive", []interface{}{arg1, arg2})
	fake.rsyncRecursiveMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) RsyncRecursiveCallCount() int {
	fake.rsyncRecursiveMutex.RLock()
	defer fake.rsyncRecursiveMutex.RUnlock()
	return len(fake.rsyncRecursiveArgsForCall)
}

func (fake *FakeStore) RsyncRecursiveCalls(stub func(string, string) error) {
	fake.rsyncRecursiveMutex.Lock()
	defer fake.rsyncRecursiveMutex.Unlock()
	fake.RsyncRecursiveStub = stub
}

func (fake *FakeStore) RsyncRecursiveArgsForCall(i int) (string, string) {
	fake.rsyncRecursiveMutex.RLock()
	defer fake.rsyncRecursiveMutex.RUnlock()
	argsForCall := fake.rsyncRecursiveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) RsyncRecursiveReturns(result1 error) {
	fake.rsyncRecursiveMutex.Lock()
	defer fake.rsyncRecursiveMutex.Unlock()
	fake.RsyncRecursiveStub = nil
	fake.rsyncRecursiveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) RsyncRecursiveReturnsOnCall(i int, result1 error) {
	fake.rsyncRecursiveMutex.Lock()
	defer fake.rsyncRecursiveMutex.Unlock()
	fake.RsyncRecursiveStub = nil
	if fake.rsyncRecursiveReturnsOnCall == nil {
		fake.rsyncRecursiveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rsyncRecursiveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) SetOptions(arg1 ...object.OptFn) {
	fake.setOptionsMutex.Lock()
	fake.setOptionsArgsForCall = append(fake.setOptionsArgsForCall, struct {
		arg1 []object.OptFn
	}{arg1})
	stub := fake.SetOptionsStub
	fake.recordInvocation("SetOptions", []interface{}{arg1})
	fake.setOptionsMutex.Unlock()
	if stub != nil {
		fake.SetOptionsStub(arg1...)
	}
}

func (fake *FakeStore) SetOptionsCallCount() int {
	fake.setOptionsMutex.RLock()
	defer fake.setOptionsMutex.RUnlock()
	return len(fake.setOptionsArgsForCall)
}

func (fake *FakeStore) SetOptionsCalls(stub func(...object.OptFn)) {
	fake.setOptionsMutex.Lock()
	defer fake.setOptionsMutex.Unlock()
	fake.SetOptionsStub = stub
}

func (fake *FakeStore) SetOptionsArgsForCall(i int) []object.OptFn {
	fake.setOptionsMutex.RLock()
	defer fake.setOptionsMutex.RUnlock()
	argsForCall := fake.setOptionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.copyBucketToBucketMutex.RLock()
	defer fake.copyBucketToBucketMutex.RUnlock()
	fake.copyToLocalMutex.RLock()
	defer fake.copyToLocalMutex.RUnlock()
	fake.copyToRemoteMutex.RLock()
	defer fake.copyToRemoteMutex.RUnlock()
	fake.getMarkerPathMutex.RLock()
	defer fake.getMarkerPathMutex.RUnlock()
	fake.getReleasePathMutex.RLock()
	defer fake.getReleasePathMutex.RUnlock()
	fake.isPathNormalizedMutex.RLock()
	defer fake.isPathNormalizedMutex.RUnlock()
	fake.normalizePathMutex.RLock()
	defer fake.normalizePathMutex.RUnlock()
	fake.pathExistsMutex.RLock()
	defer fake.pathExistsMutex.RUnlock()
	fake.rsyncRecursiveMutex.RLock()
	defer fake.rsyncRecursiveMutex.RUnlock()
	fake.setOptionsMutex.RLock()
	defer fake.setOptionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
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

var _ object.Store = new(FakeStore)