## Compile & Run Server
run the command below to compile the listener into a binary:

go build ./cmd/server/
Next, run the server as:

./server
The above command should run the TCP server and we can make connections to it.

## Testing the TCP Server
To test the TCP server, you can use redis-cli using the command:

```bash
redis-cli -p 3333
```

Try `hi` command, the server should respond with a Hi back

```bash
"Hi back!""
```

also you can use `testurl` command to check if url is reachable:

```bash
redis-cli -p 3333
127.0.0.1:3333> testurl threefold.io
(integer) 1
127.0.0.1:3333> testurl gle.co
(integer) 0
```

