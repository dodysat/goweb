# test build 1 2
# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Dody Satria <dody.satria@gmail.com>"

RUN echo "[url \"git@bitbucket.org:\"]\n\tinsteadOf = https://bitbucket.org/" >> /root/.gitconfig

# Set the Current Working Directory inside the container
WORKDIR /app

RUN go env -w GOPRIVATE=bitbucket.org/klopos

# Copy go mod and sum files
COPY go.mod go.sum ./

# RUN go get -v

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######

FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add tzdata
ENV TZ Asia/Jakarta

WORKDIR /root/

RUN mkdir log config config/files config/reference


# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/config/files/dev.json /root/config/files
COPY --from=builder /app/config/reference/ref.json /root/config/reference



# Expose port 80 to the outside world
EXPOSE 8000
EXPOSE 9090


# Command to run the executable]
CMD ["./main"]