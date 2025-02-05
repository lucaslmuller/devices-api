# Running the service

To run the service, execute the command bellow:

```
make run
```

Nice! Now the server is running on port `:8080`

To stop the execution you can run:

```
make stop
```


# Documentation

To see the documentation, import the `Test.postman_collection.json` file in your Postman.


# Testing

In order to run the tests, there are some options:

## 1. Unit Tests Only

```
make unit_test
```

## 2. E2E Tests Only

```
make e2e_test
```

## 3. Unit and E2E Tests

```
make test
```