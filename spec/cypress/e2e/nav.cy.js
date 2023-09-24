import { Mode } from "../support/mode"

describe('nav', () => {
  beforeEach(() => {
    cy.visit_home(Mode.localhost)
  })

  it('navbar displays my name', () => {
    cy.get('[data-testid="nav_bar"]')
      .contains("Nate Schieber")
  })

  it('goes to About', () => {
    cy.get('[data-testid="nav_button"]')
      .click()

    cy.get('[data-testid="nav_dropdown"]')
      .within(() => {
        cy.get('[data-testid="nav_dropdown_about"]')
          .click()
      })

    cy.url().should('eq', 'http://localhost:8080/#/about')
    
    cy.get('[data-testid="about_personal_projects"]')
      .contains('Personal projects')
    
    cy.get('[data-testid="about_technical_experience"]')
      .contains('Technical experience')
  })

  it('goes to Give Me A Sine', () => {
    cy.get('[data-testid="nav_button"]')
      .click()

    cy.get('[data-testid="nav_dropdown"]')
      .within(() => {
        cy.get('[data-testid="nav_dropdown_give_me_a_sine"]')
          .click()
      })

    cy.url().should('eq', 'http://localhost:8080/#/give-me-a-sine')

    cy.get('[data-testid="gmas_form_header"]')
      .contains("f(x) = a * sin(b*x + c)")
  })


  it('goes to System Diagram', () => {
    cy.get('[data-testid="nav_button"]')
      .click()

    cy.get('[data-testid="nav_dropdown"]')
      .within(() => {
        cy.get('[data-testid="nav_dropdown_system_diagram"]')
          .click()
      })

    cy.url().should('eq', 'http://localhost:8080/#/system-diagram')

    cy.get('[data-testid="system_diagram"]')
      .should('have.attr', 'alt', 'System Diagram')
  })
  
  // TODO: handle recaptcha in tests
  it('goes to Public Square', () => {
    cy.get('[data-testid="nav_button"]')
      .click()

    cy.get('[data-testid="nav_dropdown"]')
      .within(() => {
        cy.get('[data-testid="nav_dropdown_public_square"]')
          .click()
      })

    cy.url().should('eq', 'http://localhost:8080/#/public-square')

    cy.wait_for_loading_screen()
    cy.get('[data-testid="public_square_info_gate_welcome"]')
      .contains('Welcome to the')
    cy.get('[data-testid="public_square_info_gate_title"]')
      .contains('Public Square')
  })

  describe('goes to Magic Square', () => {
    describe('when user accepts Epiliepsy warning', () => {
      it('continues to Magic Square', () => {
        cy.get('[data-testid="nav_button"]')
          .click()

        cy.get('[data-testid="nav_dropdown"]')
          .within(() => {
            cy.get('[data-testid="nav_dropdown_magic_square"]')
              .click()
          })

        cy.url().should('eq', 'http://localhost:8080/#/magic-square')

        cy.wait_for_loading_screen()
        cy.epilepsy_warning(true)
        cy.get('[data-testid="magic_square"]')
          .should('have.attr', 'id', 'magic_square')
      })
    })


    describe('when user rejects Epiliepsy Warning', () => {
      it('goes to home', () => {
        cy.get('[data-testid="nav_button"]')
          .click()

        cy.get('[data-testid="nav_dropdown"]')
          .within(() => {
            cy.get('[data-testid="nav_dropdown_magic_square"]')
              .click()
          })

        cy.url().should('eq', 'http://localhost:8080/#/magic-square')

        cy.wait_for_loading_screen()
        cy.epilepsy_warning(false) // this command checks after navigation home URL
      })
    })
  })
})
