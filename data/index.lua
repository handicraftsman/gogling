local g = require("gogling")
local net = require("gogling.net")
net.init("text/html")
data = [[
<!DOCTYPE html>

<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">

		<title>Gogling Test Page</title>
    <style>
      body {
        font-family: monospace
      }
      .data {
        margin: 0;
        padding: 10px;

        text-aling: center
      }
      ul {
        list-style-type: none
      }
      ul, li, p {
        margin-top: 0
      }
    </style>



	</head>
	<body class="">
		<header style="background-color: #003560; color: white;">
			<h3 class="data">Gogling ]] .. g.version .. [[</h3>
		</header>
    <article class="data" style="background-color: #EBEBEB">
      <p><kbd>
        Hello, world?<br/>
        You just started blank gogling app.<br/>
        This server supports 2 languages: classical Go template language and Lua<br/>
        Let me tell more about gogling's file structure:
        <ul>
          <li>
            <b>.git/</b><br/>
            <p>Nothing interesting. Just info about git-repo</p>
          </li>
          <li>
            <b>.gitignore</b><br/>
            <p>Just to prevent publishing binaries/object code</p>
          </li>
          <li>
            <b>data/</b><br/>
            <p>Here you can store all your PUBLIC data</p>
          </li>
          <li>
            <b>internal/</b><br/>
            <p>Here you can store all files you won't make public.
            For example, your lua libraries</p>
          </li>
          <li>
            <b>err/</b><br/>
            <p>Want custom 404 page? Not problem!</p>
          </li>
          <li>
            <b>src/</b><br/>
            <p>Gogling's source codes. Published under GPLv3</p>
          </li>
          <li>
            <b>conf.json</b><br/>
            <p>Here gogling stores it's config</p>
          </li>
          <li>
            <b>.travis.yml</b><br/>
            <p>Just to allow using Travis CI</p>
          </li>
          <li>
            <b>LICENSE</b><br/>
            <p>Just file with GPLv3 license</p>
          </li>
          <li>
            <b>Makefile</b><br/>
            <p>Is it really necessary to describe this file? You just built
            gogling via it</p>
          </li>
          <li>
            <b>README.md</b><br/>
            <p>README</p>
          </li>
        </ul>
	<a href="https://handicraftsman.tk.gogling/" alt="Gogling's Site" style="color: #003560">Gogling's Site></a>
      </kbd></p>
    </article>
	</body>
</html>
]]
net.echo(data)
