# CSDD 1008 Week4 Assignment
## Build a Simple CRUD API with Go and Push to GitHub
### How to run API
<img width="1413" alt="home-page" src="https://github.com/DimanthaG/csdd1008_Go_API/blob/master/images/Terminal.png?raw=true">

Make sure GO is installed before running [Go Lang](https://go.dev/doc/install)

```
git clone https://github.com/DimanthaG/csdd1008_Go_API.git
cd csdd1008_week5
go run *.go
```
### GET /gamestats (Lists all game stats)
<img width="1413" alt="home-page" src="https://github.com/DimanthaG/csdd1008_Go_API/blob/master/images/ListAll.png?raw=true">
Lists all of the game stats that is currently available

### GET /gamestats/{id} (Lists specific id game stats)
<img width="1413" alt="home-page" src="https://github.com/DimanthaG/csdd1008_Go_API/blob/master/images/ListId.png?raw=true">
Endpoint gets a gamestat by its ID number

### POST /gamestats (Creates a new game stats)
<img width="1413" alt="home-page" src="https://github.com/DimanthaG/csdd1008_Go_API/blob/master/images/Create.png?raw=true">
Endpoint to create a new game stat 

### PUT /gamestats/{id} (Update game stats of specific id)
<img width="1413" alt="home-page" src="https://github.com/DimanthaG/csdd1008_Go_API/blob/master/images/Update.png?raw=true">
Update an already available game stat 

### DELETE /gamestats (Deletes Game Stat)
<img width="1413" alt="home-page" src="https://github.com/DimanthaG/csdd1008_Go_API/blob/master/images/Delete.png?raw=true">
<img width="1413" alt="home-page" src="https://github.com/DimanthaG/csdd1008_Go_API/blob/master/images/CheckDeleted.png?raw=true">
Deletes a game stat from the
