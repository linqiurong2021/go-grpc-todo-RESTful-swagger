var pb = require('../proto/v1/todo-service_pb');

var p = new pb.Person();
p.setName("Tom");
p.setAge(18);

var bytes = p.serializeBinary();

var unBytes = pb.Person.deserializeBinary(bytes);

var name = p.getName();
var age = p.getAge();

console.log(bytes);
console.log(unBytes);
console.log("==================================");
console.log("my name is ", name, "age is", age);