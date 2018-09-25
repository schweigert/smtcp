# SMTCP

[![Coverage Status](https://coveralls.io/repos/github/schweigert/smtcp/badge.svg?branch=master)](https://coveralls.io/github/schweigert/smtcp?branch=master)

[![Build Status](https://travis-ci.org/schweigert/smtcp.svg?branch=master)](https://travis-ci.org/schweigert/smtcp)

`Shuffled Messaging Protocol over Transmission Control Protocol` or SMTCP is a library that implements exchange of messages on the standard net go language library.

This protocol allows the implementation of complete two-way RCP systems.

It can be implemented in any other language by using the basic TCP library using the protocol specification.

## Model Specification

The exchange of messages is given by `requests`.
Requests can be sent from both sides, both `client` and `server`.
Therefore, in this context, we do not really have a tie between server and client, but only a relationship between two points, making a connection.
For this reason, we call a peer connection, but not related to P2P models.

The peer can be passive or active.
An active peer is one who receives and sends messages.
A passive peer will only receive or send messages.

In the case of a passive peer, it is implemented using a single Thread, times to receive and times to send.
An active peer will have one more thread per connection to receive calls.
It is seen that the active model is implemented on the passive model, but uses a goroutine to execute the requests. In addition, the active model must have the service record it provides.

## Request Specification

A request is a sequence of bytes.

Its contents are based on 4 initial bytes, represent an unsigned integer (LittleEndian format), which means the size of the message. You can analyze by this example:

```
0007 MESSAGE
```

The request, in turn, will contain a key / value sequence, using an unsigned integer to represent the key size and an unsigned integer to represent the value.

An example of a request calling the service `foo`, with 2 params:

  - `params['bar1']='2'`
  - `params['bar2']='3'`

```
0003 foo 0002 0004 bar1 0001 2 0004 bar2 0001 3
siz1 srv sprm siz2 prm1 siz3 c siz4 prm2 siz5 c

siz1, siz2, siz3, siz4, siz5: 4 bytes, LittleEndian unsigned int. Represents the size of the following data.
srv: size1 bytes. Service name in ASCII.
sprm: 4 bytes, LittleEndian unsigned int. Represents the amount of parameters.
prmN: Param name, with SizN lengh.
C: N bytes, the content of a message.
```

A connection terminates correctly when a message with a zero-sized service name is sent:

```
0000 0000
siz1 sprm
```
