![Master](https://img.shields.io/travis/handicraftsman/gogling/master.svg?label=Master)
![Development](https://img.shields.io/travis/handicraftsman/gogling/development.svg?label=Development)
<br/><a href="https://travis-ci.org/handicraftsman/gogling">Check Travis</a>

# Gogling

This is a complete rewrite of gogling written in 2016.

## Example

```lua
-- You can import go packages if you built them
local fmt = gogling.U.import('fmt')

-- You can call methods from imported packages
fmt.Println('Hello, World!')

-- github.com/gorilla/mux Router instance can be accessed as gogling.I.Router
gogling.I.Router:HandleFunc('/', gogling.U.wrap(function(session)
  fmt.Fprintf(session.writer, 'Hello, World!')
end))
```

## Building standard library

You can build Golang's standard library into a set of plugins which can be
loaded as lua modules using `gogling.U.import` function.

To do this you'll need ruby and bash.

First, you need to generate go-to-lua bindings by running `get-gostdlib.rb`
(was tested on ruby 2.5.0, but will probably work on earlier versions).

Then you need to build generated bindings by running `build-gostdlib.sh`.

You'll find built standard library in the `lib/` directory.

Note that not all go libraries are compiled for now because some of them cause compilation
errors when trying to relay them to the lua world. One of them is `math`.