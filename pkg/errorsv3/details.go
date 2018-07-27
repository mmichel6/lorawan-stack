// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

package errors

type detailer interface {
	Details() []interface{}
}

func (e *Error) addDetails(details ...interface{}) {
	if e.details == nil {
		e.details = details
	} else {
		e.details = append(e.details, details...)
	}
	if e.stack == nil {
		e.stack = callers(4)
	}
	e.clearGRPCStatus()
}

// WithDetails returns the error with the given details set.
// This appends to any existing details in the Error.
func (e Error) WithDetails(details ...interface{}) Error {
	e.addDetails(details...)
	return e
}

// WithDetails returns a new error from the definition, and sets the given details.
func (d Definition) WithDetails(details ...interface{}) Error {
	e := build(d, 0) // Don't refactor this to build(...).WithDetails(...)
	e.addDetails(details...)
	return e
}

// Details of the error. Usually structs from ttnpb or google.golang.org/genproto/googleapis/rpc/errdetails.
func (e Error) Details() (details []interface{}) {
	if e.cause != nil {
		details = append(details, Details(e.cause))
	}
	return append(details, e.details...)
}

// Details are not present in the error definition, so this just returns nil.
func (d Definition) Details() []interface{} { return nil }

// Details gets the details of the error.
func Details(err error) []interface{} {
	if c, ok := err.(detailer); ok {
		return c.Details()
	}
	if ttnErr, ok := From(err); ok {
		return ttnErr.Details()
	}
	return nil
}
