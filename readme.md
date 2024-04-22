## README: twclone

### intro:
This is a backend application for a microblogging platform in which the users can post and read `tweets`, users can follow and unfollow each other and users can get a `timeline` composed with `tweets` from users they follow.

### run and use it
You can either follow the steps in the dockerfile or build and run the image with docker.

#### if you don't want to use docker:
Make sure you have `go:1.22` and `sqlite` installed, then:
1. install the dependencies using
```
go mod download
```
2. run sqlite and execute the provided sql file
```
sqlite3 db.db < ./sql/setup_tables.sql
```
3. run the app
```
go run ./cmd/main.go
```
#### with docker
1. build the image
```
docker build -t TAG .
```
2. run the image
```
docker run -p 8080:8080 TAG
```
if you want to persist the storage attach a volume

### mock authentication
Before start using the app, let's explain the authentication logic that has been mocked. The app always authenticates users it does so by reading the `x-user-id` header which expects an uuid, if this header is absent the app will generate a user using a username provided via the `x-username` header, if neither is provided the app will 403.

### start tweeting

1. post a tweet:
```
curl --location 'localhost:8080/tweet' \
--header 'x-username: anon1' \
--header 'Content-Type: application/json' \
--data '{
    "text": "this is an example tweet"
}'
```
you will receive something like this:
```
{
    "message": "Tweet post was successful",
    "tweet": {
        "id": "cb108165-0f21-42b7-93f1-9ad66077eed4",
        "user_id": "c7afa779-396d-42e0-ad5a-2ad72f907238",
        "text": "this is an example tweet",
        "ts": 1713756612
    }
}
```
hold on to that `tweet.user_id`, that is the uuid generated for the username provided

2. Get the tweet:
```
curl --location 'localhost:8080/tweet/cb108165-0f21-42b7-93f1-9ad66077eed4'
```

3. Follow and Unfollow
```
curl --location --request POST 'localhost:8080/follow/USER_ID_TO_FOLLOW' \
--header 'x-user-id: YOUR_USER_ID'
```
```
curl --location --request POST 'localhost:8080/unfollow/USER_ID_TO_UNFOLLOW' \
--header 'x-user-id: YOUR_USER_ID'
```
4. Get the timeline: 
bare in mind only tweets from users you follow will appear
```
curl --location 'localhost:8080/timeline?limit=10&offset=0' \
--header 'x-user-id: YOUR_USER_ID'
```

