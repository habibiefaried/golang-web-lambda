# golang-web-lambda

# How to test

## Unit test

```
make test
```

## Test the lambda
POST

```
curl "http://my-alb-1097420442.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"oldurl":{}, "newurl": {"id": "1", "url": "https://google.com/a/b"}}' -H "Content-Type: application/json"
```

```
curl "http://my-alb-1097420442.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"oldurl":{"id": "1", "url": "https://google.com/a/b"}, "newurl": {"id": "1", "url": "https://randomzied.com/a/b"}}' -H "Content-Type: application/json"
```

```
curl "http://my-alb-1097420442.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"oldurl":{"id": "1", "url": "https://randomzied.com/a/b"}, "newurl": {}}' -H "Content-Type: application/json"
```