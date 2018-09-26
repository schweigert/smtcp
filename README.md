[![Coverage Status](https://coveralls.io/repos/github/schweigert/smtcp/badge.svg?branch=master)](https://coveralls.io/github/schweigert/smtcp?branch=master)
[![Build Status](https://travis-ci.org/schweigert/smtcp.svg?branch=master)](https://travis-ci.org/schweigert/smtcp)

![alt text](logo.png)

# SMTCP

`Shuffled Messaging Protocol over Transmission Control Protocol` or SMTCP implements two-way RPC over net.Conn.

```go
package main

import "github.com/schweigert/smtcp"

func main() {
	listener, _ := smtcp.NewTcpActiveListener(
		"3030",
		smtcp.NewLambdaSet().Set("conc", func(r *smtcp.Request) {
			p := smtcp.NewParams()
			conc := r.Params.Get("s1") + r.Params.Get("s2")

			p.Set("anw", conc)
			smtcp.NewRequest("conc_w", p, r.Peer).Send()
		}))

	for {
		peer := listener.Accept()
		peer.Work()
	}
}
```
