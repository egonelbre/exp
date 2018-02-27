# Automatic GopherJS building and reloading server

To auto-rebuild client code:

```
cd client
watchrun gopherjs build --output ../server/client/app.js
```

To auto-rebuild and restart server:

```
cd server
watchrun -care "*.go" go build . == ./server
```
