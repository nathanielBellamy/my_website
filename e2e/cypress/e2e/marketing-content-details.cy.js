describe('Marketing Content Details', () => {
  describe('About Content Details', () => {
    beforeEach(() => {
      cy.intercept('GET', '**/marketing/about?*', {
        statusCode: 200,
        body: [
          {
            id: '123',
            title: 'Test About Title',
            content: 'This is test about content. It describes the author.',
            order: 1,
          },
        ],
      }).as('getAboutContent');

      cy.intercept('GET', '**/marketing/about/123', {
        statusCode: 200,
        body: {
          id: '123',
          title: 'Test About Title',
          content: '# Test About Title\n\nThis is test about content. It describes the author.',
          order: 1,
        },
      }).as('getAboutContentById');

      cy.visit('/about');
    });

    it('should display about content cards and navigate to details page', () => {
      cy.wait('@getAboutContent');

      cy.get('[data-testid="about-header"]').scrollIntoView().should('be.visible');

      cy.contains('Test About Title', { timeout: 10000 }).should('be.visible');
      cy.contains('This is test about content.').should('be.visible');

      cy.contains('Test About Title').click();

      cy.url().should('include', '/about/123');

      cy.wait('@getAboutContentById');

      cy.contains('Back to About', { timeout: 10000 }).scrollIntoView().should('be.visible');
      cy.get('[data-testid="about-content-details-title"]')
        .contains('Test About Title')
        .should('be.visible');
      cy.contains('This is test about content.').should('be.visible');

      cy.contains('Back to About').click();

      cy.get('[data-testid="about-header"]').scrollIntoView().should('be.visible');
      cy.url().should('include', '/about');
    });

    it('should navigate directly to about content details page by URL', () => {
      cy.visit('/about/123');

      cy.wait('@getAboutContentById');

      cy.get('[data-testid="about-content-details-title"]', { timeout: 10000 })
        .contains('Test About Title')
        .should('be.visible');
      cy.contains('This is test about content.').should('be.visible');
      cy.contains('Back to About').should('be.visible');
    });
  });

  describe('GrooveJr Content Details', () => {
    beforeEach(() => {
      cy.intercept('GET', '**/marketing/groovejr?*', {
        statusCode: 200,
        body: [
          {
            id: '456',
            title: 'Test GrooveJr Title',
            content: 'This is test GrooveJr content. It describes the music player.',
            order: 1,
          },
        ],
      }).as('getGrooveJrContent');

      cy.intercept('GET', '**/marketing/groovejr/456', {
        statusCode: 200,
        body: {
          id: '456',
          title: 'Test GrooveJr Title',
          content: '# Test GrooveJr Title\n\nThis is test GrooveJr content. It describes the music player.',
          order: 1,
        },
      }).as('getGrooveJrContentById');

      cy.visit('/groovejr');
    });

    it('should display groovejr content cards and navigate to details page', () => {
      cy.wait('@getGrooveJrContent');

      cy.get('[data-testid="groovejr-header"]').first().scrollIntoView().should('be.visible');

      cy.contains('Test GrooveJr Title', { timeout: 10000 }).should('be.visible');
      cy.contains('This is test GrooveJr content.').should('be.visible');

      cy.contains('Test GrooveJr Title').click();

      cy.url().should('include', '/groovejr/456');

      cy.wait('@getGrooveJrContentById');

      cy.contains('Back to GrooveJr', { timeout: 10000 }).scrollIntoView().should('be.visible');
      cy.get('[data-testid="groove-jr-content-details-title"]')
        .contains('Test GrooveJr Title')
        .should('be.visible');
      cy.contains('This is test GrooveJr content.').should('be.visible');

      cy.contains('Back to GrooveJr').click();

      cy.get('[data-testid="groovejr-header"]').first().scrollIntoView().should('be.visible');
      cy.url().should('include', '/groovejr');
    });

    it('should navigate directly to groovejr content details page by URL', () => {
      cy.visit('/groovejr/456');

      cy.wait('@getGrooveJrContentById');

      cy.get('[data-testid="groove-jr-content-details-title"]', { timeout: 10000 })
        .contains('Test GrooveJr Title')
        .should('be.visible');
      cy.contains('This is test GrooveJr content.').should('be.visible');
      cy.contains('Back to GrooveJr').should('be.visible');
    });
  });

  describe('Home Content Details', () => {
    beforeEach(() => {
      cy.intercept('GET', '**/marketing/home?*', {
        statusCode: 200,
        body: [
          {
            id: '789',
            title: 'Test Home Title',
            content: 'This is test home content. It describes the latest update.',
            order: 1,
          },
        ],
      }).as('getHomeContent');

      cy.intercept('GET', '**/marketing/home/789', {
        statusCode: 200,
        body: {
          id: '789',
          title: 'Test Home Title',
          content: '# Test Home Title\n\nThis is test home content. It describes the latest update.',
          order: 1,
        },
      }).as('getHomeContentById');

      cy.visit('/latest-posts');
    });

    it('should display home content cards and navigate to details page', () => {
      cy.wait('@getHomeContent');

      cy.contains('Latest Posts', { timeout: 10000 }).should('be.visible');

      cy.contains('Test Home Title', { timeout: 10000 }).should('be.visible');
      cy.contains('This is test home content.').should('be.visible');

      cy.contains('Test Home Title').click();

      cy.url().should('include', '/home/789');

      cy.wait('@getHomeContentById');

      cy.contains('Back to Latest Posts', { timeout: 10000 }).scrollIntoView().should('be.visible');
      cy.get('[data-testid="home-content-details-title"]')
        .contains('Test Home Title')
        .should('be.visible');
      cy.contains('This is test home content.').should('be.visible');

      cy.contains('Back to Latest Posts').click();

      cy.contains('Latest Posts').should('be.visible');
      cy.url().should('include', '/latest-posts');
    });

    it('should navigate directly to home content details page by URL', () => {
      cy.visit('/home/789');

      cy.wait('@getHomeContentById');

      cy.get('[data-testid="home-content-details-title"]', { timeout: 10000 })
        .contains('Test Home Title')
        .should('be.visible');
      cy.contains('This is test home content.').should('be.visible');
      cy.contains('Back to Latest Posts').should('be.visible');
    });
  });
});
