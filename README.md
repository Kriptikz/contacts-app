### Go Web Server Template

This project provides a robust starting point for building web applications using Go. It is pre-configured with a Dockerfile and environment variables, simplifying the setup process. The project includes basic API endpoints that serve HTML.

The repository is ready to be cloned and deployed on most hosting platforms.


## Key Features

- Docker integration for easy containerization and deployment.
- Pre-configured environment variables for quick setup.
- Basic API endpoints serving HTML, providing a starting point for your application.
- A `/ping` route for health checks of a running container.

## Getting Started

1. Clone this repository
```bash
git clone https://github.com/kriptikz/go-web-server-template
```

2. Create your .env file. You can use the .env.example as a reference. 
```bash
cp .env.example .env
```

3. From here you can run the project directly via the Makefile command:
```bash
make start-dev
```
>Make sure you have the correct version of go installed according to the go.mod file.
>You can use this command to check your go version:
>```bash
>go version
>```

4. Or you can build the Docker image: 
```bash
make docker-build
```

5. Run the Docker container: 
```bash
make docker-start
```

> make sure you are using the same port from your .env

## Adding More Routes

You can add more routes to the RoutesHandler function in the api/handlers.go file. You can also just swap out the RoutesHandler with another package like Gorilla/Mux, Chi, etc.


## Hosting

This project is ready to be deployed on any platform that supports Docker, such as Render, Heroku, AWS, GCP, Azure, etc.

## Contributing

Feel free to fork this project and make any changes you like. If you think your changes would be beneficial to others, please open a pull request.

## License

This project is licensed under the terms of the MIT license.
