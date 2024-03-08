<h1>Receipt Processor</h1>

To run this service locally, run the following command, default port is `8080`

```go run main.go app.go```

To run the tests for this application, run

```go test```

<h3> Endpoints </h3>

To hit the service's endpoints after it is running, you can use

POST `localhost:8080/receipts/process` with the following body

```
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}

```

GET `localhost:8080/receipts/{id}/points` with the ID provided from the above call.


<h3>Considerations</h3>

Here are some considerations I had when building this app

1. Separate the calculation functions for easier testing, I used the included tests for TDD to build this.
2. Add constants for the point values so future scoring changes could be made.
3. Store the receipt point score in the DB (in-memory map) so future requests for that receipt are quicker and use less resources.

<h3>Thoughts for Future Development</h3>

1. Add E2E tests.
2. Re-evaluate application logic in REST handlers and move code to application layer as needed.
3. Create DB layer as needed.
4. Validate input using https://github.com/go-playground/validator
5. Dockerize the application.