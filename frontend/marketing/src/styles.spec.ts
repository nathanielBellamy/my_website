import { readFileSync } from 'node:fs';
import { join } from 'node:path';

describe('marketing markdown typography overrides', () => {
  it('should suppress injected backticks around inline code', () => {
    const styles = readFileSync(join(process.cwd(), 'src/styles.css'), 'utf8');

    expect(styles).toContain('@layer utilities');
    expect(styles).toContain('.prose :not(pre) > code::before');
    expect(styles).toContain('.prose :not(pre) > code::after');
    expect(styles).toContain('content: none !important;');
  });

  it('should suppress injected quotation marks inside blockquotes', () => {
    const styles = readFileSync(join(process.cwd(), 'src/styles.css'), 'utf8');

    expect(styles).toContain('.prose blockquote p:first-of-type::before');
    expect(styles).toContain('.prose blockquote p:last-of-type::after');
    expect(styles).toContain('content: none !important;');
  });
});
