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

  it('should render inline code without markdown backticks in the DOM', async () => {
    const { container } = await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: "Use `const foo = 'bar'` here.",
      },
      providers: [provideMarkdown()],
    });

    const codeElement = container.querySelector('code');

    expect(codeElement).not.toBeNull();
    expect(codeElement?.textContent).toBe("const foo = 'bar'");
    expect(container.textContent).not.toContain("`const foo = 'bar'`");
  });

  it('should render blockquotes without injecting quotation marks into the DOM', async () => {
    const { container } = await render(CardComponent, {
      componentInputs: {
        title: 'Test Title',
        content: '> foo',
      },
      providers: [provideMarkdown()],
    });

    const blockquoteElement = container.querySelector('blockquote');

    expect(blockquoteElement).not.toBeNull();
    expect(blockquoteElement?.textContent?.trim()).toBe('foo');
    expect(container.textContent).not.toContain('"foo"');
  });
});
