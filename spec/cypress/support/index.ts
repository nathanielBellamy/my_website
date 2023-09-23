import { Mode } from './mode'

declare global {
  namespace Cypress {
    interface Chainable {
      /**
       * Custom command to select DOM element by data-cy attribute.
       * @example cy.dataCy('greeting')
       */
      visit_home(value: string): Chainable<null>
    }
  }
}
