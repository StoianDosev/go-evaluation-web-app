# Instructions

## Prerequisites

Go - version 1.17 or Docker installed.

## Run

1. Open terminal in the **/EvaluationApp** folder.
2. Run:

```
 make run
```

or

```
go run ./src
```

or

```
docker run --rm -it  -p 5000:5000/tcp goevaluationappnethttp:latest
```

to run the web server.

3. Open browser on [http://localhost:5000](http://localhost:5000)

## Endpoints

- **GET:** [/ping](http://localhost:5000/ping)
- **GET:** [/errors](http://localhost:5000/errors)
- **POST:** [/evaluate](http://localhost:5000/evaluate)
- **POST:** [/validate](http://localhost:5000/validate)

## Unit tests

1. Open terminal in the **/EvaluationApp** folder
2. Run:

```
 make gotest
 ```

 or

 ```

 go test ./src/services...
 ```
 
  to run unit tests.

3. Run:

```

make gotest cover

```

or

```

go test -cover ./src/services...

```

to run test coverage.

## Usage

1.Simple arithmetics:

**/evaluate** endpoint
Accepts request body:
```
{"expression":"<simple math problem>"}
```

Response:

```
{"result":<the expression's result>}`
```

**/validate** endpoint
Accepts request body:
```
{"expression":"<simple math problem>"}
```

Response:

```
{
"valid":false,
"reason":"<the reason why the expression is invalid>"
}

or 

{"valid":true}`
```

Examples:

- Add two numbers together
  - What is 5 plus 13?
- Subtraction
  - What is 7 minus 5?
- Multiplication
  - What is 6 multiplied by 4?
- Division
  - What is 25 divided by 5?

- Set of operations, in sequence.
  - What is 3 plus 2 multiplied by 3?

**/errors** endpoint

Response:

```
[
    {
    "expression": "<a simple math problem>",
    "endpoint": "<endpoint URL>",
    "frequency": <number of times the expression failed>,
    "type": "<error type>"
    }
    ...
]
```