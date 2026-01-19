import { render, screen } from '@testing-library/angular';
import { CardComponent } from './card.component';

describe('CardComponent', () => {
  it('should render title and content', async () => {
    await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: 'Test Content',
      },
    });

    screen.getByText('Test Title');
    screen.getByText('Test Content');
  });

  it('should render tags when provided', async () => {
    await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: 'Test Content',
        tags: ['tag1', 'tag2'],
      },
    });

    screen.getByText('#tag1');
    screen.getByText('#tag2');
  });

  it('should not render tags section when tags are not provided', async () => {
    await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: 'Test Content',
      },
    });

    expect(screen.queryByText(/#/)).toBeNull();
  });

  it('should not render tags section when tags array is empty', async () => {
    await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: 'Test Content',
        tags: [],
      },
    });

    expect(screen.queryByText(/#/)).toBeNull();
  });
});

