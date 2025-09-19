
# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Set the environment variables
ENV PANEL_URL=""
ENV CLIENT_API_TOKEN=""

# Build the go application
RUN go build -o pterodactyl-exporter-go

# Expose the port
EXPOSE 9531

# Run the go application
CMD ["./pterodactyl-exporter-go"]

