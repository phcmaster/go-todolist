Simple TODO list API.

## How to run

<hr>

####  To build the application locally:

1. Clone this project.

2. Access project folder: `cd go-todolist`

3. Download the dependencies: `go get github.com/gorilla/mux`

4. Run `go build main.go`

5. Execute the binary file `./main`

#### Running with docker:

1. Build the image: `docker build -t todo-list-api .`

2. Start the todo-list-api container: `docker run -d -p 8081:8081 todo-list-api`

3. Test it: `curl -X GET http://localhost:8081/todos`

##### The output should be:

`[{"id":"1","Item":"Put the trash outsite","Completed":false},{"id":"2","Item":"Buy water","Completed":true},{"id":"3","Item":"study go","Completed":true},{"id":"4","Item":"practice exercice","Completed":false}]`