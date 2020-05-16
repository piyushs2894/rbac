# RBAC

RBAC is Role Based Access Control System. Used to assign a role to user and remove a user from the role.
And also to check access for a user for a particular resource and action type.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

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
