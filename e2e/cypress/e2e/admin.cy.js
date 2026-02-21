describe('Admin App', () => {
  beforeEach(() => {
    // Assuming auth is bypassed on localhost or handled elsewhere
    cy.visit('/admin/')
    cy.contains('Admin Panel', { timeout: 10000 }).should('be.visible')
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
    const now = Date.now();
    const testTitle = `E2E Blog ${now}`
    
    // Create
    cy.get('[data-testid="create-new-blog-post"]').click()
    cy.get('[data-testid="input-title"]').type(testTitle)
    cy.get('[data-testid="input-order"]').type('999')
    cy.get('[data-testid="input-content"]').type('Blog post content.')
    cy.get('[data-testid="input-author-name"]').type('E2E Tester')
    cy.get('[data-testid="input-tags"]').type(`e2e, test, ${now}`)
    cy.get('[data-testid="button-save"]').click()

    // Verify created
    cy.url().should('include', '/admin/blog')
    cy.contains('Blog Posts').should('be.visible')
    cy.get('[data-testid="blog-search-tags-input"]').type(`${now}`)
    cy.contains('button', now).click()
    cy.contains(testTitle, { timeout: 10000 }).should('be.visible')

    // Edit
    cy.contains('li', testTitle).within(() => {
      cy.contains('Edit').click()
    })
    const updatedTitle = `${testTitle} UPDATED`
    cy.get('[data-testid="input-title"]').clear().type(updatedTitle)
    cy.get('[data-testid="button-save"]').click()

    // Verify updated
    cy.get('[data-testid="blog-search-tags-input"]').type(`${now}`)
    cy.contains('button', now).click()
    cy.contains(updatedTitle).should('be.visible')

    // Delete
    cy.contains('li', updatedTitle).within(() => {
      cy.get('button').contains('Delete').click()
    })
    cy.on('window:confirm', () => true)
    cy.get('[data-testid="blog-search-tags-input"]').type(`${now}`)
    cy.contains('button', now).click()
    cy.contains(updatedTitle).should('not.exist')
  })

  it('should support tag assignment in blog post editor', () => {
    cy.get('[data-testid="nav-admin-blog"]').click()
    cy.get('[data-testid="create-new-blog-post"]').click()
    
    // Check if tag suggestions are visible
    cy.contains('Suggestions').should('be.visible')

    cy.get('[data-testid="blog-form-search-tags-input"]').type('Go')
    cy.contains('button', 'Go').click()
    
    // Double click a tag (assuming 'Go' exists from seed)
    cy.get('button').contains('Go').dblclick()
    
    // Verify tag is added to input
    cy.get('[data-testid="input-tags"]').should('have.value', 'Go')
    
    // Add another tag via double click
    cy.get('[data-testid="blog-form-search-tags-input"]').clear()
    cy.get('[data-testid="blog-form-search-tags-input"]').type('PostgreSQL')
    cy.get('button').contains('PostgreSQL').dblclick()
    // Value might be "Go, PostgreSQL" or just "PostgreSQL" if my split logic is weird, but expected is comma separated
    cy.get('[data-testid="input-tags"]').should('contain.value', 'Go')
    cy.get('[data-testid="input-tags"]').should('contain.value', 'PostgreSQL')
  })
})
