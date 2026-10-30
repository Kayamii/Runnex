# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.27-alpine AS build

WORKDIR /src

RUN apk add --no-cache make

# go.sum is matched with a glob because it only exists once dependencies are added.
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=unknown
ARG DATE=unknown
RUN make build VERSION="${VERSION}" COMMIT="${COMMIT}" DATE="${DATE}"

# Runtime stage: static binary on a minimal, non-root base image.
FROM gcr.io/distroless/static:nonroot

COPY --from=build /src/bin/runnex /usr/local/bin/runnex

USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/runnex"]
CMD ["help"]
