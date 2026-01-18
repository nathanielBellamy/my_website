import { render, screen } from '@testing-library/angular';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { HeaderComponent } from './components/header/header.component';

describe('AppComponent', () => {
  it('should render the navbar', async () => {
    await render(AppComponent, {
      imports: [NavbarComponent, HeaderComponent],
    });
    expect(screen.getByRole('navigation')).toBeTruthy();
  });

  it('should render the header', async () => {
    await render(AppComponent, {
      imports: [NavbarComponent, HeaderComponent],
    });
    expect(screen.getByRole('banner')).toBeTruthy();
  });
});