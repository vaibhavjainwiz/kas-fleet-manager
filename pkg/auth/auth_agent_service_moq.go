// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"sync"
)

// Ensure, that AuthAgentServiceMock does implement AuthAgentService.
// If this is not the case, regenerate this file with moq.
var _ AuthAgentService = &AuthAgentServiceMock{}

// AuthAgentServiceMock is a mock implementation of AuthAgentService.
//
// 	func TestSomethingThatUsesAuthAgentService(t *testing.T) {
//
// 		// make and configure a mocked AuthAgentService
// 		mockedAuthAgentService := &AuthAgentServiceMock{
// 			GetClientIdFunc: func(clusterId string) (string, error) {
// 				panic("mock out the GetClientId method")
// 			},
// 		}
//
// 		// use mockedAuthAgentService in code that requires AuthAgentService
// 		// and then make assertions.
//
// 	}
type AuthAgentServiceMock struct {
	// GetClientIdFunc mocks the GetClientId method.
	GetClientIdFunc func(clusterId string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetClientId holds details about calls to the GetClientId method.
		GetClientId []struct {
			// ClusterId is the clusterId argument value.
			ClusterId string
		}
	}
	lockGetClientId sync.RWMutex
}

// GetClientId calls GetClientIdFunc.
func (mock *AuthAgentServiceMock) GetClientId(clusterId string) (string, error) {
	if mock.GetClientIdFunc == nil {
		panic("AuthAgentServiceMock.GetClientIdFunc: method is nil but AuthAgentService.GetClientId was just called")
	}
	callInfo := struct {
		ClusterId string
	}{
		ClusterId: clusterId,
	}
	mock.lockGetClientId.Lock()
	mock.calls.GetClientId = append(mock.calls.GetClientId, callInfo)
	mock.lockGetClientId.Unlock()
	return mock.GetClientIdFunc(clusterId)
}

// GetClientIdCalls gets all the calls that were made to GetClientId.
// Check the length with:
//     len(mockedAuthAgentService.GetClientIdCalls())
func (mock *AuthAgentServiceMock) GetClientIdCalls() []struct {
	ClusterId string
} {
	var calls []struct {
		ClusterId string
	}
	mock.lockGetClientId.RLock()
	calls = mock.calls.GetClientId
	mock.lockGetClientId.RUnlock()
	return calls
}