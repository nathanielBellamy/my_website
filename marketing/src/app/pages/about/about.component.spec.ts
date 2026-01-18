import { render, screen } from '@testing-library/angular';
import { AboutComponent } from './about.component';

describe('AboutComponent', () => {
  it('should render about works!', async () => {
    await render(AboutComponent);
    expect(screen.getByText('about works!')).toBeTruthy();
  });
});