# Stage 1: Frontend asset builder
FROM ubuntu:22.04 AS builder-assets
WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates curl gnupg build-essential
RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_20.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list
RUN apt-get update && apt-get install -y nodejs

# Install rustup and wasm target
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --default-toolchain stable
ENV PATH="/root/.cargo/bin:${PATH}"
RUN rustup target add wasm32-unknown-unknown

RUN cargo install wasm-bindgen-cli

# Copy all package.json and package-lock.json files
COPY package.json package-lock.json ./
COPY auth/dev/package.json auth/dev/package-lock.json ./auth/dev/
COPY old-site/package.json old-site/package-lock.json ./old-site/
COPY marketing/package.json marketing/package-lock.json ./marketing/

# Install dependencies for all frontends
RUN npm install
RUN npm --prefix auth/dev install
RUN npm --prefix old-site install
RUN npm --prefix marketing install

# Copy the rest of the source code
COPY . .

# Build all frontend applications
RUN cd auth/dev && npm run build-remotedev
RUN cd old-site && npm run build-old-site-remotedev
RUN cd marketing && npm run build-marketing-remotedev

# Fix asset paths
RUN sed -i 's/\/assets/\.\/assets/g' build/auth/dev/index.html
RUN sed -i -e 's/src="\/assets/src="\/old-site\/assets/g' -e 's/href="\/assets/href="\/old-site\/assets/g' build/old-site/index.html

# Stage 2: Go backend builder
FROM golang:1.24-alpine AS builder-go
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the backend source code
COPY backend/go ./backend/go

# Build the Go application
RUN cd backend/go && CGO_ENABLED=0 GOOS=linux go build -o /app/main main.go

# Stage 3: Final application image
FROM ubuntu:latest
WORKDIR /app

# Copy the built Go binary from the Go builder stage
COPY --from=builder-go /app/main .

# Copy the built frontend assets from the frontend builder stage
COPY --from=builder-assets /app/build ./build

# Copy the environment files
COPY .env.* .

# Expose the port the backend runs on
EXPOSE 8080

# Set the command to run the backend service
CMD ["/app/main"]
