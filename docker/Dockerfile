# Stage 1: Combined builder for frontend assets and Go backend
FROM node:20-slim AS builder
WORKDIR /app

# Install build dependencies
RUN apt-get update && apt-get install -y ca-certificates curl gnupg build-essential wget

# Install Rustup and wasm target
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --default-toolchain stable
ENV PATH="/root/.cargo/bin:${PATH}"
RUN rustup target add wasm32-unknown-unknown
RUN cargo install wasm-bindgen-cli

# Install Go for arm64
RUN wget https://golang.org/dl/go1.24.0.linux-arm64.tar.gz
RUN tar -C /usr/local -xzf go1.24.0.linux-arm64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Copy version file for old-site build
COPY version.ts .

# Install frontend dependencies
COPY package.json package-lock.json ./
RUN npm install

COPY auth/admin/package.json auth/admin/package-lock.json ./auth/admin/
RUN npm --prefix auth/admin install

COPY old-site/package.json old-site/package-lock.json ./old-site/
RUN npm --prefix old-site install

COPY admin/package.json admin/package-lock.json ./admin/
RUN npm --prefix admin install

COPY marketing/package.json marketing/package-lock.json ./marketing/
RUN npm --prefix marketing install

# Build frontend applications
COPY auth/admin ./auth/admin
RUN npm run build-localhost --prefix auth/admin

COPY old-site ./old-site
RUN npm run build-old-site-localhost --prefix old-site

COPY admin ./admin
RUN npm run build-localhost --prefix admin

COPY marketing ./marketing
RUN npm run build-localhost --prefix marketing

# Copy the go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the backend source and build the Go application
COPY backend/go ./backend/go
RUN cd backend/go && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/main .

# Stage 2: Final application image
FROM gcr.io/distroless/base-debian11
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the built frontend assets from the builder stage
COPY --from=builder /app/build ./build

# Copy the environment file
COPY --from=builder /app/backend/go/.env.localhost .

# Expose the port the backend runs on
EXPOSE 8080

# Set the command to run the backend service
CMD ["./main"]
