# Protocol Buffers (PB) Directory

The `pb` directory contains Protocol Buffers (protobuf) definitions for our Golang-based monorepo project. Protocol
Buffers are used for efficient data serialization and communication between different components of the project.

## Directory Structure

- **proto**: Contains the protobuf message definitions (.proto files).
- **generated**: Generated Go code from the protobuf definitions.

## Usage

### 1. Define Protobuf Messages

Create your protobuf message definitions in the `proto` directory. Be sure to follow protobuf syntax and best practices
for defining messages and services.

Example:

```protobuf
syntax = "proto3";

package myproject;

message MyMessage {
    string field1 = 1;
    int32 field2 = 2;
}
```
