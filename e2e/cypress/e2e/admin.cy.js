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

    cy.get('[data-testid="nav-admin-groovejr"]').click()
    cy.url().should('include', '/admin/groovejr')
    cy.contains('GrooveJr Content').should('be.visible')

    cy.get('[data-testid="nav-admin-home"]').click()
    cy.url().should('include', '/admin/home')
    cy.contains('Home Content').should('be.visible')
  })

  const performCrud = (section, createTestid, title, content) => {
    const testTitle = `${title} ${Date.now()}`
    
    // Create
    cy.get(`[data-testid="${createTestid}"]`).click()
    cy.get('[data-testid="input-title"]').type(testTitle)
    cy.get('[data-testid="input-order"]').type('999')
    cy.get('[data-testid="input-content"]').type(content)
    cy.get('[data-testid="button-save"]').click()

    // Verify created in Inactive tab (since no dates set)
    cy.url().should('include', `/admin/${section}`)
    cy.get('[data-testid="status-inactive"]').click()
    cy.contains(testTitle).should('be.visible')

    // Edit
    cy.contains('li', testTitle).within(() => {
      cy.contains('Edit').click()
    })
    const updatedTitle = `${testTitle} UPDATED`
    cy.get('[data-testid="input-title"]').clear().type(updatedTitle)
    cy.get('[data-testid="button-save"]').click()

    // Verify updated
    cy.get('[data-testid="status-inactive"]').click()
    cy.contains(updatedTitle).should('be.visible')

    // Delete
    cy.contains('li', updatedTitle).within(() => {
      cy.get('button').contains('Delete').click()
    })
    cy.on('window:confirm', () => true)
    cy.contains(updatedTitle).should('not.exist')
  }

  it('should perform CRUD for Home content', () => {
    cy.get('[data-testid="nav-admin-home"]').click()
    performCrud('home', 'create-new-home-content', 'E2E Home', 'Home content test.')
  })

  it('should perform CRUD for About content', () => {
    cy.get('[data-testid="nav-admin-about"]').click()
    performCrud('about', 'create-new-about-content', 'E2E About', 'About content test.')
  })

  it('should perform CRUD for GrooveJr content', () => {
    cy.get('[data-testid="nav-admin-groovejr"]').click()
    performCrud('groovejr', 'create-new-groovejr-content', 'E2E GrooveJr', 'GrooveJr content test.')
  })

  it('should perform CRUD for Blog posts', () => {
    cy.get('[data-testid="nav-admin-blog"]').click()
    const testTitle = `E2E Blog ${Date.now()}`
    
    // Create
    cy.get('[data-testid="create-new-blog-post"]').click()
    cy.get('[data-testid="input-title"]').type(testTitle)
    cy.get('[data-testid="input-order"]').type('-100')
    cy.get('[data-testid="input-content"]').type('Blog post content.')
    cy.get('[data-testid="input-author-name"]').type('E2E Tester')
    cy.get('[data-testid="input-tags"]').type('e2e, test')
    // Clear default activatedAt to make it inactive
    cy.get('[data-testid="input-activatedAt"]').clear()
    cy.get('[data-testid="button-save"]').click()

    // Verify created
    cy.url().should('include', '/admin/blog')
    cy.contains('Blog Posts').should('be.visible')
    cy.get('[data-testid="status-inactive"]').click()
    cy.contains(testTitle, { timeout: 10000 }).should('be.visible')

    // Edit
    cy.contains('li', testTitle).within(() => {
      cy.contains('Edit').click()
    })
    const updatedTitle = `${testTitle} UPDATED`
    cy.get('[data-testid="input-title"]').clear().type(updatedTitle)
    cy.get('[data-testid="input-order"]').clear().type('-100') // Ensure order stays at top
    cy.get('[data-testid="button-save"]').click()
    cy.wait(1000)

    // Verify updated
    cy.reload() // Ensure fresh list
    cy.get('[data-testid="status-inactive"]').click()
    cy.wait(1000)
    cy.contains(updatedTitle).should('be.visible')

    // Delete
    cy.contains('li', updatedTitle).within(() => {
      cy.get('button').contains('Delete').click()
    })
    cy.on('window:confirm', () => true)
    cy.contains(updatedTitle).should('not.exist')
  })
})
