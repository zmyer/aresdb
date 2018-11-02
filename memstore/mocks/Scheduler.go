//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
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

// Code generated by mockery v1.0.0
package mocks

import "github.com/uber/aresdb/memstore/common"
import (
	"github.com/stretchr/testify/mock"
	"github.com/uber/aresdb/memstore"
)

// Scheduler is an autogenerated mock type for the Scheduler type
type Scheduler struct {
	mock.Mock
}

// DeleteTable provides a mock function with given fields: table, isFactTable
func (_m *Scheduler) DeleteTable(table string, isFactTable bool) {
	_m.Called(table, isFactTable)
}

// GetJobDetails provides a mock function with given fields: jobType
func (_m *Scheduler) GetJobDetails(jobType common.JobType) interface{} {
	ret := _m.Called(jobType)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(common.JobType) interface{}); ok {
		r0 = rf(jobType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Lock provides a mock function with given fields:
func (_m *Scheduler) Lock() {
	_m.Called()
}

// NewArchivingJob provides a mock function with given fields: tableName, shardID, cutoff
func (_m *Scheduler) NewArchivingJob(tableName string, shardID int, cutoff uint32) memstore.Job {
	ret := _m.Called(tableName, shardID, cutoff)

	var r0 memstore.Job
	if rf, ok := ret.Get(0).(func(string, int, uint32) memstore.Job); ok {
		r0 = rf(tableName, shardID, cutoff)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memstore.Job)
		}
	}

	return r0
}

// NewBackfillJob provides a mock function with given fields: tableName, shardID
func (_m *Scheduler) NewBackfillJob(tableName string, shardID int) memstore.Job {
	ret := _m.Called(tableName, shardID)

	var r0 memstore.Job
	if rf, ok := ret.Get(0).(func(string, int) memstore.Job); ok {
		r0 = rf(tableName, shardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memstore.Job)
		}
	}

	return r0
}

// NewPurgeJob provides a mock function with given fields: tableName, shardID, batchIDStart, batchIDEnd
func (_m *Scheduler) NewPurgeJob(tableName string, shardID int, batchIDStart int, batchIDEnd int) memstore.Job {
	ret := _m.Called(tableName, shardID, batchIDStart, batchIDEnd)

	var r0 memstore.Job
	if rf, ok := ret.Get(0).(func(string, int, int, int) memstore.Job); ok {
		r0 = rf(tableName, shardID, batchIDStart, batchIDEnd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memstore.Job)
		}
	}

	return r0
}

// NewSnapshotJob provides a mock function with given fields: tableName, shardID
func (_m *Scheduler) NewSnapshotJob(tableName string, shardID int) memstore.Job {
	ret := _m.Called(tableName, shardID)

	var r0 memstore.Job
	if rf, ok := ret.Get(0).(func(string, int) memstore.Job); ok {
		r0 = rf(tableName, shardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memstore.Job)
		}
	}

	return r0
}

// RLock provides a mock function with given fields:
func (_m *Scheduler) RLock() {
	_m.Called()
}

// RUnlock provides a mock function with given fields:
func (_m *Scheduler) RUnlock() {
	_m.Called()
}

// Start provides a mock function with given fields:
func (_m *Scheduler) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *Scheduler) Stop() {
	_m.Called()
}

// SubmitJob provides a mock function with given fields: job
func (_m *Scheduler) SubmitJob(job memstore.Job) chan error {
	ret := _m.Called(job)

	var r0 chan error
	if rf, ok := ret.Get(0).(func(memstore.Job) chan error); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan error)
		}
	}

	return r0
}

// Unlock provides a mock function with given fields:
func (_m *Scheduler) Unlock() {
	_m.Called()
}