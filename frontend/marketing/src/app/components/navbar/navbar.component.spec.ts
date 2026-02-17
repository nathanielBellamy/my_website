import { render, screen, fireEvent } from '@testing-library/angular';
import { NavbarComponent } from './navbar.component';

describe('NavbarComponent', () => {
  it('should render the brand name', async () => {
    await render(NavbarComponent);
    screen.getByText('Nate');
  });

  it('should render the navigation links', async () => {
    await render(NavbarComponent);
    screen.getByTestId('nav-home');
    screen.getByTestId('nav-about');
    screen.getByTestId('nav-groovejr');
    screen.getByTestId('nav-blog');
  });

  it('should render social links', async () => {
    await render(NavbarComponent);
    screen.getByTestId('navbar-mailto');
    screen.getByTestId('navbar-linked-in');
    screen.getByTestId('navbar-github');
  });

  it('should toggle mobile menu', async () => {
    await render(NavbarComponent);
    
    // Initially mobile menu should not be present
    expect(screen.queryByTestId('mobile-menu')).toBeNull();
    
    // Click toggle button
    const toggleButton = screen.getByTestId('navbar-toggle');
    fireEvent.click(toggleButton);
    
    // Mobile menu should be present
    expect(screen.getByTestId('mobile-menu')).toBeTruthy();
    
    // Check for mobile links
    screen.getByTestId('mobile-nav-home');
    screen.getByTestId('mobile-nav-about');
    screen.getByTestId('mobile-nav-groovejr');
    screen.getByTestId('mobile-nav-blog');
    screen.getByTestId('mobile-nav-old-site');
    
    // Click a link to close the menu
    const homeLink = screen.getByTestId('mobile-nav-home');
    fireEvent.click(homeLink);
    
    // Mobile menu should be closed (removed from DOM)
    expect(screen.queryByTestId('mobile-menu')).toBeNull();
  });
});