# # # # # # #
#  Builder  #
# # # # # # #

FROM golang:1.18.4-alpine3.16 as builder

# Make an /app directory in the container
RUN mkdir /app

# change default working directory in the container to /app
WORKDIR /app

# Cache dependencies
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Hot reloading with Air
RUN apk --no-cache add curl
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# copy all files in the root directory to /app in the container
COPY ./ ./

# Build the binary, call it main, put in the root directory in the container
RUN go build -o main /app/

# CMD [ "air" ]


# # # # # # #
#  Runner   #
# # # # # # #

FROM alpine:3.16 as runner
COPY --from=builder /app/main /main
ENTRYPOINT [ "/main" ]
