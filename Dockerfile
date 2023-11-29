# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory
WORKDIR /home/ahmadawab/go-training/folderprint

# Copy the current directory contents into the container
COPY . .

# Build the app
RUN go build -o main.sh

# Command to run the executable
CMD ["./main.sh"]
