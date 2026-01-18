import { render, screen } from '@testing-library/angular';
import { BlogComponent } from './blog.component';

describe('BlogComponent', () => {
  it('should render blog works!', async () => {
    await render(BlogComponent);
    expect(screen.getByText('blog works!')).toBeTruthy();
  });
});