# NATS Hashing Service
Proof-of-Concept string hashing service using the [NATS micro](https://pkg.go.dev/github.com/nats-io/nats.go/micro#readme-overview) package.

## Endpoints
- `hash.md5`
- `hash.sha1`
- `hash.sha256`
- `hash.sha512`
- `hash.blake2b`

## Usage

1. Install NATS CLI: https://github.com/nats-io/natscli#installation

2. Add and select demo NATS server:

    `nats context add demo --server demo.nats.io:4222 --select`

3. Check for running instances:

    `nats micro list HashService`

4. Run your own instance if necessary:
    - With [Task](https://taskfile.dev): `task run`
    - Directly: `go run ./cmd`

5. Send request:

    `nats req hash.md5 "hello there"`
    > Output: `161bc25962da8fed6d2f59922fb642aa`
