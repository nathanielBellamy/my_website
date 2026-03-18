import { getSnippet, stripMarkdownMedia } from './snippet';

describe('stripMarkdownMedia', () => {
  it('should remove markdown images', () => {
    const input = 'Hello ![photo](https://example.com/img.png) world';
    expect(stripMarkdownMedia(input)).toBe('Hello world');
  });

  it('should remove images with alt text and title', () => {
    const input = 'Check ![my pic](https://example.com/img.png "title") out';
    expect(stripMarkdownMedia(input)).toBe('Check out');
  });

  it('should convert inline links to plain text', () => {
    const input = 'Visit [my site](https://example.com) today';
    expect(stripMarkdownMedia(input)).toBe('Visit my site today');
  });

  it('should remove bare URLs', () => {
    const input = 'Go to https://example.com/video for more';
    expect(stripMarkdownMedia(input)).toBe('Go to for more');
  });

  it('should strip markdown formatting characters', () => {
    const input = '## Hello **world** `code` ~~strikethrough~~';
    expect(stripMarkdownMedia(input)).toBe('Hello world code strikethrough');
  });

  it('should handle embedded video markdown', () => {
    const input = 'Watch this: ![Video](https://youtube.com/watch?v=abc123) and enjoy';
    expect(stripMarkdownMedia(input)).toBe('Watch this: and enjoy');
  });

  it('should collapse whitespace', () => {
    const input = 'Hello   \n\n  world';
    expect(stripMarkdownMedia(input)).toBe('Hello world');
  });

  it('should handle empty string', () => {
    expect(stripMarkdownMedia('')).toBe('');
  });

  it('should remove reference-style links', () => {
    const input = 'See [this article][1] for details';
    expect(stripMarkdownMedia(input)).toBe('See this article for details');
  });
});

describe('getSnippet', () => {
  it('should return empty string for empty content', () => {
    expect(getSnippet('')).toBe('');
  });

  it('should return first sentence when period is within 200 chars', () => {
    expect(getSnippet('This is the first sentence. And more text here.'))
      .toBe('This is the first sentence.');
  });

  it('should truncate to 150 chars with ellipsis when no early period', () => {
    const long = 'A'.repeat(200);
    expect(getSnippet(long)).toBe('A'.repeat(150) + '...');
  });

  it('should return full content if shorter than 150 chars with no period', () => {
    expect(getSnippet('Short content')).toBe('Short content');
  });

  it('should strip markdown images before generating snippet', () => {
    const input = '![banner](https://example.com/banner.jpg) Welcome to our blog post. More content follows.';
    expect(getSnippet(input)).toBe('Welcome to our blog post.');
  });

  it('should strip inline links but keep link text in snippet', () => {
    const input = 'Check out [my project](https://github.com/example). It does cool things.';
    expect(getSnippet(input)).toBe('Check out my project.');
  });

  it('should handle content that is only a media embed', () => {
    const input = '![video](https://youtube.com/watch?v=abc123)';
    expect(getSnippet(input)).toBe('');
  });
});
