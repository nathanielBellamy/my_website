describe('Admin App', () => {
  beforeEach(() => {
    // Assuming auth is bypassed on localhost or handled elsewhere
    cy.visit('/admin/')
    cy.contains('Admin Panel').should('be.visible')
  })

  it('should display the admin panel', () => {
    cy.contains('Home Content').should('be.visible')
  })

  it('should navigate between admin sections', () => {
    cy.get('[data-testid="nav-admin-blog"]').click()
    cy.url().should('include', '/admin/blog')
    cy.contains('Blog Posts').should('be.visible')

    cy.get('[data-testid="nav-admin-about"]').click()
    cy.url().should('include', '/admin/about')
    cy.contains('About Content').should('be.visible')

    cy.get('[data-testid="nav-admin-home"]').click()
    cy.url().should('include', '/admin/home')
    cy.contains('Home Content').should('be.visible')
  })

  it('should create and then delete home content', () => {
    const testTitle = 'E2E Test Title ' + Date.now()
    const testContent = 'This is content created by an E2E test.'

    // Create
    cy.get('[data-testid="create-new-home-content"]').click()
    cy.get('[data-testid="input-title"]').type(testTitle)
    cy.get('[data-testid="input-order"]').type('999')
    cy.get('[data-testid="input-content"]').type(testContent)
    cy.get('[data-testid="button-save"]').click()

    // Verify created
    cy.url().should('include', '/admin/home')
    
    // Switch to Inactive tab since we didn't set dates
    cy.get('[data-testid="status-inactive"]').click()
    
    cy.contains(testTitle).should('be.visible')

    // Delete
    cy.contains('li', testTitle).within(() => {
      cy.get('button').contains('Delete').click()
    })

    // Handle confirm dialog
    cy.on('window:confirm', () => true);

    // Verify deleted
    cy.contains(testTitle).should('not.exist')
  })
})
