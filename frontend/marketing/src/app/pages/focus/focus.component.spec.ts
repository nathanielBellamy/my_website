import { render, screen } from '@testing-library/angular';
import { FocusComponent } from './focus.component';
import { provideRouter } from '@angular/router';

describe('FocusComponent', () => {
  it('should render ', async () => {
    await render(FocusComponent, {
      providers: [provideRouter([])]
    });
    screen.getByTestId('focus-core-values');
    screen.getByTestId('focus-tech');
  });
});
