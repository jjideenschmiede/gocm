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
	"encoding/json"
)

// MessageBody is to structure the body data
type MessageBody struct {
	Messages MessageBodyMessages `json:"messages"`
}

type MessageBodyMessages struct {
	Authentication MessageBodyAuthentication `json:"authentication"`
	Msg            []MessageBodyMsg          `json:"msg"`
}

type MessageBodyAuthentication struct {
	Producttoken string `json:"producttoken"`
}

type MessageBodyMsg struct {
	AllowedChannels []string        `json:"allowedChannels"`
	From            string          `json:"from"`
	To              []MessageBodyTo `json:"to"`
	Body            MessageBodyBody `json:"body"`
}

type MessageBodyTo struct {
	Number string `json:"number"`
}

type MessageBodyBody struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// MessageReturn is to decode the json data
type MessageReturn struct {
	Details   string `json:"details"`
	ErrorCode int    `json:"errorCode"`
	Messages  []struct {
		To               string      `json:"to"`
		Status           string      `json:"status"`
		Reference        interface{} `json:"reference"`
		Parts            int         `json:"parts"`
		MessageDetails   interface{} `json:"messageDetails"`
		MessageErrorCode int         `json:"messageErrorCode"`
	} `json:"messages"`
}

// Message is to send a new message via the business messaging api
func Message(body MessageBody) (MessageReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return MessageReturn{}, err
	}

	// Config new request
	c := Config{
		Path:   "/message",
		Method: "POST",
		Body:   convert,
	}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return MessageReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode MessageReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}
