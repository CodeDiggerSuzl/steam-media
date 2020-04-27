# Streaming module

- Static video,not RTMP(Real-Time Messaging Protocol),live stream

- Independent service, can be deploy alone.
- Unified standard api
- Streaming server
  - Upload files
  - Streaming

- This module will not have ops to db.


Both the streaming and upload files will need to keep a long connection.

## Step by step

1. Write main func, register handlers
2. Write handler code, tow handlers
3. Write response.go, define error response
4. When you app is online, there might be some user or hacker constantly send request to your server,Resulting in insufficient connections.even run all of your RAM and then you system will crash. Therefore we need the Read-Limit. Often use the token bucket algorithm.

    Bucket algorithm: TODO

    Can not use a array other type of container to store the tokens, cause each handler is a goroutine, maybe we need to add locks to the containers, Which will reduced performance of the system.

    The best way to handle this case is to use channel, which is encourage by Go.

5. Use channel to solve the bucket token.
   1. Constructor
   2. Get Conn method
   3. Release token

6. Use middle ware to add the bucket token.(just like what we do in API-module)
7. Finish handlers.
   1. Stream handler:`http.ServeContent()`
   2. uploadHandle
      1. check handler


8. Check upload handler.