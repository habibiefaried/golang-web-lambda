# golang-web-lambda

# How to test

## Unit test

```
make test
```

## Test the lambda
POST

```
curl "http://my-alb-1097420442.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"id": "1", "newurl": "https://google.com/a/b"}' -H "Content-Type: application/json"
```

```
curl "http://my-alb-1097420442.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"id": "1", "oldurl": "https://google.com/a/b", "newurl": "https://randomzied.com/a/b"}' -H "Content-Type: application/json"
```

```
curl "http://my-alb-1097420442.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"id": "1", "oldurl": "https://randomzied.com"}' -H "Content-Type: application/json"
```