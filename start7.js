var events = require('events');
var eventEmitter = new events.EventEmitter();


//Create an event handler:
var handler1 = function () {
  console.log('I hear a scream!');
}

var handler2 = function () {
	console.log('Where is it ?');
}
//Assign the event handler to an event:
eventEmitter.on('scream', handler1);
eventEmitter.on('s', handler2);

//Fire the 'scream' event:
eventEmitter.emit('scream');
eventEmitter.emit('s');