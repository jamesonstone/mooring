# mooring

A customizable Dockerfile Scanner for Potential Runtime Vulnerabilities

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

Use `#` or don't include them in the `DENY.txt` to temporarily allow a package

```text
bash
curl
git
#ssh
```

## How to Use

Check out the `main.go` file for an example of how to call Mooring.

### Static

Mooring performs static analysis on a Dockerfile using string matching in conjunction with a policy document to
determine whether a vulnerability is present.

### Runtime

Mooring preforms runtime analysis via docker scan on built container images

## Installation

```bash
go get [...]
```

## TODO

- Assess integration of static analysis with [buildkit's Dockerfile parser](https://pkg.go.dev/github.com/moby/buildkit/frontend/dockerfile/parser#Parse)
- Check for final container image user

## Support

J.Stone 2021
