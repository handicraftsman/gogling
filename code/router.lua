local fmt = gogling.U.import('fmt')
local time = gogling.U.import('time')

fmt.Println('Hello, World!')

gogling.I.Router:HandleFunc('/', gogling.U.wrap(function(session)
  fmt.Fprintln(session.writer, 'Hello, world! Current time is ' .. time.Unix(1392899576, 0):Format(time.RFC3339))
end))

gogling.I.Router:HandleFunc('/panic', gogling.U.wrap(function(session)
  gogling.I.Logger.Panic('Panicking')
end))

gogling.I.Router:HandleFunc('/reload', gogling.U.wrap(function(session)
  gogling.U.reload(session, "Done! Now all new requests will go to a new lua instance")
end))