/**
 * Simple server for local tests.
 * To use it:
 *  - install the module node-static: $ npm install node-static // only once
 *  - launch it:                      $ node local_server.js
 *
 */

var static = require('node-static');
var http = require('http');

// var file = new(static.Server)('./static');
var file = new(static.Server)();

http.createServer(function (req, res) {
    file.serve(req, res);
}).listen(8099);
