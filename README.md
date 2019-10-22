## Installation
- Just unzip the package in your go workspace where the Gopath is set. 

## Modelisation

![Alt text](source/modelisation_achievement_system.png?raw=true "Modelisation")
- Each players have a game resume which contains specific informations about every games played and for each games, specfic statistics are linked to each game.

- The game struct can be linked to the global game struct (not created for the test) which contains every informations about each games for each players. The global game struct can be represent as a table in a datasource (lika any others structures created).

## Testing
- ```go test -v``` for launching all the tests, compose of the driver and unit test for the achievements logic functions
  
## Improvements
- Use Goroutines for launching achievement systems for each player after a game.
- Use datasource for save all the data (like SQLite or PostgreSQL)
- Adding error gestion if specific errors are detected