describe('Admin App', () => {
  beforeEach(() => {
    // Assuming auth is bypassed on localhost or handled elsewhere
    cy.visit(Cypress.env('adminUrl'))
    cy.contains('Admin Panel').should('be.visible')
  })

  it('should display the dashboard by default', () => {
    cy.url().should('include', '/dashboard')
    cy.get('[data-testid="dashboard-title"]').should('be.visible')
    cy.contains('System Dashboard').should('be.visible')
  })

  it('should display dashboard health cards', () => {
    cy.get('[data-testid="dashboard-uptime-card"]', { timeout: 10000 }).should('be.visible')
    cy.get('[data-testid="dashboard-uptime-card"]').contains('Uptime')
    cy.get('[data-testid="dashboard-memory-card"]').should('be.visible')
    cy.get('[data-testid="dashboard-memory-card"]').contains('Memory')
    cy.get('[data-testid="dashboard-goroutines-card"]').should('be.visible')
    cy.get('[data-testid="dashboard-goroutines-card"]').contains('Goroutines')
    cy.get('[data-testid="dashboard-db-card"]').should('be.visible')
    cy.get('[data-testid="dashboard-db-card"]').contains(/Connected|Disconnected/)
  })

  it('should display dashboard system info', () => {
    cy.get('[data-testid="dashboard-system-info"]', { timeout: 10000 }).should('be.visible')
    cy.get('[data-testid="dashboard-system-info"]').contains('System Info')
    cy.get('[data-testid="dashboard-system-info"]').contains('Go Version')
    cy.get('[data-testid="dashboard-system-info"]').contains('CPUs')
  })

  it('should display dashboard quick links', () => {
    cy.get('[data-testid="dashboard-quick-links"]').should('be.visible')
    cy.get('[data-testid="dashboard-link-logs"]').should('be.visible')
    cy.get('[data-testid="dashboard-link-grafana"]').should('be.visible')
    cy.get('[data-testid="dashboard-link-work"]').should('be.visible')
    cy.get('[data-testid="dashboard-link-blog"]').should('be.visible')
    cy.get('[data-testid="dashboard-link-gallery"]').should('be.visible')
  })

  it('should display dashboard recent errors section', () => {
    cy.get('[data-testid="dashboard-recent-errors"]', { timeout: 10000 }).should('be.visible')
    cy.get('[data-testid="dashboard-recent-errors"]').contains('Recent Errors')
  })

  it('should have a working refresh button', () => {
    cy.get('[data-testid="dashboard-refresh-btn"]').should('be.visible')
    cy.get('[data-testid="dashboard-refresh-btn"]').click()
    cy.get('[data-testid="dashboard-uptime-card"]', { timeout: 10000 }).should('be.visible')
  })

  it('should navigate to logs via quick link', () => {
    cy.get('[data-testid="dashboard-link-logs"]').click()
    cy.url().should('include', '/logs')
    cy.get('[data-testid="logs-page-title"]').should('be.visible')
    cy.contains('System Logs').should('be.visible')
  })

  it('should display logs page controls', () => {
    cy.get('[data-testid="nav-admin-logs"]').click()
    cy.url().should('include', '/logs')

    cy.get('[data-testid="logs-live-btn"]').should('be.visible')
    cy.get('[data-testid="logs-history-btn"]').should('be.visible')
    cy.get('[data-testid="logs-autoscroll-btn"]').should('be.visible')
    cy.get('[data-testid="logs-clear-btn"]').should('be.visible')
    cy.get('[data-testid="logs-search-input"]').should('be.visible')
  })

  it('should display logs level filter chips', () => {
    cy.get('[data-testid="nav-admin-logs"]').click()

    cy.get('[data-testid="logs-level-chip-info"]').should('be.visible')
    cy.get('[data-testid="logs-level-chip-warn"]').should('be.visible')
    cy.get('[data-testid="logs-level-chip-error"]').should('be.visible')
    cy.get('[data-testid="logs-level-chip-debug"]').should('be.visible')
    cy.get('[data-testid="logs-level-chip-fatal"]').should('be.visible')
  })

  it('should stream live logs', () => {
    cy.get('[data-testid="nav-admin-logs"]').click()

    // Wait for SSE connection and backfill to populate entries
    cy.get('[data-testid="logs-container"]', { timeout: 10000 })
      .find('[data-testid^="log-entry-"]')
      .should('have.length.greaterThan', 0)
  })

  it('should switch to history mode', () => {
    cy.get('[data-testid="nav-admin-logs"]').click()
    cy.get('[data-testid="logs-history-btn"]').click()

    // Pagination should appear in history mode
    cy.get('[data-testid="logs-pagination"]', { timeout: 10000 }).should('be.visible')
    cy.get('[data-testid="logs-date-picker"]').should('be.visible')
  })

  it('should navigate between all admin sections', () => {
    cy.get('[data-testid="nav-admin-dashboard"]').click()
    cy.url().should('include', '/dashboard')
    cy.contains('System Dashboard').should('be.visible')

    cy.get('[data-testid="nav-admin-logs"]').click()
    cy.url().should('include', '/logs')
    cy.contains('System Logs').should('be.visible')

    cy.get('[data-testid="nav-admin-blog"]').click()
    cy.url().should('include', '/blog')
    cy.contains('Blog Posts').should('be.visible')

    cy.get('[data-testid="nav-admin-about"]').click()
    cy.url().should('include', '/about')
    cy.contains('About Content').should('be.visible')

    cy.get('[data-testid="nav-admin-groovejr"]').click()
    cy.url().should('include', '/groovejr')
    cy.contains('GrooveJr Content').should('be.visible')

    cy.get('[data-testid="nav-admin-work"]').click()
    cy.url().should('include', '/work')
    cy.contains('Work Content').should('be.visible')
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
    cy.url().should('include', `/${section}`)
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

  it('should perform CRUD for Work content', () => {
    cy.get('[data-testid="nav-admin-work"]').click()
    performCrud('work', 'create-new-work-content', 'E2E Work', 'Work content test.')
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
    cy.url().should('include', '/blog')
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
