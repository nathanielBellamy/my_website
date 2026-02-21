import { render, screen } from '@testing-library/angular';
import { FocusComponent } from './focus.component';
import { provideRouter } from '@angular/router';

describe('FocusComponent', () => {
  it('should render and contain scroll indicator', async () => {
    await render(FocusComponent, {
      providers: [provideRouter([])]
    });
    screen.getByTestId('featured-values-header');
    screen.getByTestId('scroll-to-latest-posts');
  });
});
