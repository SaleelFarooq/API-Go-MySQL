var http=require('http')
var dt=require('./myfirstmodule')

http.createServer(function(req,res){
	res.writeHead(404,{'Content-Type':'text/html'})
	res.write("The URL of request is :: "+req.url)
	res.write("Time :: "+ dt.myDateTime())
	res.end()
}).listen(8080);