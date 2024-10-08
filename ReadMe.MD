# CSDD 1008 Week4 Assignment
## Build a Simple CRUD API with Go and Push to GitHub
### How to run API
![](images/Terminal.PNG)
Make sure GO is installed before running [Go Lang](https://go.dev/doc/install)

```
git clone https://github.com/DimanthaG/csdd1008_week5.git
cd csdd1008_week5
go run *.go
```
### GET /gamestats (Lists all game stats)
![](images/ListAll.PNG)
Lists all of the game stats that is currently available
### GET /gamestats/{id} (Lists specific id game stats)
![](images/ListId.PNG)
Endpoint gets a gamestat by its ID number
### POST /gamestats (Creates a new game stats)
![](images/Create.PNG)
Endpoint to create a new game stat 
### PUT /gamestats/{id} (Update game stats of specific id)
![](images/Update.PNG)
Update an already available game stat 
### DELETE /gamestats (Deletes Game Stat)
![](images/Delete.PNG)
![](images/CheckDeleted.PNG)
Deletes a game stat from the api
