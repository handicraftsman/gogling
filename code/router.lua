local fmt = gogling.U.import('fmt')

fmt.Println('Hello, World!')

gogling.I.Router:HandleFunc('/', gogling.U.wrap(function(session)
  fmt.Fprintf(session.writer, 'Hello, World!')
end))