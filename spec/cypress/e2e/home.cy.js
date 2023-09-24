import { Mode } from "../support/mode"
import { version_current, version_latest_major } from '../../../version'

let version_url_base = "https://github.com/nathanielBellamy/my_website/releases/tag/"

describe('home', () => {
  beforeEach(() => {
    cy.visit_home(Mode.localhost)
  })

  describe('version number', () => {
    it('current is displayed', () => {
      cy.get('[data-testid="version_current"]')
        .contains(version_current)

      cy.get('[data-testid="version_link_current"]')
        .should('have.attr', 'href', `${version_url_base}${version_current}`)
    })

    it('latest_major is displayed', () => {
      cy.get('[data-testid="version_latest_major"]')
        .contains(version_latest_major)
      
      cy.get('[data-testid="version_link_latest_major"]')
        .should('have.attr', 'href', `${version_url_base}${version_latest_major}`)
    })
  })

  it('displays whats here', () => {
    cy.get('[data-testid="whats_here"]')
      .contains('What\'s here?')
  })
})
