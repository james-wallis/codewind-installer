/*******************************************************************************
 * Copyright (c) 2019 IBM Corporation and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 *     IBM Corporation - initial API and implementation
 *******************************************************************************/

package connections

import (
	"encoding/json"
)

// ConError : Connection package errors
type ConError struct {
	Op   string
	Err  error
	Desc string
}

const (
	errOpFileParse    = "con_parse"
	errOpFileLoad     = "con_load"
	errOpFileWrite    = "con_write"
	errOpSchemaUpdate = "con_schema_update"
	errOpConflict     = "con_conflict"
	errOpNotFound     = "con_not_found"
	errOpProtected    = "con_protected"
	errOpGetEnv       = "con_environment"
	errOpKeyring      = "con_keyring"
)

const (
	errTargetNotFound = "Target connection not found"
)

// ConError : Error formatted in JSON containing an errorOp and a description from
// either a fault condition in the CLI, or an error payload from a REST request
func (se *ConError) Error() string {
	type Output struct {
		Operation   string `json:"error"`
		Description string `json:"error_description"`
	}
	tempOutput := &Output{Operation: se.Op, Description: se.Err.Error()}
	jsonError, _ := json.Marshal(tempOutput)
	return string(jsonError)
}

// Result : status message
type Result struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
}
