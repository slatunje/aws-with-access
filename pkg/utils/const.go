// Copyright © 2018 Sylvester La-Tunje. All rights reserved.

package utils

// ExitXXX represents various exit code within the system
const (
	_                                                 = iota
	ExitExecute
	ExitRequireKeys
	ExitCredentialsFailure
	ExitCommandlineFailure
	ExitShareConfigFailure
	ExitBase64DecodeFailure
	ExitOnDebug
)
