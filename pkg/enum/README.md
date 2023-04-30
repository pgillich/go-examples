# Enum

## Enum with generics

Goals:

* Base type is an [iota](https://go.dev/ref/spec#Iota)-compliant integer, no `struct`.
* String representation to each values.
* Value and string literals must be unique.
* Implements Json Marshaller and Unmarshaller interfaces.
* The zero value is reserved for missing value in Json, so the `iota` should be started from 1.
* Implements fmt Stringer and GoStringer interface to support `%s` and `%#v` formatter directives.
* No code generators, only generics.
* Thread-safe.
* Short declarations, less copypaste work.

The integer base type is a strict limitation, because the value-string mapping is unable to store beside of the type.
The solution is registering and storing all mappings in an unexported global `sync.Map`.
