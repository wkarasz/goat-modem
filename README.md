# modem

A low level Go driver for AT modems.

[![Build Status](https://travis-ci.org/wkarasz/goat-modem.svg)](https://travis-ci.org/wkarasz/goat-modem)
[![Coverage Status](https://coveralls.io/repos/github/wkarasz/goat-modem/badge.svg?branch=master)](https://coveralls.io/github/wkarasz/goat-modem?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/wkarasz/goat-modem)](https://goreportcard.com/report/github.com/wkarasz/goat-modem)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/wkarasz/goat-modem/blob/master/LICENSE)

modem is a Go library for interacting with AT based modems.
The initial impetus was to provide functionality to send and receive SMSs via
a GSM modem, but the library may be generally useful for any device controlled
by AT commands.

The [at](at) package provides a low level driver which sits between an io.ReadWriter,
representing the physical modem, and a higher level driver or application.
The AT driver provides the ability to issue AT commands to the modem, and to
receive the info and status returned by the modem, as synchronous function calls.
Handlers for asynchronous indications from the modem, such as received SMSs,
can be registered with the driver.

The [gsm](gsm) package adds higher level SendSMS and SendSMSPDU methods to the AT driver, that allows
for sending SMSs without any knowledge of the underlying AT commands.

The [info](info) package provides utility functions to manipulate the info returned in
the responses from the modem.

The [serial](serial) package provides a simple wrapper around a third party serial driver,
so you don't have to find one yourself.

The [trace](trace) package provides a driver, which may be inserted between the AT driver
and the underlying modem, to log interactions with the modem for debugging
purposes.

The [cmd](cmd) directory contains basic commands to exercise the library and a modem, including
[retrieving details](cmd/modeminfo/modeminfo.go) from the modem, [sending](cmd/sendsms/sendsms.go)
and [receiving](cmd/waitsms/waitsms.go) SMSs, and [retrieving](cmd/phonebook/phonebook.go) the SIM phonebook.

## Features

Supports the following functionality:

- Simple synchronous interface for AT commands
- Serialises access to the modem from multiple goroutines
- Context support to allow higher layers to specify timeouts
- Asynchronous indication handling
- Tracing of messages to and from the modem
- Pluggable serial driver - any io.ReadWriter will suffice

## Usage

The [at](at) package allows you to issue commands to the modem and receive the response.
e.g., this command:

```golang
info, err := modem.Command(ctx, "I")
```

produces the following interaction with the modem (exact results will differ for your modem):

    2018/05/17 20:39:56 w: ATI
    2018/05/17 20:39:56 r:
    Manufacturer: huawei
    Model: E173
    Revision: 21.017.09.00.314
    IMEI: 1234567
    +GCAP: +CGSM,+DS,+ES

    OK

and returns this info:

```golang
info = []string{"Manufacturer: huawei", "Model: E173", "Revision: 21.017.09.00.314", "IMEI: 1234567", "+GCAP: +CGSM,+DS,+ES"}
```

Refer to the [modeminfo](cmd/modeminfo/modeminfo.go) for an example of how to create a modem object such as the one used in this example.

For more information, refer to package documentation, tests and example commands.

Package | Documentation | Tests | Example code
------- | ------------- | ----- | ------------
[at](at) | [![GoDoc](https://godoc.org/github.com/wkarasz/goat-modem/at?status.svg)](https://godoc.org/github.com/wkarasz/goat-modem/at) | [at_test](at/at_test.go) | [modeminfo](cmd/modeminfo/modeminfo.go)
[gsm](gsm) | [![GoDoc](https://godoc.org/github.com/wkarasz/goat-modem/gsm?status.svg)](https://godoc.org/github.com/wkarasz/goat-modem/gsm) | [gsm_test](gsm/gsm_test.go) | [sendsms](cmd/sendsms/sendsms.go), [waitsms](cmd/waitsms/waitsms.go)
[info](info) | [![GoDoc](https://godoc.org/github.com/wkarasz/goat-modem/info?status.svg)](https://godoc.org/github.com/wkarasz/goat-modem/info) | [info_test](info/info_test.go) | [phonebook](cmd/phonebook/phonebook.go)
[serial](serial) | [![GoDoc](https://godoc.org/github.com/wkarasz/goat-modem/serial?status.svg)](https://godoc.org/github.com/wkarasz/goat-modem/serial) | [serial_test](serial/serial_test.go) | [modeminfo](cmd/modeminfo/modeminfo.go), [sendsms](cmd/sendsms/sendsms.go), [waitsms](cmd/waitsms/waitsms.go)
[trace](trace) | [![GoDoc](https://godoc.org/github.com/wkarasz/goat-modem/trace?status.svg)](https://godoc.org/github.com/wkarasz/goat-modem/trace) | [trace_test](trace/trace_test.go) | [sendsms](cmd/sendsms/sendsms.go), [waitsms](cmd/waitsms/waitsms.go)
