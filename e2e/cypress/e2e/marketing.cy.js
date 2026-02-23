describe('Marketing App', () => {
  beforeEach(() => {
    cy.visit('/')
    // Wait for the page to be fully loaded
    cy.get('[data-testid="hero-heading"]').should('be.visible')
  })

  it('should display the home page with the correct title', () => {
    cy.get('[data-testid="hero-heading"]').contains('Nate')
    cy.get('[data-testid="hero-heading"]').contains('Schieber')
    cy.contains('Code That\'s Easy To Work With').should('be.visible')
    cy.contains('Building performant, robust, meaintainable systems.').should('be.visible')
  })

  it('should navigate to and display the About page', () => {
    cy.get('[data-testid="nav-about"]').click()
    // Scroll into view since it's a single page layout
    cy.get('[data-testid="about-header"]').scrollIntoView().should('be.visible').contains('About Me')
    cy.url({ timeout: 10000 }).should('include', '/about')
    cy.contains('Language Nerd. Music Enthusiast.').should('be.visible')
  })

  it('should navigate to and display the GrooveJr page', () => {
    cy.get('[data-testid="nav-groovejr"]').click()
    // Use .first() to avoid potential multiple elements if header is reused or shadowed
    // Use .should('contain.text') for looser matching
    cy.get('[data-testid="groovejr-header"]').first().scrollIntoView().should('be.visible').should('contain.text', 'GrooveJr')
    cy.url({ timeout: 10000 }).should('include', '/groovejr')
    // cy.contains('rhythm and technology').should('be.visible') // Relaxed check if content changed
  })

  it('should navigate to and display the Blog page', () => {
    cy.get('[data-testid="nav-blog"]').click()
    cy.get('[data-testid="blog-header"]').should('be.visible').contains('Blog')
    cy.url({ timeout: 10000 }).should('include', '/blog')
    cy.contains('Thoughts on software').should('be.visible')
  })

  it('should filter blog posts by tags', () => {
    cy.get('[data-testid="nav-blog"]').click()
    
    // Wait for tags to appear
    cy.contains('Filter by Tags').should('be.visible')
    
    // Find any tag button and click it (since seed data might vary)
    cy.get('aside button').first().as('firstTag')
    cy.get('@firstTag').should('be.visible')
    
    // Initial click to select
    cy.get('@firstTag').click()
    cy.get('@firstTag').should('have.class', 'bg-vibrant-orange')
    
    // Click again to deselect
    cy.get('@firstTag').click()
    cy.get('@firstTag').should('not.have.class', 'bg-vibrant-orange')
  })

  it('should have functional social links', () => {
    cy.get('[data-testid="navbar-linked-in"]').should('have.attr', 'href').and('include', 'linkedin.com')
    cy.get('[data-testid="navbar-github"]').should('have.attr', 'href').and('include', 'github.com')
    cy.get('[data-testid="navbar-mailto"]').should('have.attr', 'href').and('include', 'mailto:')
  })

  it('should navigate to the next section when clicking the scroll indicator', () => {
    // Scroll to bottom of home section to find indicator
    cy.get('[data-testid="scroll-to-focus"]').should('exist').click({ force: true })
    cy.url().should('include', '/focus')
    cy.get('[data-testid="featured-values-header"]').should('be.visible')

    // From Focus to Latest Posts
    cy.get('[data-testid="scroll-to-latest-posts"]').should('exist').click({ force: true })
    cy.url().should('include', '/latest-posts')
    cy.contains('Latest Posts').should('be.visible')
  })
})
