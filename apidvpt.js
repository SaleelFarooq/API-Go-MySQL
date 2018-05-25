const express = require('express');
const bodyParser = require('body-parser');
const app = express();
app.use(bodyParser.urlencoded({ extended: false }));

app.get('/', (request, response) =>  response.sendFile(`${__dirname}/index.html`));

app.post('/login', (request, response) => {
  const username = request.body.username;
  const password = request.body.password;
  response.writeHead(200,{'Content-Type' : 'text/html'});
  var text="username =" + username + "Password =" + password + "for log in";
  response.write(text);
  response.end();
  
});

app.post('/signup', (request, response) => {
  const username = request.body.username;
  const password = request.body.password;
  response.writeHead(200,{'Content-Type' : 'text/html'});
  var text="username =" + username + "Password =" + password + "for sign up";
  response.write(text);
  response.end();
  
});

app.listen(3000,()=>console.info('Running on 3000'));