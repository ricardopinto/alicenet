// Code generated by go-mockgen 1.3.3; DO NOT EDIT.

package mocks

import (
	"context"
	"sync"

	executor "github.com/alicenet/alicenet/layer1/executor"
	tasks "github.com/alicenet/alicenet/layer1/executor/tasks"
)

// MockTaskHandler is a mock implementation of the TaskHandler interface
// (from the package github.com/alicenet/alicenet/layer1/executor) used for
// unit testing.
type MockTaskHandler struct {
	// CloseFunc is an instance of a mock function object controlling the
	// behavior of the method Close.
	CloseFunc *TaskHandlerCloseFunc
	// KillTaskByIdFunc is an instance of a mock function object controlling
	// the behavior of the method KillTaskById.
	KillTaskByIdFunc *TaskHandlerKillTaskByIdFunc
	// KillTaskByTypeFunc is an instance of a mock function object
	// controlling the behavior of the method KillTaskByType.
	KillTaskByTypeFunc *TaskHandlerKillTaskByTypeFunc
	// ScheduleTaskFunc is an instance of a mock function object controlling
	// the behavior of the method ScheduleTask.
	ScheduleTaskFunc *TaskHandlerScheduleTaskFunc
	// StartFunc is an instance of a mock function object controlling the
	// behavior of the method Start.
	StartFunc *TaskHandlerStartFunc
}

// NewMockTaskHandler creates a new mock of the TaskHandler interface. All
// methods return zero values for all results, unless overwritten.
func NewMockTaskHandler() *MockTaskHandler {
	return &MockTaskHandler{
		CloseFunc: &TaskHandlerCloseFunc{
			defaultHook: func() {
				return
			},
		},
		KillTaskByIdFunc: &TaskHandlerKillTaskByIdFunc{
			defaultHook: func(context.Context, string) (r0 *executor.HandlerResponse, r1 error) {
				return
			},
		},
		KillTaskByTypeFunc: &TaskHandlerKillTaskByTypeFunc{
			defaultHook: func(context.Context, tasks.Task) (r0 *executor.HandlerResponse, r1 error) {
				return
			},
		},
		ScheduleTaskFunc: &TaskHandlerScheduleTaskFunc{
			defaultHook: func(context.Context, tasks.Task, string) (r0 *executor.HandlerResponse, r1 error) {
				return
			},
		},
		StartFunc: &TaskHandlerStartFunc{
			defaultHook: func() {
				return
			},
		},
	}
}

// NewStrictMockTaskHandler creates a new mock of the TaskHandler interface.
// All methods panic on invocation, unless overwritten.
func NewStrictMockTaskHandler() *MockTaskHandler {
	return &MockTaskHandler{
		CloseFunc: &TaskHandlerCloseFunc{
			defaultHook: func() {
				panic("unexpected invocation of MockTaskHandler.Close")
			},
		},
		KillTaskByIdFunc: &TaskHandlerKillTaskByIdFunc{
			defaultHook: func(context.Context, string) (*executor.HandlerResponse, error) {
				panic("unexpected invocation of MockTaskHandler.KillTaskById")
			},
		},
		KillTaskByTypeFunc: &TaskHandlerKillTaskByTypeFunc{
			defaultHook: func(context.Context, tasks.Task) (*executor.HandlerResponse, error) {
				panic("unexpected invocation of MockTaskHandler.KillTaskByType")
			},
		},
		ScheduleTaskFunc: &TaskHandlerScheduleTaskFunc{
			defaultHook: func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error) {
				panic("unexpected invocation of MockTaskHandler.ScheduleTask")
			},
		},
		StartFunc: &TaskHandlerStartFunc{
			defaultHook: func() {
				panic("unexpected invocation of MockTaskHandler.Start")
			},
		},
	}
}

// NewMockTaskHandlerFrom creates a new mock of the MockTaskHandler
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockTaskHandlerFrom(i executor.TaskHandler) *MockTaskHandler {
	return &MockTaskHandler{
		CloseFunc: &TaskHandlerCloseFunc{
			defaultHook: i.Close,
		},
		KillTaskByIdFunc: &TaskHandlerKillTaskByIdFunc{
			defaultHook: i.KillTaskById,
		},
		KillTaskByTypeFunc: &TaskHandlerKillTaskByTypeFunc{
			defaultHook: i.KillTaskByType,
		},
		ScheduleTaskFunc: &TaskHandlerScheduleTaskFunc{
			defaultHook: i.ScheduleTask,
		},
		StartFunc: &TaskHandlerStartFunc{
			defaultHook: i.Start,
		},
	}
}

