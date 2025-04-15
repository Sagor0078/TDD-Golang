# Notes
Acceptance tests should validate the system from an external perspective, so ideally you'd run these tests against a real instance of your app or via a test HTTP server.

For bigger systems (e.g., microservices, CLI apps), acceptance tests may include:

    Spinning up Docker containers

    Using test fixtures or seed data

    Calling external APIs or database mock servers