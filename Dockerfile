FROM golang:1.13.1 AS build
ARG service_name

RUN mkdir /app
COPY . /app
RUN ls /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -mod vendor -o $service_name

FROM alpine:3.9
ARG service_name

RUN apk add --update tzdata
RUN apk add --update ca-certificates
RUN apk add --update bash
ENV SERVICE_NAME=$service_name

COPY --from=build /app/$service_name /app/$service_name
RUN chmod +x /app/$service_name
RUN bash -c 'printf "#! /bin/bash\n/app/${SERVICE_NAME} \$1" > entrypoint.sh'
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["./entrypoint.sh", "serve"]
