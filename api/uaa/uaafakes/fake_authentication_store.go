// This file was generated by counterfeiter
package uaafakes

import (
	"sync"

	"code.cloudfoundry.org/cli/api/uaa"
)

type FakeAuthenticationStore struct {
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns     struct {
		result1 string
	}
	RefreshTokenStub        func() string
	refreshTokenMutex       sync.RWMutex
	refreshTokenArgsForCall []struct{}
	refreshTokenReturns     struct {
		result1 string
	}
	SetAccessTokenStub        func(token string)
	setAccessTokenMutex       sync.RWMutex
	setAccessTokenArgsForCall []struct {
		token string
	}
	ClientIDStub        func() string
	clientIDMutex       sync.RWMutex
	clientIDArgsForCall []struct{}
	clientIDReturns     struct {
		result1 string
	}
	ClientSecretStub        func() string
	clientSecretMutex       sync.RWMutex
	clientSecretArgsForCall []struct{}
	clientSecretReturns     struct {
		result1 string
	}
	SkipSSLValidationStub        func() bool
	skipSSLValidationMutex       sync.RWMutex
	skipSSLValidationArgsForCall []struct{}
	skipSSLValidationReturns     struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthenticationStore) AccessToken() string {
	fake.accessTokenMutex.Lock()
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	} else {
		return fake.accessTokenReturns.result1
	}
}

func (fake *FakeAuthenticationStore) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeAuthenticationStore) AccessTokenReturns(result1 string) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthenticationStore) RefreshToken() string {
	fake.refreshTokenMutex.Lock()
	fake.refreshTokenArgsForCall = append(fake.refreshTokenArgsForCall, struct{}{})
	fake.recordInvocation("RefreshToken", []interface{}{})
	fake.refreshTokenMutex.Unlock()
	if fake.RefreshTokenStub != nil {
		return fake.RefreshTokenStub()
	} else {
		return fake.refreshTokenReturns.result1
	}
}

func (fake *FakeAuthenticationStore) RefreshTokenCallCount() int {
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	return len(fake.refreshTokenArgsForCall)
}

func (fake *FakeAuthenticationStore) RefreshTokenReturns(result1 string) {
	fake.RefreshTokenStub = nil
	fake.refreshTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthenticationStore) SetAccessToken(token string) {
	fake.setAccessTokenMutex.Lock()
	fake.setAccessTokenArgsForCall = append(fake.setAccessTokenArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("SetAccessToken", []interface{}{token})
	fake.setAccessTokenMutex.Unlock()
	if fake.SetAccessTokenStub != nil {
		fake.SetAccessTokenStub(token)
	}
}

func (fake *FakeAuthenticationStore) SetAccessTokenCallCount() int {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return len(fake.setAccessTokenArgsForCall)
}

func (fake *FakeAuthenticationStore) SetAccessTokenArgsForCall(i int) string {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return fake.setAccessTokenArgsForCall[i].token
}

func (fake *FakeAuthenticationStore) ClientID() string {
	fake.clientIDMutex.Lock()
	fake.clientIDArgsForCall = append(fake.clientIDArgsForCall, struct{}{})
	fake.recordInvocation("ClientID", []interface{}{})
	fake.clientIDMutex.Unlock()
	if fake.ClientIDStub != nil {
		return fake.ClientIDStub()
	} else {
		return fake.clientIDReturns.result1
	}
}

func (fake *FakeAuthenticationStore) ClientIDCallCount() int {
	fake.clientIDMutex.RLock()
	defer fake.clientIDMutex.RUnlock()
	return len(fake.clientIDArgsForCall)
}

func (fake *FakeAuthenticationStore) ClientIDReturns(result1 string) {
	fake.ClientIDStub = nil
	fake.clientIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthenticationStore) ClientSecret() string {
	fake.clientSecretMutex.Lock()
	fake.clientSecretArgsForCall = append(fake.clientSecretArgsForCall, struct{}{})
	fake.recordInvocation("ClientSecret", []interface{}{})
	fake.clientSecretMutex.Unlock()
	if fake.ClientSecretStub != nil {
		return fake.ClientSecretStub()
	} else {
		return fake.clientSecretReturns.result1
	}
}

func (fake *FakeAuthenticationStore) ClientSecretCallCount() int {
	fake.clientSecretMutex.RLock()
	defer fake.clientSecretMutex.RUnlock()
	return len(fake.clientSecretArgsForCall)
}

func (fake *FakeAuthenticationStore) ClientSecretReturns(result1 string) {
	fake.ClientSecretStub = nil
	fake.clientSecretReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthenticationStore) SkipSSLValidation() bool {
	fake.skipSSLValidationMutex.Lock()
	fake.skipSSLValidationArgsForCall = append(fake.skipSSLValidationArgsForCall, struct{}{})
	fake.recordInvocation("SkipSSLValidation", []interface{}{})
	fake.skipSSLValidationMutex.Unlock()
	if fake.SkipSSLValidationStub != nil {
		return fake.SkipSSLValidationStub()
	} else {
		return fake.skipSSLValidationReturns.result1
	}
}

func (fake *FakeAuthenticationStore) SkipSSLValidationCallCount() int {
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	return len(fake.skipSSLValidationArgsForCall)
}

func (fake *FakeAuthenticationStore) SkipSSLValidationReturns(result1 bool) {
	fake.SkipSSLValidationStub = nil
	fake.skipSSLValidationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAuthenticationStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	fake.clientIDMutex.RLock()
	defer fake.clientIDMutex.RUnlock()
	fake.clientSecretMutex.RLock()
	defer fake.clientSecretMutex.RUnlock()
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAuthenticationStore) recordInvocation(key string, args []interface{}) {
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

var _ uaa.AuthenticationStore = new(FakeAuthenticationStore)