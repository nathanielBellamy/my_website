import { render, screen } from '@testing-library/angular';
import { HeaderComponent } from './header.component';

describe('HeaderComponent', () => {
  it('should render the header text', async () => {
    await render(HeaderComponent);
    expect(screen.getByText('Hi, my name is Nate.')).toBeTruthy();
  });
});