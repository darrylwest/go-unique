# Go Unique!


```
   ___                     _   _             _     __ _                   
  / __|    ___      o O O | | | |  _ _      (_)   / _` |   _  _     ___   
 | (_ |   / _ \    o      | |_| | | ' \     | |   \__, |  | +| |   / -_)  
  \___|   \___/   TS__[O]  \___/  |_||_|   _|_|_   __|_|   \_,_|   \___|  
_|"""""|_|"""""| {======|_|"""""|_|"""""|_|"""""|_|"""""|_|"""""|_|"""""| 
"`-0-0-'"`-0-0-'./o--000'"`-0-0-'"`-0-0-'"`-0-0-'"`-0-0-'"`-0-0-'"`-0-0-' 
```

_A unique id generator to create uuid, ulid and other unique numbers..._

[![build](https://travis-ci.org/darrylwest/go-unique.svg?branch=master)](https://travis-ci.org/darrylwest/go-unique/)
[![reportcard](https://goreportcard.com/badge/github.com/darrylwest/go-unique)](https://goreportcard.com/report/github.com/darrylwest/go-unique)

## Overview

A module that supports a handy collection of unique id generators including:

* ulid - a 26 character sortable
* uuid - from standard high-quality uuid
* guid - a uuid with dashes stripped

And some ids suitable for sessions, etc:

* tsid - a 12 character sortable id, good to the nano second
* txid - a 16 character sortable id, tsid with random bytes

## Installation

`go get github.com/darrylwest/go-unique/unique`

or

`git clone https://github.com/darrylwest/go-unique.git`

## Use

```go
import "github.com/darrylwest/go-unique/unique"

func main() {
    uuid := unique.CreateUUID()
    fmt.Printf("uuid : %s\n", uuid)

    ulid := unique.CreateULID()
    fmt.Printf("ulid : %s\n", ulid)
    
    txid := unique.CreateTXID()
    fmt.Printf("txid : %s\n", txid)
    
    if buf, err := unique.RandomBytes(24); err == nil {
    	str := fmt.Sprintf("%x", buf)
    	fmt.Printf("%s (%d)\n", s, len(str));
    }
}
```

### Command Line

`unique --version` // shows version

`unique [ --ulid --uuid --guid --tsid --txid --bytes ]`

Generates the specified id or byte stream.

_The make file includes a `make install` target that installs unique in /usr/local/bin/ and links ulid, uuid, guid, tsid, txid to enable invoking them without flags..._

### ULID

Universally Unique Lexicographically Sortable Identifier designed by Alizain Feerasta (similar to mongo's ObjectId).

Example: `01BFJA617JMJXEW6G7TDDXNSHX`

### UUID

Standard uuid created with rand/crypto tools.

Example: `63ba8ab3-ae69-92ac-50fc-b408876999bc`

### GUID

Standard uuid with the dashes stripped (16 bytes random)

Example: `f8ceaacb6f9145c62d9ab1e9079b8e6e`

### Time Stamp Id (TSID)

Base 36 unix nano second that generates a 12 character sortable id. 

Example: `bcodzisneczc` 

### TXID - Time Stamp with Random Bytes

Base 36 unix nano with 2 additional random bytes, 4 characters.  Sortable, short, not universal but good for sessions and local ids.

Example: `bcoexh3rdic67523`

### RandomBytes

Generates a crypto-strength random byte stream.

Example: `269f58bd8796a774e86c7f61c824d4391bbc356bcb190ef3` // 24 bytes, 48 characters

## Services

### TCP Service

Unique not includes as a separate runnable, a TCP service that responds to unique requests for uuid, ulid, guid, tsid, txid and bytes.  The service runs stand-alone or inside a container.   This is convenient when you want to centralize ID generation or use go-unique in an alternate client language.  

_The examples folder has clients in go, node, java and python._

## License

Apache 2.0

## To Do

* added http interface 
* create docker containers for tcp and http services
* implement mongodb ObjectId from mongo spec
* implement a DOI struct with id, dateCreated, lastUpdated, and version

###### darryl.west@raincitysoftware.com | Version 1.0.6
