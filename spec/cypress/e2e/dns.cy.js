import { Mode } from "../support/mode"

describe('all domains', () => {
  it('redirect to nateschieber.dev', () => {
    cy.visit_home(Mode.prod)
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
