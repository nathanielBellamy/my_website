import { render, screen } from '@testing-library/angular';
import { HomeComponent } from './home.component';
import { provideRouter } from '@angular/router';

describe('HomeComponent', () => {
  it('should render', async () => {
    await render(HomeComponent, {
      providers: [provideRouter([])]
    });
    screen.getByTestId('hero-heading');
    screen.getByTestId('scroll-to-focus');
  });
});
