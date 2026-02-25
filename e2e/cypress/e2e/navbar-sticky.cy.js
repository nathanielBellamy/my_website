describe('Navbar Sticky Behavior', () => {
  it('should remain visible at the top when scrolling down', () => {
    cy.visit('/');

    // Ensure the navbar is visible initially
    cy.get('app-navbar').should('be.visible');

    // Wait for layout
    cy.wait(100);

    // Navigate
    cy.get('[data-testid="nav-blog"]').click()

    // Ensure the navbar is still visible
    cy.get('app-navbar').should('be.visible');

    // Check if it's at the top of the viewport
    cy.get('app-navbar').then(($nav) => {
      const rect = $nav[0].getBoundingClientRect();
      // It should stick to the top of the viewport (top: 0 relative to viewport)
      expect(rect.top).to.be.closeTo(0, 1); 
    });
  });
});
