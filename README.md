[![Go Report Card](https://goreportcard.com/badge/edgexfoundry-holding/edgex-cli)](https://goreportcard.com/report/edgexfoundry-holding/edgex-cli)

# edgex-cli

## Introduction

A command line interface to interact with EdgeX microservices. Replaces the need to manually construct complex CURL commands and/or maintain developer scripts.

```
 _____    _           __  __  _____                     _            
| ____|__| | __ _  ___\ \/ / |  ___|__  _   _ _ __   __| |_ __ _   _ 
|  _| / _` |/ _` |/ _ \\  /  | |_ / _ \| | | | '_ \ / _` | '__| | | |
| |__| (_| | (_| |  __//  \  |  _| (_) | |_| | | | | (_| | |  | |_| |
|_____\__,_|\__, |\___/_/\_\ |_|  \___/ \__,_|_| |_|\__,_|_|   \__, |
            |___/                                              |___/ 


https://www.edgexfoundry.org/

Usage:
  edgex [command]

Available Commands:
  addressable   Addressable command command
  db            Purges entire EdgeX Database. [USE WITH CAUTION]
  device        Device command
  deviceservice Device service command
  event         Event command
  help          Help about any command
  interval      Interval command
  notification  Notification command
  profile       Device profile command.
  reading       Reading command
  status        Checks the current status of each microservice.
  subscription  Subscription command
  version       Version command

Flags:
  -h, --help       help for edgex
      --no-pager   Do not pipe output into a pager.
  -u, --url        Print URL(s) used by the entered command.
  -v, --verbose    Print entire HTTP response.

Use "edgex [command] --help" for more information about a command.
```

## Installation

In order to run this tool, you will need a **running EdgeX instance** and Go 1.12 or higher installed on your machine.

* Clone the git repo:

```
$ git clone https://github.com/edgexfoundry-holding/edgex-cli
```

* Change directory:

```
$ cd edgex-cli
```

* Install the CLI:

```
$ make install
```

You can now use the CLI by entering `edgex-cli` anywhere on your machine. 


### Developers

To try out your changes, you can either build the binary or calling `go run`.

* Build and run:

```
$ make build
$ ./edgex-cli
```

* Use `go run`:

```
$ go run main.go [COMMAND]
```

* Running tests:

```
$ make test
```

This will generate a coverage.out file in the root directory of the repo which you can use to see test coverage of code by running

```
$ go tool cover -html=coverage.out
```
