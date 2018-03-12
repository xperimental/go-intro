# go-intro

This repository contains slides and examples used for my Go introductory workshop.

## Usage

### View online

The slides can be viewed online using a public server:

- [Part 1](https://go-talks.appspot.com/github.com/xperimental/go-intro/go-intro-part1.slide)
- [Part 2](https://go-talks.appspot.com/github.com/xperimental/go-intro/go-intro-part2.slide)

### Local server

If you have a working Go installation you can use `go get` to download this repo to your machine for playing with the examples:

```bash
go get -d github.com/xperimental/go-intro
```

The slides are in a format understood by [present](https://godoc.org/golang.org/x/tools/cmd/present). To view the slides locally use the following (after downloading the repo):

```bash
# Get present tool
go get -u golang.org/x/tools/cmd/present
# With $GOPATH/bin in $PATH
# Inside go-intro repo
present
# Then open browser at http://localhost:3999
# and navigate to slides.
```

## Talks

- [v1](https://github.com/xperimental/go-intro/tree/v1) held at [FFBW::Camp 2017](https://ffbsee.de/doku.php?id=events:camp)
- Workshop at [HolidayCheck AG](https://www.holidaycheck.de): [v2](https://github.com/xperimental/go-intro/tree/v2) first part; [v3](https://github.com/xperimental/go-intro/tree/v3) second part
