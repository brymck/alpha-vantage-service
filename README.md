alpha-vantage-service
=====================

![Build](https://github.com/brymck/alpha-vantage-service/workflows/Build/badge.svg)
[![codecov](https://codecov.io/gh/brymck/alpha-vantage-service/branch/master/graph/badge.svg)](https://codecov.io/gh/brymck/alpha-vantage-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/brymck/alpha-vantage-service)](https://goreportcard.com/report/github.com/brymck/alpha-vantage-service)

A service utilizing the Alpha Vantage API.

Features
--------

* [CODEOWNERS](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/about-code-owners) for
  defining roles and responsibilities
* Hermetic Docker builds
* An extremely small Docker container (< 4 MB) from some optimizations in how the Go binary is created

Usage
-----

[Get an API key from Alpha Vantage](https://www.alphavantage.co/support/#api-key) and export it as an environment
variable named `$ALPHA_VANTAGE_API_KEY`. Then run either the commands below for a standard or containerized build:

```zsh
# Build a standard binary and run it...
make run

# ...or build a container and run it
make docker
docker run -it --rm -p 50051:50051 -e ALPHA_VANTAGE_API_KEY=$ALPHA_VANTAGE_API_KEY docker.pkg.github.com/brymck/alpha-vantage-service/alpha-vantage-service
```

If you have [gRPCurl](https://github.com/fullstorydev/grpcurl) installed, you can test the service with the scripts in
`bin/`:

```zsh
bin/ts HYG
bin/quote IVV
```

Once deployed, you can

```zsh
make client
ADDR=$(gcloud run services describe alpha-vantage-service --platform managed --format 'value(status.url.scope())'):443

./client IVV
```
