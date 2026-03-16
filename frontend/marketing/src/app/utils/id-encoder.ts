/**
 * Encodes a UUID to a hex string by removing dashes for use in URLs.
 * This obscures the raw UUID format while keeping the ID URL-friendly.
 *
 * Example: '550e8400-e29b-41d4-a716-446655440000' → '550e8400e29b41d4a716446655440000'
 */
export function encodeId(id: string): string {
  return id.replace(/-/g, '');
}

/**
 * Decodes a hex-encoded UUID (32 lowercase hex chars, no dashes) back to
 * standard UUID format. If the input is already UUID-formatted, a
 * non-UUID string, or any other value, it is returned unchanged.
 *
 * Example: '550e8400e29b41d4a716446655440000' → '550e8400-e29b-41d4-a716-446655440000'
 */
export function decodeId(id: string): string {
  if (/^[0-9a-f]{32}$/i.test(id)) {
    return `${id.slice(0, 8)}-${id.slice(8, 12)}-${id.slice(12, 16)}-${id.slice(16, 20)}-${id.slice(20)}`;
  }
  return id;
}
