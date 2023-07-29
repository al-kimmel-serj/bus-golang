# Event bus golang implementation

[![GoDoc](https://pkg.go.dev/badge/al-kimmel-serj/bus-golang)](https://pkg.go.dev/github.com/al-kimmel-serj/bus-golang)

Golang event bus library with service discovery and protobuf support. This repository provides the basic types required to implement event bus pattern.

## Publisher implementations

* [AWS Kinesis](https://github.com/al-kimmel-serj/bus-golang-publisher-aws-kinesis)
* [ZeroMQ](https://github.com/al-kimmel-serj/bus-golang-publisher-zmq)

## Publishers registry implementations

* [AWS Service Discovery aka CloudMap](https://github.com/al-kimmel-serj/bus-golang-publishers-registry-aws-servicediscovery)
* [Consul](https://github.com/al-kimmel-serj/bus-golang-publishers-registry-consul)
* [Stub implementation for unit tests](https://github.com/al-kimmel-serj/bus-golang-publishers-registry-stub)

## Subscriber implementations

* [ZeroMQ](https://github.com/al-kimmel-serj/bus-golang-subscriber-zmq)
