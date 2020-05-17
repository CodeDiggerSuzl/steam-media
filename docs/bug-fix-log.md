# Bug fix log of this project

- [x] 1. Can't pass all test in package: dbops

    **Solution**:
    1. Did not know that **`init`** method is a special method in golang.
    2. And because of my careless problem: didn't use `:=` and `=` in coding. Shame on me. :/
    3. Thanks to the warm-hearted buddy in a QQ group,which added me and help me to solve the bug. :)

## Date 5.4: A lot of wired shit just happened:

- [ ] 1. System stoped:

    After run `go build`, then system just stopped, not listening to the port, after debugging,that it seems that:
    `http.ListenAndServe(":9000", middleWareHandler)` doesn't work, then debug:
    ```go
    func (srv *Server) ListenAndServe() error {
    	if srv.shuttingDown() {
    		return ErrServerClosed
    	}
    	addr := srv.Addr
    	if addr == "" {
    		addr = ":http"
    	}
    	ln, err := net.Listen("tcp", addr)
    	if err != nil {
            // bug happens here
    		return err
    	}
    	return srv.Serve(ln)
    }
    ```
    So the address is in `dial.go` func `Dial` method,
    ```
    // For TCP, UDP and IP networks, if the host is empty or a literal
    // unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
    // TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
    // assumed.
    ```
    but it doesn't work.


- [ ] 2. Redirect to to new page after uploading video

    After uploading a new video, then the address and just jump to `videos/video-id/`, need to fix this shit.

- [ ] 3. The release connection is not working right.
    ```go
    //ReleaseConn Release Connection
    func (c *ConnLimiter) ReleaseConn() {
	    conn := <-c.bucket
	    log.Printf("Release connection: %d", conn)
    }
    ```
    here are the logs:
    ```log
    020/05/04 23:15:08 Did not reach the rate limitation
    2020/05/04 23:15:08 Did not reach the rate limitation
    2020/05/04 23:15:08 Release connection: -1
    2020/05/04 23:15:08 Release connection: -1
    2020/05/04 23:15:21 Did not reach the rate limitation
    2020/05/04 23:15:21 Release connection: -1
    2020/05/04 23:15:21 Did not reach the rate limitation
    2020/05/04 23:15:22 Did not reach the rate limitation
    ```
    TWICE !!!
- [ ] 4. The video id should be random.
- [ ] 5. The limiter is not working right, opened 4 tabs in chrome and is still did not get the `Too much requests` response, maybe is the http stuff working.
- [ ] Open a video in Edge is ok, and open the same page in Chrome just can not working,the request just hanging..., HTTP problem ?

    ![](https://tva1.sinaimg.cn/large/007S8ZIlly1geguxs8mpaj30yo08a75h.jpg)
- [ ] 6. Can upload all kinds of files, and will be treated as mp4.
Need to fix. Ps, the project is not that good.