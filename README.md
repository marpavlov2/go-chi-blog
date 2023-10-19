# go-chi-blog

## API ROUTES

GET ALL POSTS
```sh
curl "http://localhost:8080/posts" | jq
```

CREATE POST
```sh
curl -v -X POST -H "Content-Type: application/json" -d '{"title": "My First Post", "content": "My First Description"}' "http://localhost:8080/posts" | jq
```
