1. Set up
```go mod vendor```

2. Running application
```CGO_ENABLED=1 go run main.go```

3. Example cURL
``` 
curl --location 'localhost:8000/ingredients' \
--header 'Content-Type: application/json' \
--data '{
    "name":"foobar",
    "cause_alergy": true,
    "type": 1,
    "status":1
}'
```

4. For the database i'm using SQLite to saving time.