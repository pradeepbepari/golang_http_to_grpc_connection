{ pkgs ? import (builtins.fetchGit {
  name = "dev-go";
  url = "https://github.com/NixOS/nixpkgs";
  ref = "refs/heads/nixpkgs-unstable";
  rev = "9355fa86e6f27422963132c2c9aeedb0fb963d93";
}) {} }:

with pkgs;

mkShell {
  buildInputs = [
    clang-tools
    gitlint
    gnupg
    go_1_23
    go-tools
    golangci-lint
    goreleaser
    gosec
    gotools
    gofumpt
    golint
    openapi-generator-cli
    postgresql
    pre-commit
    protobuf
    protoc-gen-go
    protoc-gen-go-grpc
    awscli2
    sops
  ];

  shellHook = ''
    export PATH="$(go env GOPATH)/bin:$PATH"
    export NIXPKGS_ALLOW_UNFREE=1

    # Install pre-commit hooks
    pre-commit install

    echo "[INFO] Checking and installing Go proto plugins..."

    # Ensure proto plugins are installed (idempotent)
    which protoc-gen-grpc-gateway || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
    which protoc-gen-openapiv2 || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
    which protoc-gen-go || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    which protoc-gen-go-grpc || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

    # Bonus tools for generation and mocking
    which enumer || go install github.com/dmarkham/enumer@latest
    which gocritic || go install github.com/go-critic/go-critic/cmd/gocritic@latest
    which goreturns || go install github.com/sqs/goreturns@latest
    which swag || go install github.com/swaggo/swag/cmd/swag@latest
    which mockgen || go install github.com/golang/mock/mockgen@v1.6.0

    # Git config
    git config --local include.path ../.gitconfig || true

    clear
  '';
}
