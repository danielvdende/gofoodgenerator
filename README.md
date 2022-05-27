# GoFoodGenerator

The GoFoodGenerator is a small app written in Golang to generate weekly menu's. The idea is that coming up with what 
to eat in a week can be hard (and can be automated relatively easily). The GoFoodGenerator will generate a menu for a week
for you given a ruleset (which is hardcoded for now). These rules include things such as preparation time (where the idea
is that you don't want to prepare meals that take a long time to make during weekdays)

## Endpoints

### `GET /`

### `GET /food`

### `POST /food/:day`

## How to run
First build the docker image:
```bash
docker build -t gofoodgenerator .
```
then run the built image:
```bash
docker run -p 8080:8080 gofoodgenerator
```

Note: the commands above provide zero persistence (i.e. once the Docker container stops, any changes made to the database will be lost.
As we currently don't support updates to the database via the app, this is fine). If you do want persistence, you 


## Features still to come
- Setup ci/cd
- Allow for custom rulesets. This might be done by configuring it via the UI, or by passing a file.
- Support other rules than PreparationTime. This will require more flexible use of SQL generation. 
- Allow addition/modification/deletion of meals from the database.
