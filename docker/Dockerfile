# Stage 1: Combined builder for frontend assets and Go backend
FROM node:20-slim AS builder
WORKDIR /app

# Install build dependencies
RUN apt-get update && apt-get install -y ca-certificates curl gnupg build-essential

# Install Rustup and wasm target
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --default-toolchain stable
ENV PATH="/root/.cargo/bin:${PATH}"
RUN rustup target add wasm32-unknown-unknown
RUN cargo install wasm-bindgen-cli --version 0.2.113
RUN cargo install wasm-pack

# Install Go (Architecture Agnostic)
COPY --from=golang:1.24 /usr/local/go /usr/local/go
ENV PATH="/usr/local/go/bin:${PATH}"

# Copy version file for old-site build
COPY version.ts .

# Install frontend dependencies
COPY package.json package-lock.json ./
RUN npm install

COPY frontend/auth/admin/package.json frontend/auth/admin/package-lock.json ./auth/admin/
RUN npm --prefix auth/admin install

COPY frontend/old-site/package.json frontend/old-site/package-lock.json ./old-site/
RUN npm --prefix old-site install

COPY frontend/admin/package.json frontend/admin/package-lock.json ./admin/
RUN npm --prefix admin install

COPY frontend/marketing/package.json frontend/marketing/package-lock.json ./marketing/
RUN npm --prefix marketing install

# Copy environment variables before building SPAs so set-env scripts can read them
COPY .env/ ./.env/

# Accept MODE as a build argument so it's available during RUN steps
ARG MODE

# Build frontend applications
COPY frontend/auth/admin ./auth/admin
RUN npm run build-${MODE} --prefix auth/admin

COPY frontend/old-site ./old-site
RUN npm run build-old-site-${MODE} --prefix old-site

COPY frontend/admin ./admin
RUN npm run build-${MODE} --prefix admin

COPY frontend/marketing ./marketing
RUN npm run build-marketing-${MODE} --prefix marketing

# Copy the go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the backend source and build the Go application
COPY backend/go ./backend/go
RUN cd backend/go && CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2: Final application image
FROM gcr.io/distroless/base-debian11
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the built frontend assets from the builder stage
COPY --from=builder /build ./build

# Copy the environment file
COPY --from=builder /app/.env/ ./.env/

# Expose the port the backend runs on
EXPOSE 8080

# Set the command to run the backend service
CMD ["./main"]
