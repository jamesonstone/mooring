# mooring

A Dockerfile Scanner for Potential Runtime Vulnerabilities

## Getting Started

Mooring scans each Dockerfile statically and at runtime to detect potential vulnerabilities.

Mooring implements a deny policy that allows you to block certain packages from being included in the final container image

Append to the DENY.txt file:

```text
bash
curl
git
ssh
```

### Static

Mooring performs static analysis on a Dockerfile using string matching and regular expressions, and a policy document to
determine whether a vulnerability is present.

### Runtime

Mooring preforms runtime analysis but buildings and running the container image

## Installation

```bash
go get [...]
```

## TODO

Assess integration of static analysis with [buildkit's Dockerfile parser](https://pkg.go.dev/github.com/moby/buildkit/frontend/dockerfile/parser#Parse)

## Support

J.Stone 2021
