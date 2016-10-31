![Master](https://img.shields.io/travis/handicraftsman/gogling/master.svg?label=Master)
![Development](https://img.shields.io/travis/handicraftsman/gogling/development.svg?label=Development)
<br/><a href="https://travis-ci.org/handicraftsman/gogling">Check Travis</a>

# Gogling - installation
## Step 0: install Go
You can skip this step if you already have Go<br/>
Otherwise, install it via your package manager or get it from https://golang.org/
## Step 1: clone repo
`git clone https://github.com/handicraftsman/gogling`
## Step 2: pull deps & sources
`make pull`
## Step 3: build Gogling
`make build`
## Step 4: run Gogling
`./gogling`

<br/>
## For developers:
You can build & run Gogling with `make br` command

# Stable Version
I'll publish first "stable" version after:

1. Packing it into go-styled package
2. Adding "Show Error Traceback" feature for pages like "500" (lua only)
3. Adding HTTPS
