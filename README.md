# golang-web-lambda

# How to test

## Unit test

```
make test
```

## Test the lambda
POST

```
curl "http://my-alb-323469957.ap-northeast-1.elb.amazonaws.com/whitelist" -d '{"domain": "itb.ac.id"}' -H "Content-Type: application/json"
```

GET

```
curl "http://my-alb-323469957.ap-northeast-1.elb.amazonaws.com/is-whitelisted?domain=itb.ac.id&port=443"
```