import { render, screen } from '@testing-library/angular';
import { GrooveJrComponent } from './groove-jr.component';

describe('GrooveJrComponent', () => {
  it('should render groove-jr works!', async () => {
    await render(GrooveJrComponent);
    expect(screen.getByText('groove-jr works!')).toBeTruthy();
  });
});