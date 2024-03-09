# API Latency Checker

This is a golang tool to check latency of endpoints in the specified URL.

## Why is this tool important?
API latency refers to the duration it takes for an API to respond to a request. Having high API latency can affect your user experience, therefore it is key metric to evaluate the performance of your API. 

## How does this tool work?
This tool helps the developer check the URL provided and outputs the endpoints in a ```.txt``` file. Thereafter, it checks each URL in the file for its latency and prints it out in the console.


## How to run the code

```
go run main.go --url "URL" -output "FILENAME.txt"
```
## How to run the tests

```
go test -v
```

