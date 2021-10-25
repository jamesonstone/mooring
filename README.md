# mooring

A Dockerfile Scanner for Potential Runtime Vulnerabilities

## Getting Started

Mooring scans each Dockerfile statically and at runtime to detect potential vulnerabilities.

### Static

Mooring performs static analysis on a Dockerfile using string matching and regular expressions, and a policy document to
determine whether a vulnerability is present.

### Runtime

Mooring preforms runtime analysis but buildings and running the container image

## Installation

```bash
go get [...]
```

## Support

J.Stone 2021
