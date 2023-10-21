# go-chi-blog

## API ROUTES

Get all posts
```sh
curl "http://localhost:8080/posts" | jq
```

Create post
```sh
curl -v -X POST -H "Content-Type: application/json" -d '{"title": "My First Post", "content": "My First Description"}' "http://localhost:8080/posts" | jq
```
