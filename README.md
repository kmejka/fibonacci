# fibonacci
Fibonacci counter REST service, implemented in golang
________

## Usage
### Installation and start
```sh
$> go install
$> fibonacci
```
The service by default starts on port '8080'.

### Calls
The service right now has two endpoints:

The get the typical 'Hello world' response call '/':
```sh
$> curl localhost:8080/
   Welcome to fib service, call /fibonacci/:n
```

To get the first 'n' values of the fibonacci sequence call '/fibonacci/:n':
```sh
$> curl localhost:8080/fibonacci/10
   1, 1, 2, 3, 5, 8, 13, 21, 34, 55
```

## Technology
The service is written in golang, with [gin](https://github.com/gin-gonic/gin) serving as the web framework and
[go-check]("https://github.com/go-check/check") in tests.

The `fibservice.go` is the place where go counting takes place. The calculations occur in a seperate go routine.
Calculation is done with the help of a channel from which the consecutive values are read.