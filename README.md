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

## Overview

This is a handy collection of unique id generators including:

* ulid - a 26 character sortable
* uuid - from Hashicorp's repo for standard high-quality uuid
* guid - a uuid with dashes stripped

And some ids suitable for sessions, etc:

* tsid - a 12 character sortable id, good to the nano second
* txid - a 16 character sortable id, tsid with random bytes

## Installation

`go get github.com/darrylwest/go-unique`

## Use

```go
import "github.com/darrylwest/go-unique"

func main() {
	uuid := unique.CreateUUID()
   fmt.Printf("uuid : %s\n", uuid)

	ulid := unique.CreateULID()
   fmt.Printf("ulid : %s\n", uLid)
}
```

### ULID

Universally Unique Lexicographically Sortable Identifier designed by Alizain Feerasta (similar to mongo's ObjectId).

Example: `01BFJA617JMJXEW6G7TDDXNSHX`

### UUID

Standard V4 uuid from the Hashicorp go-uuid repo (not the deprecated one).

Example: `63ba8ab3-ae69-92ac-50fc-b408876999bc`

### GUID

Standard uuid with the dashes stripped (16 bytes random)

Example: `f8ceaacb6f9145c62d9ab1e9079b8e6e`

### Time Stamp Id (TSID)

Base 36 unix nano second that generates a 12 character sortable id. 

Example: `bcodzisneczc` 

### TXID - Time Stamp with Random Bytes

Bas36 unix nano with 4 additional random bytes.

Example: `bcoe5fkmvgox8723`

###### darryl.west@raincitysoftware.com | Version 1.0.0
