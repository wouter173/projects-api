FROM golang:1.15.6-alpine AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -o projects-api .

FROM busybox:1.32.1

WORKDIR /app
COPY --from=build /src/projects-api .

ENV PROJECTS_PATH ./projects

CMD ["./projects-api"]
