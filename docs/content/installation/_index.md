---
title: "Installation"
date: 2019-03-03T16:39:46+01:00
weight: 1
draft: false
---

## Binaries

To get the binary just download the latest release for your OS/Arch from [the release page](https://github.com/go-acme/lego/releases) and put the binary somewhere convenient.
lego does not assume anything about the location you run it from.

## From Docker

```bash
docker run goacme/lego -h
```

## From package managers

- [ArchLinux (AUR)](https://aur.archlinux.org/packages/lego):

```bash
yay -S lego
```

**Note**: only the package manager for Arch Linux is officially supported by the lego team.

- [FreeBSD (Ports)](https://www.freshports.org/security/lego):

```bash
cd /usr/ports/security/lego && make install clean
```

or

```bash
pkg install lego
```

## From sources

Requirements:

- `go` v1.15+
- environment variable: `GO111MODULE=on`

To install the latest development version from sources, just run:

```bash
go get -u github.com/casbin/lego/v4/cmd/lego
```

or

```bash
git clone git@github.com:go-acme/lego.git
make        # tests + doc + build
make build  # only build
```
