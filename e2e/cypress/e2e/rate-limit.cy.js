// Rate-limiting behaviour differs by MODE:
//   localhost  – permissive limits (1000 req/s, burst 5000) so that Docker-NAT'd
//                CI/dev traffic is never throttled. All requests should return 200.
//   other      – strict limits (5 req/s, burst 10). Rapid-fire requests must
//                eventually produce 429s.
describe('Rate Limiting', () => {
  const REQUEST_COUNT = 20;
  const mode = Cypress.env('mode') || 'localhost';

  it('should not rate-limit traffic in localhost/CI mode', { skip: mode !== 'localhost' }, () => {
    const statuses = [];

    for (let i = 0; i < REQUEST_COUNT; i++) {
      cy.request({
        url: '/v1/api/marketing/home',
        failOnStatusCode: false,
      }).then(response => {
        statuses.push(response.status);
      });
    }

    cy.then(() => {
      const rateLimitedCount = statuses.filter(s => s === 429).length;
      cy.log(`200 OK: ${statuses.filter(s => s === 200).length}, 429: ${rateLimitedCount}`);
      expect(rateLimitedCount).to.equal(
        0,
        'Localhost mode uses permissive rate limits — no request should be throttled',
      );
    });
  });

  it.skip('should return 429 Too Many Requests when burst limit is exceeded', { skip: mode === 'localhost' }, () => {
    const statuses = [];

    // Fire requests sequentially.
    // The Go backend is fast enough that these execute within one second,
    // triggering the strict rate limiter (burst of 10, 5 req/s regeneration).
    for (let i = 0; i < REQUEST_COUNT; i++) {
      cy.request({
        url: '/v1/api/marketing/home',
        failOnStatusCode: false,
      }).then(response => {
        statuses.push(response.status);
      });
    }

    cy.then(() => {
      const rateLimitedCount = statuses.filter(s => s === 429).length;
      const successCount = statuses.filter(s => s === 200).length;

      cy.log(`200 OK: ${successCount}, 429 Too Many Requests: ${rateLimitedCount}`);

      expect(rateLimitedCount).to.be.greaterThan(
        0,
        'Expected at least one request to be rate limited and return 429',
      );
      expect(successCount).to.be.greaterThan(
        0,
        'Expected the first burst of requests to succeed with 200',
      );
    });
  });
});
