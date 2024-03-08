<h1>Receipt Processor</h1>

To run this service locally, run the following command

`go run main.go app.go`

To run the tests for this application, run

`go test`

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