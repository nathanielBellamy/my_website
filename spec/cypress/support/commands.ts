import { Mode } from './mode'

Cypress.Commands.add('epilepsy_warning', (accept: boolean) => {
  cy.get('[data-testid="epilepsy_warning"]')
    .contains('Epilepsy Warning')

  switch (accept) {
    case true:
      cy.get('[data-testid="epilepsy_warning_accept"]')
        .click()
      break

    case false:
      cy.get('[data-testid="epilepsy_warning_go_home"]')
        .click()
      cy.url().should('eq', 'http://localhost:8080/#/')
      break
  }
})

Cypress.Commands.add('visit_home', (mode: Mode) => {
  switch (mode) {
    case Mode.localhost:
      cy.visit('localhost:8080')
      cy.get('[data-testid="pw_input"]')
        .type('foo')
        .type('{enter}')
      break
    case Mode.remotedev:
      break
    default:
      cy.visit('https://nateschieber.dev/')
      break
  }
})

Cypress.Commands.add('wait_for_loading_screen', () => {
  cy.get('[data-testid="loading_title"]')
    .contains("Loading...")

  cy.get('[data-testid="loading"]').should('not.exist');
})
