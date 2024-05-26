# golang-web-lambda

# How to test

## Unit test

```
make test
```

## Test the lambda
POST

```
curl "http://my-alb-1404328830.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"id": "1", "newurl": "https://google.com/a/b"}' -H "Content-Type: application/json"
```

```
curl "http://my-alb-1404328830.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"id": "1", "newurl": "https://1.2.3.4/web.aspx"}' -H "Content-Type: application/json"
```

```
curl "http://my-alb-1404328830.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"id": "1", "oldurl": "https://google.com/a/b", "newurl": "https://randomzied.com/a/b"}' -H "Content-Type: application/json"
```

```
curl "http://my-alb-1404328830.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"id": "1", "oldurl": "https://randomzied.com"}' -H "Content-Type: application/json"
```

TODO
1. Missed testcase, adding 3 new rule in the same ID resulting rejection. it should be able to receive all of them