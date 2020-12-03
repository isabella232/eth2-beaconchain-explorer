# The dockerfile is currently still WIP and might be broken
FROM golang:1.14-alpine AS build-env
RUN go get -d -v ./...
RUN apk --no-cache add build-base git musl-dev linux-headers npm
ADD . /src
RUN cd /src && make all

# final stage
FROM alpine
WORKDIR /app
RUN apk --no-cache add libstdc++ libgcc
COPY --from=build-env /src/bin /app/

COPY ./entrypoint.sh /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

EXPOSE 3333

#CMD ["./explorer -config config.yml"]
CMD ["./explorer"]
