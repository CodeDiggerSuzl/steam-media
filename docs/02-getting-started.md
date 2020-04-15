# Architecture of the stream-media website
Why stream-media website ?
- Go is a net-work programming language. Excellent at handling network IO or disk IO.
- Stream-media website contains much more tips at work.
- Top-notch native http lib and template engine(do not need other frame works.)

## Architecture

![](https://tva1.sinaimg.cn/large/007S8ZIlly1gdus1zi7pcj312q0qs7dy.jpg)


### Decoupling
Before the coding we need to know what is decoupling of front end and back end ?
- Render the front end pages and services by normal web engine.
- The back-end data is called post-processing and rendering through the rendered page script.

Advantages:
- Improve cooperation efficiency.
- The architecture is more flexible and easy to deploy, in line with the design of micro-services.
- High ability and reliability.

Disadvantages:
- Much more work.
- Team costs and learning costs are higher.
- Higher complexity.


# Back-end
## API design
- REST(Representational Status Transfer) API.
- REST is a design style not a standard.
- Now the RESTful API commonly use HTTP as its transfer protocol and JSON as data type.


RESTful API Features
- Uniform Interface.
- Stateless.

  No matter when or who call this api, it will return the same data.
- Cache-able

  Reduce the back-end pressure.

- Layered System
- C-S Architecture

API Design Principle
- Use URL(Uniform Source Location) style to design the API.
- USE different method( GET/POST/DELETE/PUT ) to differentiate access to resources(CRUD).
- The status code that returns should meet the requirements of HTTP resource description.


## Project API Design

![](https://tva1.sinaimg.cn/large/007S8ZIlly1gdusuwmzdbj30ue0oqgrk.jpg)

#### API Design: User
Source is user.

| Action              | URL              | Method | SC                   |
| :------------------ | :--------------- | :----- | :------------------- |
| Create User         | `/user`          | POST   | 201,400,500          |
| User Login          | `/user/username` | POST   | 200,400              |
| Get User Info       | `/user/username` | GET    | 200,400,401,403,500  |
| User Cancel Account | `/user/username` | DELETE | 204,400,401,403, 500 |

- 204:
- 401: don't verify user.
- 403: user's permission denied.
 
#### API Design: User Videos

| Action           | URL                             | Method | SC                  |
| :--------------- | :------------------------------ | :----- | :------------------ |
| List Videos      | `/user/:username/videos`        | GET    | 200,400,500         |
| Get One Video    | `/user/:username/videos/vid-id` | GET    | 200,400,500         |
| Delete One Video | `/user/:username/videos/vid-id` | DELETE | 204,400,401,403,500 |


#### API Design: User Comments

| Action             | URL                                 | Method | SC                  |
| :----------------- | :---------------------------------- | :----- | :------------------ |
| Show Comments      | `/video/vid-id/comments`            | GET    | 200,400,500         |
| Post One Comment   | `/video/vid-id/comments`            | POST   | 201,400,500         |
| Delete One Comment | `/video/vid-id/comments/comment-id` | DELETE | 204,400,401,403,500 |

### User API details
#### Project tree

router usage
- [ ] RESTful and RESTful alike api