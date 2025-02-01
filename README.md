# Running the service

To run the service, execute the command bellow:

```
docker compose down -v && docker compose build --no-cache && docker compose up
```

Nice! Now the server is running on port `:8080`


# Documentation

To see the documentation, import the `Test.postman_collection.json` file in your Postman.


# Testing

In order to run the tests, execute the command bellow:

```
make test
```