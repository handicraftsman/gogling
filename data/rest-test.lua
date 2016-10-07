-- Will be deleted after finishing developing REST module
local net = require("gogling.net")
local rest = require("gogling.net.rest")
net.init("text/html")
data = rest.getAll()["foo"]


net.echo(data)
