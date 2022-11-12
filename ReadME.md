# Golang excel Example

## How to run project

- docker-compose build --no-cache
- docker-compose up -d
- docker-compose ps
- docker-compose logs -f {service-name}
- send POST request to API' , http://127.0.0.1:8080/api/v1/files/upload , for upload excel file (.xlsx, .xlx formats)
- docker-compose down (remove and stop containers)

____

## Project Details
- visit http://localhost:9000/ (Portainer) , docker containers GUI
- visit http://localhost:8081/ (Mongo-Express), MongoDB GUI
- open golang container logs, see the logs during the run of the application