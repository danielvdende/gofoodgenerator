# Go Food Generator

## TODO
- setup repo
- setup ci/cd
- finalize endpoint for regeneration

## Endpoints

### `/food`


## How to run
First build the docker image:
```bash
docker build -t gofoodgenerator .
```
then run the built image:
```bash
docker run -p 8080:8080 gofoodgenerator
```