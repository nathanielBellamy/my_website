/**
 * Strips markdown media/link syntax from raw content so that
 * snippets display only plain text.
 *
 * Removes:
 *  - images:  ![alt](url)
 *  - videos / iframes embedded via markdown image syntax
 *  - inline links: [text](url) → keeps "text"
 *  - reference links: [text][ref]
 *  - bare URLs: http(s)://...
 *  - remaining markdown formatting: #, *, `, ~
 */
export function stripMarkdownMedia(raw: string): string {
  return raw
    // Remove images / videos: ![alt](url "title")
    .replace(/!\[[^\]]*\]\([^)]*\)/g, '')
    // Convert inline links to just their text: [text](url) → text
    .replace(/\[([^\]]*)\]\([^)]*\)/g, '$1')
    // Remove reference-style links: [text][ref]
    .replace(/\[([^\]]*)\]\[[^\]]*\]/g, '$1')
    // Remove bare URLs
    .replace(/https?:\/\/\S+/g, '')
    // Strip common markdown formatting chars
    .replace(/[#*`~>]/g, '')
    // Collapse multiple spaces / newlines into a single space
    .replace(/\s+/g, ' ')
    .trim();
}

export function getSnippet(content: string): string {
  if (!content) return '';
  const clean = stripMarkdownMedia(content);
  const firstPeriod = clean.indexOf('.');
  if (firstPeriod > -1 && firstPeriod < 200) {
    return clean.substring(0, firstPeriod + 1);
  }
  return clean.length > 150 ? clean.substring(0, 150) + '...' : clean;
}
