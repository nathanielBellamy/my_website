describe('Marketing App', () => {
  beforeEach(() => {
    cy.visit('/')
    // Wait for the page to be fully loaded
    cy.get('[data-testid="hero-heading"]').should('be.visible')
  })

  it('should display the home page with the correct title', () => {
    cy.get('[data-testid="hero-heading"]').contains('Nate')
    cy.get('[data-testid="hero-heading"]').contains('Schieber')
    cy.contains('Clean Coding. Code Cleaning. Software Builder.').should('be.visible')
  })

  it('should navigate to and display the About page', () => {
    cy.get('[data-testid="nav-about"]').click()
    // Scroll into view since it's a single page layout
    cy.get('[data-testid="about-header"]').scrollIntoView().should('be.visible').contains('About Me')
    cy.url({ timeout: 10000 }).should('include', '/about')
    cy.contains('Engineer by trade').should('be.visible')
  })

  it('should navigate to and display the GrooveJr page', () => {
    cy.get('[data-testid="nav-groovejr"]').click()
    cy.get('[data-testid="groovejr-header"]').scrollIntoView().should('be.visible').contains('GrooveJr')
    cy.url({ timeout: 10000 }).should('include', '/groovejr')
    cy.contains('rhythm and technology').should('be.visible')
  })

  it('should navigate to and display the Blog page', () => {
    cy.get('[data-testid="nav-blog"]').click()
    cy.get('[data-testid="blog-header"]').scrollIntoView().should('be.visible').contains('Blog')
    cy.url({ timeout: 10000 }).should('include', '/blog')
    cy.contains('software engineering').should('be.visible')
  })

  it('should filter blog posts by tags', () => {
    cy.get('[data-testid="nav-blog"]').click()
    
    // Wait for tags to appear
    cy.contains('Filter by Tags').should('be.visible')
    
    // Find a tag (e.g., 'Go') and click it
    // We use a flexible selector because exact tag names depend on seed data
    cy.get('button').contains('Go').as('goTag')
    cy.get('@goTag').should('be.visible')
    
    // Initial click to select
    cy.get('@goTag').click()
    cy.get('@goTag').should('have.class', 'bg-vibrant-orange')
    
    // Verify URL or content update (mocking backend or checking DOM)
    // Checking if 'Getting Started with Go' is visible (from seed)
    cy.contains('Getting Started with Go').should('be.visible')
    
    // Click again to deselect
    cy.get('@goTag').click()
    cy.get('@goTag').should('not.have.class', 'bg-vibrant-orange')
  })

  it('should have functional social links', () => {
    cy.get('[data-testid="navbar-linked-in"]').should('have.attr', 'href').and('include', 'linkedin.com')
    cy.get('[data-testid="navbar-github"]').should('have.attr', 'href').and('include', 'github.com')
    cy.get('[data-testid="navbar-mailto"]').should('have.attr', 'href').and('include', 'mailto:')
  })
})
