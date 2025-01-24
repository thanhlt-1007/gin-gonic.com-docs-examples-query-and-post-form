# gin-gonic.com-docs-examples-query-and-post-form

- Query and post form

- Reference: https://gin-gonic.com/docs/examples/query-and-post-form/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go get .

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

- POST /paging

```sh
curl --location 'localhost:8080/paging?page=1&per=20' \
--form 'message="MESSAGE"'
```
