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

  it('should navigate to the About page', () => {
    cy.get('[data-testid="nav-about"]').click()
    cy.contains('About').should('be.visible')
    // URL might take a moment to stabilize due to scroll logic
    cy.url({ timeout: 10000 }).should('include', '/about')
  })

  it('should navigate to the GrooveJr page', () => {
    cy.get('[data-testid="nav-groovejr"]').click()
    cy.contains('GrooveJr').should('be.visible')
    cy.url({ timeout: 10000 }).should('include', '/groovejr')
  })

  it('should navigate to the Blog page', () => {
    cy.get('[data-testid="nav-blog"]').click()
    cy.get('[data-testid="blog-header"]').should('be.visible')
    cy.url({ timeout: 10000 }).should('include', '/blog')
  })

  it('should have functional social links', () => {
    cy.get('[data-testid="navbar-linked-in"]').should('have.attr', 'href').and('include', 'linkedin.com')
    cy.get('[data-testid="navbar-github"]').should('have.attr', 'href').and('include', 'github.com')
    cy.get('[data-testid="navbar-mailto"]').should('have.attr', 'href').and('include', 'mailto:')
  })
})
