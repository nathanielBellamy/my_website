import { Mode } from "../support/mode"

describe('home', () => {
  it('says my name in the NavBar', () => {
    cy.visit_home(Mode.localhost)
    cy.get('[data-testid="nav_bar"]')
      .contains("Nate Schieber")
  })

  it('has my social links in the footer', () => {
    cy.visit_home(Mode.localhost)
    cy.get('[data-testid="footer"]')
      .within(() => {
        cy.get('[data-testid="github_social_icon"]')
          .should('have.attr', 'href', 'https://github.com/nathanielBellamy')
        cy.get('[data-testid="linkedin_social_icon"]')
          .should('have.attr', 'href', 'https://www.linkedin.com/in/nateschieber/')
        cy.get('[data-testid="mailto_social_icon"]')
          .should('have.attr', 'href', 'mailto:nbschieber@gmail.com')
      })
  })
})
