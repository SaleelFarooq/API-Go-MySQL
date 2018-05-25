var fs = require('fs');

fs.writeFile('mynewfile3.txt', 'This is my text', function (err) {
  if (err) throw err;
  console.log('Replaced!');
});



fs.writeFile('mynewfile3.txt', 'This is my new text,yes,I have changed it', function (err) {
  if (err) throw err;
  console.log('Replaced!');
});

var http = require('http');
var fs = require('fs');
http.createServer(function (req, res) {
  fs.readFile('mynewfile3.txt', function(err, data) {
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.write(data);
    res.end();
  });
}).listen(8080);
