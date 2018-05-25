var http = require('http');
var toupper=require('upper-case')

http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'text/html'});
    text="Hello world" + "\n" + toupper('Hello World')
    res.write(text)
    res.end();
}).listen(8080);