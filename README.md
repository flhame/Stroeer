# Stroeer

### REST API for handling user information and their comments
### Following Endpoints are supported:
#### Getting all users with their comments:

http://localhost:8081/api/v1/users-with-comments

#### Getting single user with its comments:
http://localhost:8081/api/v1/users-with-comments?userId=1

####
Unit tests are located in:
```internal/app/userservice/service_test.go```
####
Mocks are located in: ```internal/app/mock```
