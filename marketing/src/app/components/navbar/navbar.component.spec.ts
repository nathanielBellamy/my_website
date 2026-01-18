import { render, screen } from '@testing-library/angular';
import { NavbarComponent } from './navbar.component';

describe('NavbarComponent', () => {
  it('should render the brand name', async () => {
    await render(NavbarComponent);
    expect(screen.getByText('Nate')).toBeTruthy();
  });

  it('should render the navigation links', async () => {
    await render(NavbarComponent);
    expect(screen.getByText('Home')).toBeTruthy();
    expect(screen.getByText('About')).toBeTruthy();
    expect(screen.getByText('GrooveJr')).toBeTruthy();
    expect(screen.getByText('Blog')).toBeTruthy();
  });

  it('should render social links', async () => {
    await render(NavbarComponent);
    expect(screen.getByTestId('navbar-mailto')).toBeTruthy();
    expect(screen.getByTestId('navbar-linked-in')).toBeTruthy();
    expect(screen.getByTestId('navbar-github')).toBeTruthy();
  });
});