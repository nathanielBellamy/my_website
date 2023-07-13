describe('/', () => {
  beforeEach(() => {
    cy.visit('localhost:8080')
  }),

  it('loads nav bar with dropdowns', () => {
    const nav = cy.get('nav')
    
    nav.should('have.length', 1)
    nav.within(() => {
      cy.get('button').should('have.length', 2)


      // Site Section Dropdown
      const siteSectionDropdownButton = cy.get('#siteSectionDropdown')
      siteSectionDropdownButton.click()
      
      var siteSectionDropdown = cy.get('ul')
      siteSectionDropdown.should('have.length', 1)
      siteSectionDropdown = siteSectionDropdown.first()

      siteSectionDropdown
        .within(() => {
          const listItems = cy.get('li')
          listItems.should('have.length', 4)
          listItems.first().should('have.text', 'Home')
          listItems.next().should('have.text', 'About')
          listItems.next().should('have.text', 'Magic square')
          listItems.next().should('have.text', 'Give me a sine')
        })

      // close dropdown
      siteSectionDropdownButton.click()

      // Contact Info Dropdown
      const contactInfoDropdownButton = cy.get('#contactInfo')
      contactInfoDropdownButton.click()
      
      var contactInfoDropdown = cy.get('ul')
      contactInfoDropdown.should('have.length', 1)
      contactInfoDropdown = contactInfoDropdown.first()

      contactInfoDropdown
        .within(() => {
          const listItems = cy.get('li')
          listItems.should('have.length', 4)
          listItems.first().should('have.text', 'nbschieber@gmail.com')
          listItems.next().should('have.text', 'in/nateschieber')
          listItems.next().should('have.text', 'github.com/nathanielBellamy')
          listItems.next().should('have.text', 'PORTLAND, OR')
        })
      contactInfoDropdownButton.click()
    })
  })
})
