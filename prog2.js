var http = require('http');
var dt = require('./myfirstmodule')


http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.write("The Date and Time are currently :: "+ dt.myDateTime())
    res.end('See you!');
}).listen(8080);