# gRPC-gateway REST API with Fiber

This project provides an automated gRPC-gateway REST API based on Fiber framework in Go.

## Prerequisites

- Go (at least Go 1.21)
- Fiber framework

## Installation

1. Clone the project repository:

```bash

git clone https://github.com/your-username/your-project.git

```

2. Rename the project to your desired name:

```bash

make rename

```

3. Add the gRPC service to your project:

```bash

make add-grpc-service

```

4. Generate the necessary code:

```bash

make generate

```

5. Customize the handler YAML file:

Edit the `handler.yaml` file to define your REST endpoints and their corresponding gRPC services.

6. Run the application:

```bash

make run

```

## Usage

Once the application is running, you can access the REST API endpoints defined in the `handler.yaml` file.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an
issue or submit a pull request.

## License

This project is licensed under the [MIT  License](LICENSE).
  

