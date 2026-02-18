describe('Navbar Sticky Behavior', () => {
  it('should remain visible at the top when scrolling down', () => {
    cy.visit('/');

    // Ensure the navbar is visible initially
    cy.get('app-navbar').should('be.visible');

    // Force the scroll container to be scrollable by adding a tall element inside it
    cy.get('[data-testid="main-scroll-container"]').then(($container) => {
      const div = document.createElement('div');
      div.style.height = '2000px';
      div.style.width = '100%';
      // Append to the container, not the body
      $container[0].appendChild(div);
    });

    // Wait for layout
    cy.wait(100);

    // Scroll the container, not the window
    cy.get('[data-testid="main-scroll-container"]').scrollTo(0, 500);

    // Ensure the navbar is still visible
    cy.get('app-navbar').scrollIntoView().should('be.visible');

    // Verify the container is still scrolled down (proving the navbar stuck to the top of the viewport while scrolled)
    cy.get('[data-testid="main-scroll-container"]').invoke('scrollTop').should('be.gt', 0);

    // Check if it's at the top of the viewport
    cy.get('app-navbar').then(($nav) => {
      const rect = $nav[0].getBoundingClientRect();
      // It should stick to the top of the viewport (top: 0 relative to viewport)
      expect(rect.top).to.be.closeTo(0, 1); 
    });
  });
});
