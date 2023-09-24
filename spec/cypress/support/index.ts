import { Mode } from './mode'

declare global {
  namespace Cypress {
    interface Chainable {
      epilepsy_warning(accept: boolean): Chainable<null>
      visit_home(value: string): Chainable<null>
      wait_for_loading_screen(): Chainable<null>
    }
  }
}
