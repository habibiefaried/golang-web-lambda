# golang-web-lambda

# How to test

## Unit test

```
make test
```

## Test the lambda
POST

```
curl "https://ahkr2sycw333enfo2gqwv3x5ia0bmvvf.lambda-url.ap-northeast-1.on.aws/whitelist" -d '{"domain": "itb.ac.id"}' -H "Content-Type: application/json"
```

GET

```
curl "https://ahkr2sycw333enfo2gqwv3x5ia0bmvvf.lambda-url.ap-northeast-1.on.aws/is-whitelisted?domain=itb.ac.id"
```