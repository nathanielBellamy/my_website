import { render, screen } from '@testing-library/angular';
import { provideMarkdown } from 'ngx-markdown';
import { CardComponent } from './card.component';

describe('CardComponent', () => {
  it('should render the card with title and content', async () => {
    await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: 'Test Content',
      },
      providers: [provideMarkdown()],
    });

    expect(screen.getByText('Test Title')).toBeTruthy();
    expect(screen.getByText('Test Content')).toBeTruthy();
  });

  it('should render tags if provided', async () => {
    await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: 'Test Content',
        tags: ['tag1', 'tag2'],
      },
      providers: [provideMarkdown()],
    });

    expect(screen.getByText('#tag1')).toBeTruthy();
    expect(screen.getByText('#tag2')).toBeTruthy();
  });

  it('should not render tags if not provided', async () => {
    await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: 'Test Content',
      },
      providers: [provideMarkdown()],
    });

    expect(screen.queryByText(/#/)).toBeNull();
  });
});
