FROM golang:1.13 AS build-env
COPY . /build
WORKDIR /build

# execute tests
RUN go test ./... -cover

## ONLY tests so far, building to be added later
# execute build
RUN export GIT_COMMIT=$(git rev-parse HEAD) && \
    export GIT_REPOSITORY=$(git config --get remote.origin.url) && \
    CGO_ENABLED=0 go build \
        -ldflags \
            "-X github.com/SAP/jenkins-library/cmd.GitCommit=${GIT_COMMIT} \
            -X github.com/SAP/jenkins-library/pkg/log.LibraryRepository=${GIT_REPOSITORY}" \
        -o piper

FROM maven:3-jdk-8-slim
COPY --from=build-env /build/piper /bin/piper

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
