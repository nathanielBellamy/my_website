describe('Marketing Blog', () => {
  const mockBlogPost = {
    id: '123',
    title: 'Test Blog Post',
    content: '# Test Blog Post\n\nThis is a test blog post content. It is long enough to have a snippet.',
    author: { id: '1', name: 'Test Author' },
    tags: [{ id: '1', name: 'test' }, { id: '2', name: 'cypress' }],
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
    order: 1
  };

  beforeEach(() => {
    // Intercept blog list request
    cy.intercept('GET', '**/marketing/blog?*', {
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

    // Intercept single blog post request (ID '123' has no dashes, so encodeId('123') === '123')
    cy.intercept('GET', '**/marketing/blog/123', {
      statusCode: 200,
      body: mockBlogPost
    }).as('getBlogPost');

    cy.visit('/blog');
  });

  it('should display blog posts and navigate to details page', () => {
    // Wait for initial load
    cy.wait('@getBlogPosts');

    // Ensure the section is in view to trigger fade-in
    cy.get('[data-testid="blog-header"]').scrollIntoView().should('be.visible');

    // Check if the blog post card is displayed
    // Allow more timeout for the fade-in animation to complete
    cy.contains('Test Blog Post', { timeout: 10000 }).should('be.visible');
    cy.contains('This is a test blog post content.').should('be.visible');
    cy.contains('#test').should('be.visible');

    // Click on the card
    cy.contains('Test Blog Post').click();

    // Check URL change
    cy.url().should('include', '/blog/123');

    // Wait for details page load
    cy.wait('@getBlogPost');

    // Verify details page content
    cy.contains('Back to Blog', { timeout: 10000 }).scrollIntoView().should('be.visible');
    cy.get('h1').contains('Test Blog Post').should('be.visible');
    cy.contains('This is a test blog post content.').should('be.visible');

    // Verify back navigation
    cy.contains('Back to Blog').click();
    
    // Ensure the blog list is visible again
    cy.get('[data-testid="blog-header"]').scrollIntoView().should('be.visible');
    cy.contains('Test Blog Post', { timeout: 10000 }).should('be.visible');
    
    // Scroll to the blog header to ensure we are in the right section (helps with URL update)
    cy.get('[data-testid="blog-header"]').scrollIntoView();
    cy.url().should('include', '/blog');
  });

  it('should load blog post details on direct URL navigation', () => {
    // Intercept the single blog post request for direct navigation
    cy.intercept('GET', '**/marketing/blog/123', {
      statusCode: 200,
      body: mockBlogPost
    }).as('getBlogPostDirect');

    // Navigate directly to the content-details URL (simulates browser refresh)
    cy.visit('/blog/123');

    // Wait for the API call to complete
    cy.wait('@getBlogPostDirect');

    // Verify the page renders the blog post content correctly
    cy.contains('Back to Blog', { timeout: 10000 }).scrollIntoView().should('be.visible');
    cy.get('h1').contains('Test Blog Post').should('be.visible');
    cy.contains('This is a test blog post content.').should('be.visible');
  });
});
