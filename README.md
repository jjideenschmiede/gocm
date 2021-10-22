# gocm

[![Go](https://github.com/jjideenschmiede/gocm/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gocm/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gocm)](https://goreportcard.com/report/github.com/jjideenschmiede/gocm) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gocm?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gocm)

With this small library it should be possible to send SMS, WhatsApp & Co messages via [cm.com](https://www.cm.com/de-de/). And, of course, to make other functions of the API usable.

## Business Messaging

In order to send one, or more messages via the Business Messages API, you can use the following function. [Here](https://www.cm.com/app/docs/en/api/business-messaging-api/1.0/index#introduction) you can find an additional description from the manufacturer.

Currently the following channels can be used: **WhatsApp, Push, RCS, Viber, SMS**

```go
// Define body
body := MessageBody{
    Messages: MessageBodyMessages{
        Authentication: MessageBodyAuthentication{
            Producttoken: "",
        },
        Msg: []MessageBodyMsg{},
    },
}

// Create a message
body.Messages.Msg = append(body.Messages.Msg, MessageBodyMsg{
    AllowedChannels: []string{"SMS"},
    From:            "Test",
    To:              []MessageBodyTo{},
    Body: MessageBodyBody{
        Type:    "auto",
        Content: "Test message",
    },
})

// Add receiver
body.Messages.Msg[0].To = append(body.Messages.Msg[0].To, MessageBodyTo{
    Number: "004941521234567",
})

// Send message
message, err := Message(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(message)
}
```
