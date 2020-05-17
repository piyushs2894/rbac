# RBAC

RBAC is Role Based Access Control System. Used to assign a role to user and remove a user from the role.
And also to check access for a user for a particular resource and action type.

### Background

In rbac project, there are 6 main entities listed below, which are defined in `src/model` package. This basically contains methods which are communicating with DB or fileSystem.
1) `action` representing actionType
2) `role` defining user role like Software Developer, Quality Assurance etc
3) `resources` representing service or resource like DB, server for different services
4) `user` for handling user signup and login flow
5) `userRoleMapping` defining userId and roleId mapping. It can be many to many mapping. Here userId and roleId combination will be unique and is validated on its insertion.
6) `roleResourceMapping` which represents for a particular role and actionType which resource it has access to. Here combination of roleId, resourceId and ActionId will be unique.

And there is `src/manager` package, which contains rbacImplementation for different functionalities required by the system. It represents business logic layer, which calls methods of different modules as per functionality required.

Also in this project, I have defined multiple modules like `userRoleMappingModule`, `roleResourceMappingModule`, `actionModule`, `resourceModule` which are initialized on app start in this call `managerModule := manager.New()`. Defining these modules has great advantage, because it makes our system flexible as in these modules, we can also initialize DB config, Cache config etc. Currently I have not used any DB configs, I have used maps in place for fetching values corresponding to different entities similar to DB. And as it can change for different services, or for unit testing, mock modules are required inside a project, then we just need to initialize those in app start and pass as a parameter. There we don't need to change code in entities or business logic layer. For more information, this video can be referred https://www.youtube.com/watch?v=o_TH-Y78tt4.

### Prerequisites

This project is written in golang. So for this golang should be installed. Here 1.14.2 version of golang is used.

### Installing

For installing golang, please refer this link - https://www.geeksforgeeks.org/how-to-install-golang-on-macos/ . 

1) Install go 1.14.2 from this https://golang.org/dl/ .
2) Check if go is successfully installed by checking version
```
go version
```
3) Create a folder named `go` inside this `/Users/<userName>` directory.

4) Set GOPATH using this command. Replace <userName> with your user name and run below two commands.

```
cd ~
echo "export GOPATH=/Users/<userName>/go" >> .bash_profile
```
5) Check that your .bash_profile conatins the following path using the following command:
```
cat .bash_profile
```
6) Check if GOPATH is set using following command:
```
echo $GOPATH
```
7) In `go` folder, create folder `src` and inside `src`, create another folder named `github.com`

8) Move RBAC project inside `github.com` folder.

8) Go inside `RBAC` directory. And run `main.go` by using following command:
```
go run main.go
```

### Assumptions

1) Input from /etc/files on app start is assumed to be valid.
2) ActionType defined as Action are independent and unique. No parent-child relationship should exist.

### Testing

1) For Assign a role to user, use this cURL request similar to below: 
```
curl --location --request POST 'http://localhost:8080/api/v1/addOrUpdateRoletoUser' \
--header 'Content-Type: application/json' \
--data-raw '{
	"user_id": 3,
	"role_id": 2
}'
```
It will give `Add role to User is Success` on success otherwise error on failure.

2) For Removing a role from user, use this cURL request similar to below: 
```
curl --location --request POST 'http://localhost:8080/api/v1/removeRolefromUser' \
--header 'Content-Type: application/json' \
--data-raw '{
	"user_id": 3,
	"role_id": 2
}'
```
It will give `Remove role from User is Success` on success otherwise error on failure.

3) For Checking user access for particular resource and action, use this cURL: 
```
curl --location --request GET 'http://localhost:8080/api/v1/checkUserAccess?user_id=3&action_type=Delete&resource_name=Fulfilment%20Service'
```
It will give `User is Authorized for accessing this resource` if user has access otherwise `Unauthorized Access for this resource` if user is not allowed to access.
