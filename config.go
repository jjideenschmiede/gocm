//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gocm.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gocm

import (
	"bytes"
	"net/http"
)

const (
	baseUrl = "https://gw.cmtelecom.com/v1.0"
)

// Config is to define config data
type Config struct {
	Path, Method string
	Body         []byte
}

// Send is to send a new request
func (c Config) Send() (*http.Response, error) {

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, baseUrl+c.Path, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Add header
	request.Header.Add("Content-Type", "application/json")

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
