describe('all domains', () => {
  it('redirect to nateschieber.dev', () => {
    cy.visit_home()
    cy.url().should('eq', 'https://nateschieber.dev/')

    cy.visit('https://nateschieber.com')
    cy.url().should('eq', 'https://nateschieber.dev/')

    cy.visit('https://nateschieber.net')
    cy.url().should('eq', 'https://nateschieber.dev/')

    cy.visit('https://nateschieber.org')
    cy.url().should('eq', 'https://nateschieber.dev/')

    cy.visit('https://nathanschieber.com')
    cy.url().should('eq', 'https://nateschieber.dev/')
  })
})