// TaskHandlerCloseFunc describes the behavior when the Close method of the
// parent MockTaskHandler instance is invoked.
type TaskHandlerCloseFunc struct {
	defaultHook func()
	hooks       []func()
	history     []TaskHandlerCloseFuncCall
	mutex       sync.Mutex
}

// Close delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockTaskHandler) Close() {
	m.CloseFunc.nextHook()()
	m.CloseFunc.appendCall(TaskHandlerCloseFuncCall{})
	return
}

// SetDefaultHook sets function that is called when the Close method of the
// parent MockTaskHandler instance is invoked and the hook queue is empty.
func (f *TaskHandlerCloseFunc) SetDefaultHook(hook func()) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Close method of the parent MockTaskHandler instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *TaskHandlerCloseFunc) PushHook(hook func()) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *TaskHandlerCloseFunc) SetDefaultReturn() {
	f.SetDefaultHook(func() {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *TaskHandlerCloseFunc) PushReturn() {
	f.PushHook(func() {
		return
	})
}

func (f *TaskHandlerCloseFunc) nextHook() func() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *TaskHandlerCloseFunc) appendCall(r0 TaskHandlerCloseFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of TaskHandlerCloseFuncCall objects describing
// the invocations of this function.
func (f *TaskHandlerCloseFunc) History() []TaskHandlerCloseFuncCall {
	f.mutex.Lock()
	history := make([]TaskHandlerCloseFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// TaskHandlerCloseFuncCall is an object that describes an invocation of
// method Close on an instance of MockTaskHandler.
type TaskHandlerCloseFuncCall struct{}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c TaskHandlerCloseFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c TaskHandlerCloseFuncCall) Results() []interface{} {
	return []interface{}{}
}

// TaskHandlerKillTaskByIdFunc describes the behavior when the KillTaskById
// method of the parent MockTaskHandler instance is invoked.
type TaskHandlerKillTaskByIdFunc struct {
	defaultHook func(context.Context, string) (*executor.HandlerResponse, error)
	hooks       []func(context.Context, string) (*executor.HandlerResponse, error)
	history     []TaskHandlerKillTaskByIdFuncCall
	mutex       sync.Mutex
}

// KillTaskById delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockTaskHandler) KillTaskById(v0 context.Context, v1 string) (*executor.HandlerResponse, error) {
	r0, r1 := m.KillTaskByIdFunc.nextHook()(v0, v1)
	m.KillTaskByIdFunc.appendCall(TaskHandlerKillTaskByIdFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the KillTaskById method
// of the parent MockTaskHandler instance is invoked and the hook queue is
// empty.
func (f *TaskHandlerKillTaskByIdFunc) SetDefaultHook(hook func(context.Context, string) (*executor.HandlerResponse, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// KillTaskById method of the parent MockTaskHandler instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *TaskHandlerKillTaskByIdFunc) PushHook(hook func(context.Context, string) (*executor.HandlerResponse, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *TaskHandlerKillTaskByIdFunc) SetDefaultReturn(r0 *executor.HandlerResponse, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*executor.HandlerResponse, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *TaskHandlerKillTaskByIdFunc) PushReturn(r0 *executor.HandlerResponse, r1 error) {
	f.PushHook(func(context.Context, string) (*executor.HandlerResponse, error) {
		return r0, r1
	})
}

func (f *TaskHandlerKillTaskByIdFunc) nextHook() func(context.Context, string) (*executor.HandlerResponse, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *TaskHandlerKillTaskByIdFunc) appendCall(r0 TaskHandlerKillTaskByIdFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of TaskHandlerKillTaskByIdFuncCall objects
// describing the invocations of this function.
func (f *TaskHandlerKillTaskByIdFunc) History() []TaskHandlerKillTaskByIdFuncCall {
	f.mutex.Lock()
	history := make([]TaskHandlerKillTaskByIdFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// TaskHandlerKillTaskByIdFuncCall is an object that describes an invocation
// of method KillTaskById on an instance of MockTaskHandler.
type TaskHandlerKillTaskByIdFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *executor.HandlerResponse
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c TaskHandlerKillTaskByIdFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c TaskHandlerKillTaskByIdFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// TaskHandlerKillTaskByTypeFunc describes the behavior when the
// KillTaskByType method of the parent MockTaskHandler instance is invoked.
type TaskHandlerKillTaskByTypeFunc struct {
	defaultHook func(context.Context, tasks.Task) (*executor.HandlerResponse, error)
	hooks       []func(context.Context, tasks.Task) (*executor.HandlerResponse, error)
	history     []TaskHandlerKillTaskByTypeFuncCall
	mutex       sync.Mutex
}

// KillTaskByType delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockTaskHandler) KillTaskByType(v0 context.Context, v1 tasks.Task) (*executor.HandlerResponse, error) {
	r0, r1 := m.KillTaskByTypeFunc.nextHook()(v0, v1)
	m.KillTaskByTypeFunc.appendCall(TaskHandlerKillTaskByTypeFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the KillTaskByType
// method of the parent MockTaskHandler instance is invoked and the hook
// queue is empty.
func (f *TaskHandlerKillTaskByTypeFunc) SetDefaultHook(hook func(context.Context, tasks.Task) (*executor.HandlerResponse, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// KillTaskByType method of the parent MockTaskHandler instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *TaskHandlerKillTaskByTypeFunc) PushHook(hook func(context.Context, tasks.Task) (*executor.HandlerResponse, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *TaskHandlerKillTaskByTypeFunc) SetDefaultReturn(r0 *executor.HandlerResponse, r1 error) {
	f.SetDefaultHook(func(context.Context, tasks.Task) (*executor.HandlerResponse, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *TaskHandlerKillTaskByTypeFunc) PushReturn(r0 *executor.HandlerResponse, r1 error) {
	f.PushHook(func(context.Context, tasks.Task) (*executor.HandlerResponse, error) {
		return r0, r1
	})
}

func (f *TaskHandlerKillTaskByTypeFunc) nextHook() func(context.Context, tasks.Task) (*executor.HandlerResponse, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *TaskHandlerKillTaskByTypeFunc) appendCall(r0 TaskHandlerKillTaskByTypeFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of TaskHandlerKillTaskByTypeFuncCall objects
// describing the invocations of this function.
func (f *TaskHandlerKillTaskByTypeFunc) History() []TaskHandlerKillTaskByTypeFuncCall {
	f.mutex.Lock()
	history := make([]TaskHandlerKillTaskByTypeFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// TaskHandlerKillTaskByTypeFuncCall is an object that describes an
// invocation of method KillTaskByType on an instance of MockTaskHandler.
type TaskHandlerKillTaskByTypeFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 tasks.Task
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *executor.HandlerResponse
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c TaskHandlerKillTaskByTypeFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c TaskHandlerKillTaskByTypeFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// TaskHandlerScheduleTaskFunc describes the behavior when the ScheduleTask
// method of the parent MockTaskHandler instance is invoked.
type TaskHandlerScheduleTaskFunc struct {
	defaultHook func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error)
	hooks       []func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error)
	history     []TaskHandlerScheduleTaskFuncCall
	mutex       sync.Mutex
}

// ScheduleTask delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockTaskHandler) ScheduleTask(v0 context.Context, v1 tasks.Task, v2 string) (*executor.HandlerResponse, error) {
	r0, r1 := m.ScheduleTaskFunc.nextHook()(v0, v1, v2)
	m.ScheduleTaskFunc.appendCall(TaskHandlerScheduleTaskFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the ScheduleTask method
// of the parent MockTaskHandler instance is invoked and the hook queue is
// empty.
func (f *TaskHandlerScheduleTaskFunc) SetDefaultHook(hook func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ScheduleTask method of the parent MockTaskHandler instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *TaskHandlerScheduleTaskFunc) PushHook(hook func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *TaskHandlerScheduleTaskFunc) SetDefaultReturn(r0 *executor.HandlerResponse, r1 error) {
	f.SetDefaultHook(func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *TaskHandlerScheduleTaskFunc) PushReturn(r0 *executor.HandlerResponse, r1 error) {
	f.PushHook(func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error) {
		return r0, r1
	})
}

func (f *TaskHandlerScheduleTaskFunc) nextHook() func(context.Context, tasks.Task, string) (*executor.HandlerResponse, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *TaskHandlerScheduleTaskFunc) appendCall(r0 TaskHandlerScheduleTaskFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of TaskHandlerScheduleTaskFuncCall objects
// describing the invocations of this function.
func (f *TaskHandlerScheduleTaskFunc) History() []TaskHandlerScheduleTaskFuncCall {
	f.mutex.Lock()
	history := make([]TaskHandlerScheduleTaskFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// TaskHandlerScheduleTaskFuncCall is an object that describes an invocation
// of method ScheduleTask on an instance of MockTaskHandler.
type TaskHandlerScheduleTaskFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 tasks.Task
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *executor.HandlerResponse
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c TaskHandlerScheduleTaskFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c TaskHandlerScheduleTaskFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// TaskHandlerStartFunc describes the behavior when the Start method of the
// parent MockTaskHandler instance is invoked.
type TaskHandlerStartFunc struct {
	defaultHook func()
	hooks       []func()
	history     []TaskHandlerStartFuncCall
	mutex       sync.Mutex
}

// Start delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockTaskHandler) Start() {
	m.StartFunc.nextHook()()
	m.StartFunc.appendCall(TaskHandlerStartFuncCall{})
	return
}

// SetDefaultHook sets function that is called when the Start method of the
// parent MockTaskHandler instance is invoked and the hook queue is empty.
func (f *TaskHandlerStartFunc) SetDefaultHook(hook func()) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Start method of the parent MockTaskHandler instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *TaskHandlerStartFunc) PushHook(hook func()) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *TaskHandlerStartFunc) SetDefaultReturn() {
	f.SetDefaultHook(func() {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *TaskHandlerStartFunc) PushReturn() {
	f.PushHook(func() {
		return
	})
}

func (f *TaskHandlerStartFunc) nextHook() func() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *TaskHandlerStartFunc) appendCall(r0 TaskHandlerStartFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of TaskHandlerStartFuncCall objects describing
// the invocations of this function.
func (f *TaskHandlerStartFunc) History() []TaskHandlerStartFuncCall {
	f.mutex.Lock()
	history := make([]TaskHandlerStartFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// TaskHandlerStartFuncCall is an object that describes an invocation of
// method Start on an instance of MockTaskHandler.
type TaskHandlerStartFuncCall struct{}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c TaskHandlerStartFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c TaskHandlerStartFuncCall) Results() []interface{} {
	return []interface{}{}
}

// MockTaskResponse is a mock implementation of the TaskResponse interface
// (from the package github.com/alicenet/alicenet/layer1/executor) used for
// unit testing.
type MockTaskResponse struct {
	// GetResponseBlockingFunc is an instance of a mock function object
	// controlling the behavior of the method GetResponseBlocking.
	GetResponseBlockingFunc *TaskResponseGetResponseBlockingFunc
	// IsReadyFunc is an instance of a mock function object controlling the
	// behavior of the method IsReady.
	IsReadyFunc *TaskResponseIsReadyFunc
}

// NewMockTaskResponse creates a new mock of the TaskResponse interface. All
// methods return zero values for all results, unless overwritten.
func NewMockTaskResponse() *MockTaskResponse {
	return &MockTaskResponse{
		GetResponseBlockingFunc: &TaskResponseGetResponseBlockingFunc{
			defaultHook: func(context.Context) (r0 error) {
				return
			},
		},
		IsReadyFunc: &TaskResponseIsReadyFunc{
			defaultHook: func() (r0 bool) {
				return
			},
		},
	}
}

// NewStrictMockTaskResponse creates a new mock of the TaskResponse
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockTaskResponse() *MockTaskResponse {
	return &MockTaskResponse{
		GetResponseBlockingFunc: &TaskResponseGetResponseBlockingFunc{
			defaultHook: func(context.Context) error {
				panic("unexpected invocation of MockTaskResponse.GetResponseBlocking")
			},
		},
		IsReadyFunc: &TaskResponseIsReadyFunc{
			defaultHook: func() bool {
				panic("unexpected invocation of MockTaskResponse.IsReady")
			},
		},
	}
}

// NewMockTaskResponseFrom creates a new mock of the MockTaskResponse
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockTaskResponseFrom(i executor.TaskResponse) *MockTaskResponse {
	return &MockTaskResponse{
		GetResponseBlockingFunc: &TaskResponseGetResponseBlockingFunc{
			defaultHook: i.GetResponseBlocking,
		},
		IsReadyFunc: &TaskResponseIsReadyFunc{
			defaultHook: i.IsReady,
		},
	}
}

// TaskResponseGetResponseBlockingFunc describes the behavior when the
// GetResponseBlocking method of the parent MockTaskResponse instance is
// invoked.
type TaskResponseGetResponseBlockingFunc struct {
	defaultHook func(context.Context) error
	hooks       []func(context.Context) error
	history     []TaskResponseGetResponseBlockingFuncCall
	mutex       sync.Mutex
}

// GetResponseBlocking delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockTaskResponse) GetResponseBlocking(v0 context.Context) error {
	r0 := m.GetResponseBlockingFunc.nextHook()(v0)
	m.GetResponseBlockingFunc.appendCall(TaskResponseGetResponseBlockingFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the GetResponseBlocking
// method of the parent MockTaskResponse instance is invoked and the hook
// queue is empty.
func (f *TaskResponseGetResponseBlockingFunc) SetDefaultHook(hook func(context.Context) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetResponseBlocking method of the parent MockTaskResponse instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *TaskResponseGetResponseBlockingFunc) PushHook(hook func(context.Context) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *TaskResponseGetResponseBlockingFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *TaskResponseGetResponseBlockingFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context) error {
		return r0
	})
}

func (f *TaskResponseGetResponseBlockingFunc) nextHook() func(context.Context) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *TaskResponseGetResponseBlockingFunc) appendCall(r0 TaskResponseGetResponseBlockingFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of TaskResponseGetResponseBlockingFuncCall
// objects describing the invocations of this function.
func (f *TaskResponseGetResponseBlockingFunc) History() []TaskResponseGetResponseBlockingFuncCall {
	f.mutex.Lock()
	history := make([]TaskResponseGetResponseBlockingFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// TaskResponseGetResponseBlockingFuncCall is an object that describes an
// invocation of method GetResponseBlocking on an instance of
// MockTaskResponse.
type TaskResponseGetResponseBlockingFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c TaskResponseGetResponseBlockingFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c TaskResponseGetResponseBlockingFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// TaskResponseIsReadyFunc describes the behavior when the IsReady method of
// the parent MockTaskResponse instance is invoked.
type TaskResponseIsReadyFunc struct {
	defaultHook func() bool
	hooks       []func() bool
	history     []TaskResponseIsReadyFuncCall
	mutex       sync.Mutex
}

// IsReady delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockTaskResponse) IsReady() bool {
	r0 := m.IsReadyFunc.nextHook()()
	m.IsReadyFunc.appendCall(TaskResponseIsReadyFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the IsReady method of
// the parent MockTaskResponse instance is invoked and the hook queue is
// empty.
func (f *TaskResponseIsReadyFunc) SetDefaultHook(hook func() bool) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// IsReady method of the parent MockTaskResponse instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *TaskResponseIsReadyFunc) PushHook(hook func() bool) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *TaskResponseIsReadyFunc) SetDefaultReturn(r0 bool) {
	f.SetDefaultHook(func() bool {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *TaskResponseIsReadyFunc) PushReturn(r0 bool) {
	f.PushHook(func() bool {
		return r0
	})
}

func (f *TaskResponseIsReadyFunc) nextHook() func() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *TaskResponseIsReadyFunc) appendCall(r0 TaskResponseIsReadyFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of TaskResponseIsReadyFuncCall objects
// describing the invocations of this function.
func (f *TaskResponseIsReadyFunc) History() []TaskResponseIsReadyFuncCall {
	f.mutex.Lock()
	history := make([]TaskResponseIsReadyFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// TaskResponseIsReadyFuncCall is an object that describes an invocation of
// method IsReady on an instance of MockTaskResponse.
type TaskResponseIsReadyFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 bool
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c TaskResponseIsReadyFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c TaskResponseIsReadyFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
