import { MarkedExtension } from 'marked';

/**
 * Custom marked extension that converts YouTube links to embedded iFrames.
 * Syntax: [video title](https://www.youtube.com/watch?v=VIDEO_ID)
 * Or:     [video title](https://youtu.be/VIDEO_ID)
 * Examples:
 *   [My Video](https://www.youtube.com/watch?v=dQw4w9WgXcQ)  → embedded iFrame
 *   [My Video](https://youtu.be/dQw4w9WgXcQ)                 → embedded iFrame
 */
export function extractYouTubeVideoId(href: string): string | null {
  const watchMatch = href.match(/(?:youtube\.com\/watch\?(?:[^#&]*&)*v=)([a-zA-Z0-9_-]{11})/);
  if (watchMatch) return watchMatch[1];

  const shortMatch = href.match(/youtu\.be\/([a-zA-Z0-9_-]{11})/);
  if (shortMatch) return shortMatch[1];

  const embedMatch = href.match(/youtube\.com\/embed\/([a-zA-Z0-9_-]{11})/);
  if (embedMatch) return embedMatch[1];

  return null;
}

export const markedYouTubeExtension: MarkedExtension = {
  renderer: {
    link({ href, title, text }: { href: string; title: string | null; text: string }) {
      const videoId = extractYouTubeVideoId(href);
      if (!videoId) {
        return false;
      }

      const embedUrl = `https://www.youtube.com/embed/${videoId}`;
      const titleAttr = title || text || 'YouTube video';

      return [
        `<div`,
        `  class="youtube-embed"`,
        `  style="position:relative;padding-bottom:56.25%;height:0;overflow:hidden;max-width:100%"`,
        `>`,
        `  <iframe`,
        `    src="${embedUrl}"`,
        `    title="${titleAttr}"`,
        `    style="position:absolute;top:0;left:0;width:100%;height:100%;border:0"`,
        `    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"`,
        `    allowfullscreen`,
        `    loading="lazy"`,
        `  ></iframe>`,
        `</div>`,
      ].join('\n');
    },
  },
};
