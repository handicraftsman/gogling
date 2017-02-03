local g = require("gogling")
local net = require("gogling.net")
local cookie = require("gogling.net.cookie")
net.init("text/html")

net.echo("Gogling v"..g.version.." - ONLINE")
