package database

import "errors"

type opErr struct {
	inner error
}

func (e LockedError) InnerError() error {
	if e.inner != nil {
		return e.inner
	}
	return errors.New("")
}

type LockedError opErr

func (e LockedError) Error() string {
	return "Forbidden."
}

type NotFoundError opErr

func (e NotFoundError) Error() string {
	return "Not Found."
}

type AlreadyExistsError opErr

func (e AlreadyExistsError) Error() string {
	return "Already Exists."
}

type TooBigError opErr

func (e TooBigError) Error() string {
	return "Too Big."
}

type InvalidArgsError opErr

func (e InvalidArgsError) Error() string {
	return "Invalid Arguments."
}

type BadDeltaError opErr

func (e BadDeltaError) Error() string {
	return "Bad Delta."
}

type NotSupportedError opErr

func (e NotSupportedError) Error() string {
	return "Not Supported."
}

type NotImplementedError opErr

func (e NotImplementedError) Error() string {
	return "Not Implemented."
}

type InternalError opErr

func (e InternalError) Error() string {
	return "Internal Error."
}
