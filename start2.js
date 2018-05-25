var fs = require('fs');

fs.appendFile('mynewfile1.txt', 'Hello content!', function (err) {
  if (err) throw err;
  console.log('Saved!');
});
fs.appendFile('mynewfile1.txt', 'Here goes the second line!', function (err) {
  if (err) throw err;
  console.log('Saved!');
});

var http = require('http');
var fs = require('fs');
http.createServer(function (req, res) {
  fs.readFile('mynewfile1.txt', function(err, data) {
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.write(data);
    res.end();
  });
}).listen(8080);