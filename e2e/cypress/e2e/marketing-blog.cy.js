describe('Marketing Blog', () => {
  beforeEach(() => {
    // Intercept blog list request
    cy.intercept('GET', '**/api/marketing/blog?*', {
      statusCode: 200,
      body: [
        {
          id: '123',
          title: 'Test Blog Post',
          content: 'This is a test blog post content. It is long enough to have a snippet.',
          author: { id: '1', name: 'Test Author' },
          tags: [{ id: '1', name: 'test' }, { id: '2', name: 'cypress' }],
          createdAt: new Date().toISOString(),
          updatedAt: new Date().toISOString(),
          order: 1
        }
      ]
    }).as('getBlogPosts');

    // Intercept single blog post request
    cy.intercept('GET', '**/api/marketing/blog/123', {
      statusCode: 200,
      body: {
        id: '123',
        title: 'Test Blog Post',
        content: '# Test Blog Post\n\nThis is a test blog post content. It is long enough to have a snippet.',
        author: { id: '1', name: 'Test Author' },
        tags: [{ id: '1', name: 'test' }, { id: '2', name: 'cypress' }],
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
        order: 1
      }
    }).as('getBlogPost');

    cy.visit('/');
    cy.get('[data-testid="nav-blog"]').click();
  });

  it('should display blog posts and navigate to details page', () => {
    // Wait for initial load
    cy.wait('@getBlogPosts');

    // Ensure we are in the blog section
    cy.get('[data-testid="blog-header"]', { timeout: 10000 }).should('be.visible');

    // Check if the blog post card is displayed
    cy.get('app-blog').contains('Test Blog Post').should('be.visible');
    cy.get('app-blog').contains('This is a test blog post content.').should('be.visible');
    cy.get('app-blog').contains('#test').should('be.visible');

    // Click on the card
    cy.get('app-blog').contains('Test Blog Post').click();

    // Check URL change
    cy.url({ timeout: 10000 }).should('include', '/blog/123');

    // Wait for details page load
    cy.wait('@getBlogPost');

    // Verify details page content
    cy.get('h1').contains('Test Blog Post').should('be.visible');
    cy.contains('This is a test blog post content.').should('be.visible');
    cy.contains('Back to Blog').should('be.visible');

    // Verify back navigation
    cy.get('[data-testid="back-to-blog"]').should('be.visible').and('have.attr', 'href', '/blog');
    cy.get('[data-testid="back-to-blog"]').click();
    
    // Check URL change - using a more robust check
    cy.location('pathname', { timeout: 10000 }).should('eq', '/blog');
    
    // Ensure we are back in the blog section and it's visible
    cy.get('[data-testid="blog-header"]', { timeout: 10000 }).should('be.visible');
  });
});
