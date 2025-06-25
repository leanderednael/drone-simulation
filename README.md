# Drone Simulation

## Scenario

There are two automatic drones that fly around London and report on traffic conditions. When a drone flies over a tube station, it assesses what the traffic condition is like in the area, and reports on it.

A simulation has one dispatcher and two drones. Each drone should "move" independently on different processes. The dispatcher should send the coordinates to each drone detailing where the drone's next position should be. The dispatcher should also be responsible for terminating the program.

When the drone receives a new coordinate, it moves, checks if there is a tube station in the area, and if so, reports on the traffic conditions there.

- The simulation should finish at 08:10, when the drones will receive a "SHUTDOWN" signal.
- The two drones have IDs `6043` and `5937`. There is a file containing their lat/lon points for their routes. The csv file layout is `drone-id,latitude,longitude,time`
- There is also a file with the lat/lon points for London tube stations: `station,lat,lon`
- Traffic reports should have the following format:
  - Drone ID
  - Time
  - Speed
  - Conditions of Traffic (`HEAVY`, `LIGHT`, `MODERATE`). This can be chosen randomly.

## Remarks

1. Assume that the drones follow a straight line between each point, travelling at constant speed.
2. Disregard the fact that the start time is not in synch. The dispatcher can start pumping data as soon as it has read the files.
3. A nearby station should be less than 350 meters from the drone's position.
4. Bonus point: Put a constraint on each drone to have limited memory, so they can only consume ten points at a time.

## Requirements

[Docker](https://docs.docker.com/get-docker/), and / or [Go 1.24](https://golang.org/doc/install)

## Instructions

### To run the simulation

- `go run main.go`, or
- `go build -o simulation && ./simulation`

- or in a Docker container:
  - `docker build -q -t simulation .`
  - `docker run simulation`

### To run the tests

- `go test ./...`
- with detailed coverage: `go test -coverprofile cov ./... && go tool cover -html=cov && rm cov`

- or in a Docker container:
  - `docker build -q -t simulation .`
  - `docker run -e cmd=test simulation`
