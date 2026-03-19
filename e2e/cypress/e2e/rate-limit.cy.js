describe('Rate Limiting', () => {
  it('should return 429 Too Many Requests when burst limit is exceeded', () => {
    const statuses = [];
    const requestCount = 20; // 20 to ensure we blow past the burst limit of 10 and 5/sec regeneration

    // Fire requests sequentially. 
    // Our Go backend is fast enough that these should easily execute within a single second,
    // triggering the rate limiter which allows a burst of 10.
    for (let i = 0; i < requestCount; i++) {
      cy.request({
        url: '/v1/api/marketing/home',
        failOnStatusCode: false // Prevent Cypress from failing the test when we hit 429
      }).then(response => {
        statuses.push(response.status);
      });
    }

    // After all requests in the Cypress queue have completed, verify the statuses
    cy.then(() => {
      const rateLimitedCount = statuses.filter(s => s === 429).length;
      const successCount = statuses.filter(s => s === 200).length;
      
      cy.log(`200 OK: ${successCount}, 429 Too Many Requests: ${rateLimitedCount}`);
      
      expect(rateLimitedCount).to.be.greaterThan(
        0, 
        'Expected at least one request to be rate limited and return 429'
      );
      expect(successCount).to.be.greaterThan(
        0, 
        'Expected the first burst of requests to succeed with 200'
      );
    });
  });
});
