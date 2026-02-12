# GitHub Actions Workflows

This directory contains the CI/CD workflows for the my_website project.

## pr-checks.yml

Runs comprehensive checks on all pull requests to ensure code quality and functionality.

### Jobs

#### 1. Backend Tests (`backend-tests`)
- **Purpose**: Validates Go backend functionality
- **Runs**: Go tests for all packages in `backend/go`
- **Requirements**: Go 1.20
- **Triggers**: All PRs
- **Duration**: ~30 seconds

**Packages Tested:**
- Main backend routes and handlers
- Authentication (dev auth, cookie validation)
- Environment configuration
- Marketing controller (blog posts API)
- Old site controller (file serving, WebSocket)

#### 2. Marketing Build & Test (`marketing-build-test`)
- **Purpose**: Validates Angular marketing app
- **Runs**: Jest unit tests and production build
- **Requirements**: Node.js 20
- **Triggers**: All PRs
- **Duration**: ~1-2 minutes

**What's Tested:**
- Service layer tests (home, about, blog, groove-jr services)
- Component tests (pages and shared components)
- Production build verification

**Test Framework:** Jest with Angular Testing Library (no TestBed)

#### 3. Docker Build (`docker-build`)
- **Purpose**: Ensures the complete application can be containerized
- **Builds**: Full multi-stage Docker image with all components
- **Requirements**: Docker Buildx
- **Triggers**: All PRs
- **Duration**: ~15-20 minutes (with caching)

**What's Built:**
- Node.js dependencies for all frontend apps
- Rust/WASM compilation for old-site
- Go backend compilation (linux/arm64)
- All frontend SPAs (auth, old-site, marketing)

**Image Platform:** linux/arm64 (for NixOS VM deployment)

### Workflow Triggers

The workflow triggers on:
```yaml
on:
  pull_request:
    branches:
      - '**'
```

This means any PR to any branch will run these checks.

### Caching

To speed up builds, the workflow caches:
- Go modules (`~/go/pkg/mod`)
- NPM packages (`~/.npm`)

### Future Enhancements

When the `admin` Angular app is created, add:
```yaml
admin-build-test:
  name: Admin App Build & Test
  runs-on: ubuntu-latest
  steps:
    - name: Checkout code
      uses: actions/checkout@v4
    - name: Set up Node.js
      uses: actions/setup-node@v4
      with:
        node-version: '20'
    - name: Install dependencies
      run: npm ci
      working-directory: ./admin
    - name: Run tests
      run: npm test
      working-directory: ./admin
    - name: Build admin app
      run: npm run build-localhost
      working-directory: ./admin
```

### Debugging Failed Workflows

If a workflow fails:

1. **Backend Tests Fail:**
   - Check the test output in the Actions tab
   - Run tests locally: `cd backend/go && MODE=localhost go test -v ./...`
   - Ensure all mocks are properly configured

2. **Marketing Tests Fail:**
   - Check Jest output for specific test failures
   - Run locally: `cd marketing && npm test`
   - Verify all services and components are tested with Angular Testing Library patterns

3. **Docker Build Fails:**
   - Check which stage failed in the multi-stage build
   - Common issues:
     - Missing dependencies in package.json files
     - Rust compilation errors (check old-site/src-rust)
     - Go compilation errors
     - Network issues downloading dependencies

### Local Testing

To test workflows locally before pushing:

```bash
# Backend tests
cd backend/go
MODE=localhost go test -v ./...

# Marketing tests and build
cd marketing
npm ci
npm test
npm run build-localhost

# Docker build (note: requires linux/arm64 or Rosetta on Apple Silicon)
docker build -f docker/Dockerfile -t my_website:test .
```
