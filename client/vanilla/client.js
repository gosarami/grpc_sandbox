const {HelloRequest, HelloReply} = require('./helloworld_pb.js');
const {GreeterClient} = require('./helloworld_grpc_web_pb.js');

// for grpc-web debug
const enableDevTools = window.__GRPCWEB_DEVTOOLS__ || (() => {});
const client = new GreeterClient('http://localhost:8081');
enableDevTools([
    client,
]);

const request = new HelloRequest();
request.setName('Worldaaaa');

client.sayHello(request, {}, (err, response) => {
  console.log(response.getMessage());
});