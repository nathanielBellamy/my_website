# Security Policy

## Supported Versions

Only the latest version of the `my_website` project is supported for security updates.

## Reporting a Vulnerability

If you discover a security vulnerability in this project, please report it via GitHub Issues or contact the maintainer directly.

## Security Architecture

The project follows several security best practices:

- **Defense in Depth**: Multiple layers of protection including MFA (OTP/TOTP), SHA256 password hashing with challenges, and secure cookie configurations.
- **Least Privilege**: The `marketing` app has read-only access to the database, while the `admin` app is guarded by a robust authentication layer.
- **Zero Trust Principles**: Verification is required for all administrative actions.

### Authentication Flow
1. **Challenge**: Client requests a challenge.
2. **Password**: Client sends SHA256(ADMIN_PW + challenge).
3. **MFA**: After password validation, an OTP is sent via email or verified via TOTP.
4. **Session**: A secure, HTTP-only session cookie is issued.

## Automated Security Checks

We run the following security scans on every push:

- **Trivy**: Vulnerability scanning for filesystem and dependencies.
- **Gosec**: Static analysis for Go code to detect common security pitfalls.
- **Squirrel**: Dynamic website audit for security headers, SSL, and other web-level vulnerabilities.
- **npm audit**: Checking for known vulnerabilities in Node.js dependencies.

## Security Roadmap

- [ ] Implement a full Content Security Policy (CSP).
- [ ] Add explicit CSRF protection for all mutating administrative endpoints.
- [ ] Implement Rate Limiting for all API endpoints.
- [ ] Enhance audit logging for all administrative actions.
