# Use the official C++ image to build the application
FROM gcc:latest

# Install dependencies (in this case, for JSON processing)
RUN apt-get update && apt-get install -y \
    cmake \
    libjsoncpp-dev \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . /app

# Build the application
RUN g++ -o crud-api main.cpp -I. -L. -ljsoncpp -std=c++11

# Expose the port that the app will listen on
EXPOSE 8080

# Run the server
CMD ["./crud-api"]
