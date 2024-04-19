# golang-web-lambda

## How to test

POST

```
curl "https://ahkr2sycw333enfo2gqwv3x5ia0bmvvf.lambda-url.ap-northeast-1.on.aws/post" -d '{"test": "Hello, world!"}' -H "Content-Type: application/json"
```

GET

```
curl "https://ahkr2sycw333enfo2gqwv3x5ia0bmvvf.lambda-url.ap-northeast-1.on.aws/reflect?test=a"
```