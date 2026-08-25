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
# In Docker, SPAs are at /app/<spa>/ (1 level deep) but their set-env paths use ../../.env/
# (designed for local layout: frontend/<spa>/, 2 levels deep). Symlink bridges the gap.
RUN ln -s /app/.env /.env

# Accept MODE as a build argument so it's available during RUN steps
ARG MODE

# Accept New Relic browser agent config as build args (sourced from the mode's .env
# file via docker-compose --env-file). Mirrors the sed substitution done against
# frontend/*/src/index.html in .github/workflows/deploy.yml for CI/prod builds, so
# local docker-compose builds (which build the SPAs from source inside the
# container, bypassing lifecycle/build.sh's host-side injection) get real values too.
ARG NEW_RELIC_ACCOUNT_ID
ARG NEW_RELIC_TRUST_KEY
ARG NEW_RELIC_BROWSER_LICENSE_KEY
ARG NEW_RELIC_AGENT_ID_MARKETING
ARG NEW_RELIC_APP_ID_MARKETING
ARG NEW_RELIC_AGENT_ID_ADMIN
ARG NEW_RELIC_APP_ID_ADMIN
ARG NEW_RELIC_AGENT_ID_AUTH
ARG NEW_RELIC_APP_ID_AUTH
ARG NEW_RELIC_AGENT_ID_OLDSITE
ARG NEW_RELIC_APP_ID_OLDSITE

# Build frontend applications
COPY frontend/auth/admin ./auth/admin
RUN sed -i \
      -e "s|__NEW_RELIC_ACCOUNT_ID__|${NEW_RELIC_ACCOUNT_ID}|g" \
      -e "s|__NEW_RELIC_TRUST_KEY__|${NEW_RELIC_TRUST_KEY}|g" \
      -e "s|__NEW_RELIC_BROWSER_LICENSE_KEY__|${NEW_RELIC_BROWSER_LICENSE_KEY}|g" \
      -e "s|__NEW_RELIC_AGENT_ID_AUTH__|${NEW_RELIC_AGENT_ID_AUTH}|g" \
      -e "s|__NEW_RELIC_APP_ID_AUTH__|${NEW_RELIC_APP_ID_AUTH}|g" \
      auth/admin/src/index.html
RUN npm run build-${MODE} --prefix auth/admin

COPY frontend/old-site ./old-site
RUN sed -i \
      -e "s|__NEW_RELIC_ACCOUNT_ID__|${NEW_RELIC_ACCOUNT_ID}|g" \
      -e "s|__NEW_RELIC_TRUST_KEY__|${NEW_RELIC_TRUST_KEY}|g" \
      -e "s|__NEW_RELIC_BROWSER_LICENSE_KEY__|${NEW_RELIC_BROWSER_LICENSE_KEY}|g" \
      -e "s|__NEW_RELIC_AGENT_ID_OLDSITE__|${NEW_RELIC_AGENT_ID_OLDSITE}|g" \
      -e "s|__NEW_RELIC_APP_ID_OLDSITE__|${NEW_RELIC_APP_ID_OLDSITE}|g" \
      old-site/index.html
RUN npm run build-old-site-${MODE} --prefix old-site

COPY frontend/admin ./admin
RUN sed -i \
      -e "s|__NEW_RELIC_ACCOUNT_ID__|${NEW_RELIC_ACCOUNT_ID}|g" \
      -e "s|__NEW_RELIC_TRUST_KEY__|${NEW_RELIC_TRUST_KEY}|g" \
      -e "s|__NEW_RELIC_BROWSER_LICENSE_KEY__|${NEW_RELIC_BROWSER_LICENSE_KEY}|g" \
      -e "s|__NEW_RELIC_AGENT_ID_ADMIN__|${NEW_RELIC_AGENT_ID_ADMIN}|g" \
      -e "s|__NEW_RELIC_APP_ID_ADMIN__|${NEW_RELIC_APP_ID_ADMIN}|g" \
      admin/src/index.html
RUN npm run build-${MODE} --prefix admin

COPY frontend/marketing ./marketing
RUN sed -i \
      -e "s|__NEW_RELIC_ACCOUNT_ID__|${NEW_RELIC_ACCOUNT_ID}|g" \
      -e "s|__NEW_RELIC_TRUST_KEY__|${NEW_RELIC_TRUST_KEY}|g" \
      -e "s|__NEW_RELIC_BROWSER_LICENSE_KEY__|${NEW_RELIC_BROWSER_LICENSE_KEY}|g" \
      -e "s|__NEW_RELIC_AGENT_ID_MARKETING__|${NEW_RELIC_AGENT_ID_MARKETING}|g" \
      -e "s|__NEW_RELIC_APP_ID_MARKETING__|${NEW_RELIC_APP_ID_MARKETING}|g" \
      marketing/src/index.html
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
