# goclean-architecture
Sample code to demonstrate how to implement clean architecture in golang

# Install
Clone project to go folder tree
I take `go-workspace` as an example name for my go workspace
```bash
cd ~/go-workspace/src
git clone git@github.com:MrBigKuma/goclean-architecture.git goclean-architechture
```

Setup gopath, get library & test
```bash
cd ~/go-workspace

export GOPATH=`pwd`
go get ./...
go test goclean-architecture
```

# Source Overview
This source has 4 packages:

1. `entity`: only store core data model.
2. `usecase`: store business logic
3. `interfaceacdapter`: store controller & database service (repo).
4. `infrastructure`: has helpers, utils to interact with outside.

## entity
This package should has list of struct that is used for usecase.
These struct can be passed to `interfaceadapter` layer.

## usecase
1. Has interfaces for of outside layer, like: database service (repo), mail.
2. Implement business logic for application.

## interfaceadapter
Currently has 2 main subpackages:

**controller:** for receive request from outside. This one should handle things like request data validation, create json response...
These function will call usecase's function to execute logic and represent the usecase's response to API request. 

**repository:** database services. It should has its own data model with some extra field to be stored to database.
E.g. `CreatedTime`, `UpdatedTime` should be presented in every record in database but it's not required in entity.

## infrastructure
There should be very few code here like common database service, third-party lib helpers...