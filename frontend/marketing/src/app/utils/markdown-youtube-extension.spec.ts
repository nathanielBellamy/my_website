import { extractYouTubeVideoId, markedYouTubeExtension } from './markdown-youtube-extension';

describe('extractYouTubeVideoId', () => {
  it('extracts video id from a standard watch URL', () => {
    expect(extractYouTubeVideoId('https://www.youtube.com/watch?v=dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
  });

  it('extracts video id from a watch URL with extra query params', () => {
    expect(extractYouTubeVideoId('https://www.youtube.com/watch?list=PL1&v=dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
  });

  it('extracts video id from a youtu.be short URL', () => {
    expect(extractYouTubeVideoId('https://youtu.be/dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
  });

  it('extracts video id from an embed URL', () => {
    expect(extractYouTubeVideoId('https://www.youtube.com/embed/dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
  });

  it('returns null for a non-YouTube URL', () => {
    expect(extractYouTubeVideoId('https://example.com/video')).toBeNull();
  });

  it('returns null for an empty string', () => {
    expect(extractYouTubeVideoId('')).toBeNull();
  });
});

describe('markedYouTubeExtension link renderer', () => {
  const renderLink = markedYouTubeExtension.renderer!['link' as keyof typeof markedYouTubeExtension.renderer] as Function;

  it('renders an iFrame for a YouTube watch URL', () => {
    const result = renderLink({ href: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ', title: null, text: 'My Video' });
    expect(result).toContain('<iframe');
    expect(result).toContain('https://www.youtube.com/embed/dQw4w9WgXcQ');
    expect(result).toContain('allowfullscreen');
    expect(result).toContain('loading="lazy"');
  });

  it('uses the link title as the iframe title attribute when provided', () => {
    const result = renderLink({ href: 'https://youtu.be/dQw4w9WgXcQ', title: 'Custom Title', text: 'Video' });
    expect(result).toContain('title="Custom Title"');
  });

  it('falls back to link text as the iframe title when no title is provided', () => {
    const result = renderLink({ href: 'https://youtu.be/dQw4w9WgXcQ', title: null, text: 'Video Text' });
    expect(result).toContain('title="Video Text"');
  });

  it('falls back to default title when neither title nor text is provided', () => {
    const result = renderLink({ href: 'https://youtu.be/dQw4w9WgXcQ', title: null, text: '' });
    expect(result).toContain('title="YouTube video"');
  });

  it('returns false for a non-YouTube link so default rendering is used', () => {
    const result = renderLink({ href: 'https://example.com', title: null, text: 'Example' });
    expect(result).toBe(false);
  });

  it('wraps iFrame in a responsive container div', () => {
    const result = renderLink({ href: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ', title: null, text: 'Video' });
    expect(result).toContain('class="youtube-embed"');
    expect(result).toContain('padding-bottom:56.25%');
  });
});
