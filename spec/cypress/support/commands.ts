// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
import { Mode } from './mode'

// -- This is a parent command --
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
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })
