# grpc-gateway-test

Demo project for testing the gRPC gateway and OpenAPI spec generation.

## Running

Prerequisites: Go 1.17

Build the project by executing:
```bash
./build.sh -b
```

Run the server by executing:

```bash
./build.sh -r
```

This will start gRPC on port `8989` and the Gateway + Swagger UI on port `8990`.

To call the gRPC server directly you can use the included client:

```bash
./bin/client
```

To view the Swagger UI, point your browser to http://localhost:8990/static/swagger-ui/ 
You can test the API right from the UI.

## Changing

Generated gRPC and protobuf code is in the `gen` folder. This is generated from the definitions in the `proto` folder.

If you need to re-generate these files, you can do the following:

```bash
./build.sh -i       # to install the tools
./build.sh -g       # to run the generator
```

Note: all required tools for code generation are vendored in the `tools` folder, and will be installed in the `tools/bin` folder.
