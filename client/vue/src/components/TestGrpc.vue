<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <button @click="helloWorld">HelloButton</button>
  </div>
</template>

<script>
// eslint-disable-next-line 
const {HelloRequest, HelloReply} = require('../assets/protobuf/helloworld_pb.js');
const {GreeterClient} = require('../assets/protobuf/helloworld_grpc_web_pb.js');

// for grpc-web debug
const enableDevTools = window.__GRPCWEB_DEVTOOLS__ || (() => {});
const client = new GreeterClient('http://localhost:8081');
enableDevTools([
    client,
]);

const request = new HelloRequest();
request.setName('Worldaaaa');

export default {
  name: 'HelloWorld',
  data: function() {
    return {
      msg: ""      
    }
  },
  methods: {
    helloWorld: function(){
      // eslint-disable-next-line
      const request = new HelloRequest();
      request.setName('grpc-web with vue');
      client.sayHello(request, {}, (err, response) => {
        //console.log(response.getMessage());
        this.msg = response.getMessage();
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
