
Simple REST API Go microservice built as hexagonal architecture (echo framework)
For frontend I use Angular framework. This application is for learning purpose.

Backend generate hash code for each url what comes from the frontend.
That Url-hash pair can be listed, deleted

Hexagonal Architecture tackles this issue by building the application around the core. The main objective is to create fully testable systems that can be driven equally by users, programs and batch scripts in isolation of database.
The core alone is not very useful. Something has to drive this application, call the business logic methods. It may be a HTTP request, automatic test or integration API. These interfaces that drive the application we call the primary ports and the modules that use them are primary adapters.


Install:
docker-compose up
