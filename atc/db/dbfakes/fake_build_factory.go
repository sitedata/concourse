// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"

	db "github.com/concourse/concourse/atc/db"
)

type FakeBuildFactory struct {
	BuildStub        func(int) (db.Build, bool, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 int
	}
	buildReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	buildReturnsOnCall map[int]struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	GetAllStartedBuildsStub        func() ([]db.Build, error)
	getAllStartedBuildsMutex       sync.RWMutex
	getAllStartedBuildsArgsForCall []struct {
	}
	getAllStartedBuildsReturns struct {
		result1 []db.Build
		result2 error
	}
	getAllStartedBuildsReturnsOnCall map[int]struct {
		result1 []db.Build
		result2 error
	}
	GetDrainableBuildsStub        func() ([]db.Build, error)
	getDrainableBuildsMutex       sync.RWMutex
	getDrainableBuildsArgsForCall []struct {
	}
	getDrainableBuildsReturns struct {
		result1 []db.Build
		result2 error
	}
	getDrainableBuildsReturnsOnCall map[int]struct {
		result1 []db.Build
		result2 error
	}
	MarkNonInterceptibleBuildsStub        func() error
	markNonInterceptibleBuildsMutex       sync.RWMutex
	markNonInterceptibleBuildsArgsForCall []struct {
	}
	markNonInterceptibleBuildsReturns struct {
		result1 error
	}
	markNonInterceptibleBuildsReturnsOnCall map[int]struct {
		result1 error
	}
	PublicBuildsStub        func(db.Page) ([]db.Build, db.Pagination, error)
	publicBuildsMutex       sync.RWMutex
	publicBuildsArgsForCall []struct {
		arg1 db.Page
	}
	publicBuildsReturns struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	publicBuildsReturnsOnCall map[int]struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	VisibleBuildsStub        func([]string, db.Page) ([]db.Build, db.Pagination, error)
	visibleBuildsMutex       sync.RWMutex
	visibleBuildsArgsForCall []struct {
		arg1 []string
		arg2 db.Page
	}
	visibleBuildsReturns struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	visibleBuildsReturnsOnCall map[int]struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	VisibleBuildsWithTimeStub        func([]string, db.Page) ([]db.Build, db.Pagination, error)
	visibleBuildsWithTimeMutex       sync.RWMutex
	visibleBuildsWithTimeArgsForCall []struct {
		arg1 []string
		arg2 db.Page
	}
	visibleBuildsWithTimeReturns struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	visibleBuildsWithTimeReturnsOnCall map[int]struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildFactory) Build(arg1 int) (db.Build, bool, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Build", []interface{}{arg1})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.buildReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBuildFactory) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeBuildFactory) BuildArgsForCall(i int) int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildFactory) BuildReturns(result1 db.Build, result2 bool, result3 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) BuildReturnsOnCall(i int, result1 db.Build, result2 bool, result3 error) {
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 db.Build
			result2 bool
			result3 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) GetAllStartedBuilds() ([]db.Build, error) {
	fake.getAllStartedBuildsMutex.Lock()
	ret, specificReturn := fake.getAllStartedBuildsReturnsOnCall[len(fake.getAllStartedBuildsArgsForCall)]
	fake.getAllStartedBuildsArgsForCall = append(fake.getAllStartedBuildsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetAllStartedBuilds", []interface{}{})
	fake.getAllStartedBuildsMutex.Unlock()
	if fake.GetAllStartedBuildsStub != nil {
		return fake.GetAllStartedBuildsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAllStartedBuildsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildFactory) GetAllStartedBuildsCallCount() int {
	fake.getAllStartedBuildsMutex.RLock()
	defer fake.getAllStartedBuildsMutex.RUnlock()
	return len(fake.getAllStartedBuildsArgsForCall)
}

func (fake *FakeBuildFactory) GetAllStartedBuildsReturns(result1 []db.Build, result2 error) {
	fake.GetAllStartedBuildsStub = nil
	fake.getAllStartedBuildsReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildFactory) GetAllStartedBuildsReturnsOnCall(i int, result1 []db.Build, result2 error) {
	fake.GetAllStartedBuildsStub = nil
	if fake.getAllStartedBuildsReturnsOnCall == nil {
		fake.getAllStartedBuildsReturnsOnCall = make(map[int]struct {
			result1 []db.Build
			result2 error
		})
	}
	fake.getAllStartedBuildsReturnsOnCall[i] = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildFactory) GetDrainableBuilds() ([]db.Build, error) {
	fake.getDrainableBuildsMutex.Lock()
	ret, specificReturn := fake.getDrainableBuildsReturnsOnCall[len(fake.getDrainableBuildsArgsForCall)]
	fake.getDrainableBuildsArgsForCall = append(fake.getDrainableBuildsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetDrainableBuilds", []interface{}{})
	fake.getDrainableBuildsMutex.Unlock()
	if fake.GetDrainableBuildsStub != nil {
		return fake.GetDrainableBuildsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getDrainableBuildsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildFactory) GetDrainableBuildsCallCount() int {
	fake.getDrainableBuildsMutex.RLock()
	defer fake.getDrainableBuildsMutex.RUnlock()
	return len(fake.getDrainableBuildsArgsForCall)
}

func (fake *FakeBuildFactory) GetDrainableBuildsReturns(result1 []db.Build, result2 error) {
	fake.GetDrainableBuildsStub = nil
	fake.getDrainableBuildsReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildFactory) GetDrainableBuildsReturnsOnCall(i int, result1 []db.Build, result2 error) {
	fake.GetDrainableBuildsStub = nil
	if fake.getDrainableBuildsReturnsOnCall == nil {
		fake.getDrainableBuildsReturnsOnCall = make(map[int]struct {
			result1 []db.Build
			result2 error
		})
	}
	fake.getDrainableBuildsReturnsOnCall[i] = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildFactory) MarkNonInterceptibleBuilds() error {
	fake.markNonInterceptibleBuildsMutex.Lock()
	ret, specificReturn := fake.markNonInterceptibleBuildsReturnsOnCall[len(fake.markNonInterceptibleBuildsArgsForCall)]
	fake.markNonInterceptibleBuildsArgsForCall = append(fake.markNonInterceptibleBuildsArgsForCall, struct {
	}{})
	fake.recordInvocation("MarkNonInterceptibleBuilds", []interface{}{})
	fake.markNonInterceptibleBuildsMutex.Unlock()
	if fake.MarkNonInterceptibleBuildsStub != nil {
		return fake.MarkNonInterceptibleBuildsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.markNonInterceptibleBuildsReturns
	return fakeReturns.result1
}

func (fake *FakeBuildFactory) MarkNonInterceptibleBuildsCallCount() int {
	fake.markNonInterceptibleBuildsMutex.RLock()
	defer fake.markNonInterceptibleBuildsMutex.RUnlock()
	return len(fake.markNonInterceptibleBuildsArgsForCall)
}

func (fake *FakeBuildFactory) MarkNonInterceptibleBuildsReturns(result1 error) {
	fake.MarkNonInterceptibleBuildsStub = nil
	fake.markNonInterceptibleBuildsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildFactory) MarkNonInterceptibleBuildsReturnsOnCall(i int, result1 error) {
	fake.MarkNonInterceptibleBuildsStub = nil
	if fake.markNonInterceptibleBuildsReturnsOnCall == nil {
		fake.markNonInterceptibleBuildsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.markNonInterceptibleBuildsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildFactory) PublicBuilds(arg1 db.Page) ([]db.Build, db.Pagination, error) {
	fake.publicBuildsMutex.Lock()
	ret, specificReturn := fake.publicBuildsReturnsOnCall[len(fake.publicBuildsArgsForCall)]
	fake.publicBuildsArgsForCall = append(fake.publicBuildsArgsForCall, struct {
		arg1 db.Page
	}{arg1})
	fake.recordInvocation("PublicBuilds", []interface{}{arg1})
	fake.publicBuildsMutex.Unlock()
	if fake.PublicBuildsStub != nil {
		return fake.PublicBuildsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.publicBuildsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBuildFactory) PublicBuildsCallCount() int {
	fake.publicBuildsMutex.RLock()
	defer fake.publicBuildsMutex.RUnlock()
	return len(fake.publicBuildsArgsForCall)
}

func (fake *FakeBuildFactory) PublicBuildsArgsForCall(i int) db.Page {
	fake.publicBuildsMutex.RLock()
	defer fake.publicBuildsMutex.RUnlock()
	argsForCall := fake.publicBuildsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildFactory) PublicBuildsReturns(result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.PublicBuildsStub = nil
	fake.publicBuildsReturns = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) PublicBuildsReturnsOnCall(i int, result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.PublicBuildsStub = nil
	if fake.publicBuildsReturnsOnCall == nil {
		fake.publicBuildsReturnsOnCall = make(map[int]struct {
			result1 []db.Build
			result2 db.Pagination
			result3 error
		})
	}
	fake.publicBuildsReturnsOnCall[i] = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) VisibleBuilds(arg1 []string, arg2 db.Page) ([]db.Build, db.Pagination, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visibleBuildsMutex.Lock()
	ret, specificReturn := fake.visibleBuildsReturnsOnCall[len(fake.visibleBuildsArgsForCall)]
	fake.visibleBuildsArgsForCall = append(fake.visibleBuildsArgsForCall, struct {
		arg1 []string
		arg2 db.Page
	}{arg1Copy, arg2})
	fake.recordInvocation("VisibleBuilds", []interface{}{arg1Copy, arg2})
	fake.visibleBuildsMutex.Unlock()
	if fake.VisibleBuildsStub != nil {
		return fake.VisibleBuildsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.visibleBuildsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBuildFactory) VisibleBuildsCallCount() int {
	fake.visibleBuildsMutex.RLock()
	defer fake.visibleBuildsMutex.RUnlock()
	return len(fake.visibleBuildsArgsForCall)
}

func (fake *FakeBuildFactory) VisibleBuildsArgsForCall(i int) ([]string, db.Page) {
	fake.visibleBuildsMutex.RLock()
	defer fake.visibleBuildsMutex.RUnlock()
	argsForCall := fake.visibleBuildsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildFactory) VisibleBuildsReturns(result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.VisibleBuildsStub = nil
	fake.visibleBuildsReturns = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) VisibleBuildsReturnsOnCall(i int, result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.VisibleBuildsStub = nil
	if fake.visibleBuildsReturnsOnCall == nil {
		fake.visibleBuildsReturnsOnCall = make(map[int]struct {
			result1 []db.Build
			result2 db.Pagination
			result3 error
		})
	}
	fake.visibleBuildsReturnsOnCall[i] = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) VisibleBuildsWithTime(arg1 []string, arg2 db.Page) ([]db.Build, db.Pagination, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visibleBuildsWithTimeMutex.Lock()
	ret, specificReturn := fake.visibleBuildsWithTimeReturnsOnCall[len(fake.visibleBuildsWithTimeArgsForCall)]
	fake.visibleBuildsWithTimeArgsForCall = append(fake.visibleBuildsWithTimeArgsForCall, struct {
		arg1 []string
		arg2 db.Page
	}{arg1Copy, arg2})
	fake.recordInvocation("VisibleBuildsWithTime", []interface{}{arg1Copy, arg2})
	fake.visibleBuildsWithTimeMutex.Unlock()
	if fake.VisibleBuildsWithTimeStub != nil {
		return fake.VisibleBuildsWithTimeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.visibleBuildsWithTimeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBuildFactory) VisibleBuildsWithTimeCallCount() int {
	fake.visibleBuildsWithTimeMutex.RLock()
	defer fake.visibleBuildsWithTimeMutex.RUnlock()
	return len(fake.visibleBuildsWithTimeArgsForCall)
}

func (fake *FakeBuildFactory) VisibleBuildsWithTimeArgsForCall(i int) ([]string, db.Page) {
	fake.visibleBuildsWithTimeMutex.RLock()
	defer fake.visibleBuildsWithTimeMutex.RUnlock()
	argsForCall := fake.visibleBuildsWithTimeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildFactory) VisibleBuildsWithTimeReturns(result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.VisibleBuildsWithTimeStub = nil
	fake.visibleBuildsWithTimeReturns = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) VisibleBuildsWithTimeReturnsOnCall(i int, result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.VisibleBuildsWithTimeStub = nil
	if fake.visibleBuildsWithTimeReturnsOnCall == nil {
		fake.visibleBuildsWithTimeReturnsOnCall = make(map[int]struct {
			result1 []db.Build
			result2 db.Pagination
			result3 error
		})
	}
	fake.visibleBuildsWithTimeReturnsOnCall[i] = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.getAllStartedBuildsMutex.RLock()
	defer fake.getAllStartedBuildsMutex.RUnlock()
	fake.getDrainableBuildsMutex.RLock()
	defer fake.getDrainableBuildsMutex.RUnlock()
	fake.markNonInterceptibleBuildsMutex.RLock()
	defer fake.markNonInterceptibleBuildsMutex.RUnlock()
	fake.publicBuildsMutex.RLock()
	defer fake.publicBuildsMutex.RUnlock()
	fake.visibleBuildsMutex.RLock()
	defer fake.visibleBuildsMutex.RUnlock()
	fake.visibleBuildsWithTimeMutex.RLock()
	defer fake.visibleBuildsWithTimeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.BuildFactory = new(FakeBuildFactory)
