# NU-Task

## Requirements 
Go v1.21

## Build 
```
go build -o .\bin\ .\cmd\counter .\cmd\counter-server
```

## Running

### CLI

```
.\bin\counter.exe --word a --case=true --whole=true
```

### Host

```
 .\bin\counter-server.exe    
```

Sample request
```
curl -X POST -H "Content-Type: application/json" -d '{"directory":"corpus", "word":"john"}' http://localhost:8080/counter
```

Sample request on Windows:
```
curl.exe -X POST -H "Content-Type: application/json" -d '{\"directory\":\"internal\\counter\", \"word\":\"a\", \"case\":false, \"whole\":false}' http://localhost:8080/counter
```

## Output

Example 
```
$ .\bin\counter.exe --dir ./cmd/ --word func --case=true --whole=true  
2023/10/16 06:02:08 The word "func" appears 2 times in file cmd\counter-server\main.go
2023/10/16 06:02:08 The word "func" appears 1 times in file cmd\counter\main.go
2023/10/16 06:02:08 The word "func" appears 3 times in total under the directory "./cmd/"
3
```

## Tests
Use tests runner.