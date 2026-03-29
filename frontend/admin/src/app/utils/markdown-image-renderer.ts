import { MarkedExtension } from 'marked';

/**
 * Custom marked extension that parses image size from alt text.
 * Syntax: ![alt text|WIDTHxHEIGHT](url)
 * Examples:
 *   ![photo|400x300](/v1/api/images/pic.jpg)  → 400×300
 *   ![photo](/v1/api/images/pic.jpg)          → natural size
 */
export const markedImageResizeExtension: MarkedExtension = {
  renderer: {
    image({ href, title, text }: { href: string; title: string | null; text: string }) {
      let alt = text;
      let width: string | null = null;
      let height: string | null = null;

      const sizeMatch = alt.match(/^(.*?)\|(\d+)x(\d+)$/);
      if (sizeMatch) {
        alt = sizeMatch[1].trim();
        width = sizeMatch[2];
        height = sizeMatch[3];
      }

      const attrs: string[] = [
        `src="${href}"`,
        `alt="${alt}"`,
        'loading="lazy"',
        'style="max-width:100%;height:auto"',
      ];

      if (title) {
        attrs.push(`title="${title}"`);
      }
      if (width) {
        attrs.push(`width="${width}"`);
      }
      if (height) {
        attrs.push(`height="${height}"`);
      }

      return `<img ${attrs.join(' ')} />`;
    },
  },
};
