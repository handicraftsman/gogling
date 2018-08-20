gogling.I.Router:HandleFunc("/", gogling.U.wrap(function(session)
  session.writer:Write("Hello, World!")
end))