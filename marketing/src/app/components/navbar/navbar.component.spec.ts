import { render, screen } from '@testing-library/angular';
import { NavbarComponent } from './navbar.component';

describe('NavbarComponent', () => {
  it('should render the brand name', async () => {
    await render(NavbarComponent);
    screen.getByText('Nate');
  });

  it('should render the navigation links', async () => {
    await render(NavbarComponent);
    screen.getByText('Home');
    screen.getByText('About');
    screen.getByText('GrooveJr');
    screen.getByText('Blog');
  });

  it('should render social links', async () => {
    await render(NavbarComponent);
    screen.getByTestId('navbar-mailto');
    screen.getByTestId('navbar-linked-in');
    screen.getByTestId('navbar-github');
  });
});