import { encodeId, decodeId } from './id-encoder';

describe('encodeId', () => {
  it('removes dashes from a standard UUID', () => {
    expect(encodeId('550e8400-e29b-41d4-a716-446655440000')).toBe('550e8400e29b41d4a716446655440000');
  });

  it('returns a string with no dashes unchanged', () => {
    expect(encodeId('550e8400e29b41d4a716446655440000')).toBe('550e8400e29b41d4a716446655440000');
  });

  it('returns a short non-UUID string unchanged', () => {
    expect(encodeId('123')).toBe('123');
  });

  it('removes all dashes from a string with multiple dashes', () => {
    expect(encodeId('a-b-c-d')).toBe('abcd');
  });
});

describe('decodeId', () => {
  it('adds dashes to a 32-char hex string to form a UUID', () => {
    expect(decodeId('550e8400e29b41d4a716446655440000')).toBe('550e8400-e29b-41d4-a716-446655440000');
  });

  it('returns an already-formatted UUID unchanged', () => {
    expect(decodeId('550e8400-e29b-41d4-a716-446655440000')).toBe('550e8400-e29b-41d4-a716-446655440000');
  });

  it('returns a short non-UUID string unchanged', () => {
    expect(decodeId('123')).toBe('123');
  });

  it('returns a 31-char hex string unchanged (too short for UUID)', () => {
    expect(decodeId('550e8400e29b41d4a71644665544000')).toBe('550e8400e29b41d4a71644665544000');
  });

  it('returns a 33-char hex string unchanged (too long for UUID)', () => {
    expect(decodeId('550e8400e29b41d4a7164466554400001')).toBe('550e8400e29b41d4a7164466554400001');
  });

  it('encodeId and decodeId are inverse operations for a standard UUID', () => {
    const uuid = '550e8400-e29b-41d4-a716-446655440000';
    expect(decodeId(encodeId(uuid))).toBe(uuid);
  });

  it('handles uppercase hex chars', () => {
    expect(decodeId('550E8400E29B41D4A716446655440000')).toBe('550E8400-E29B-41D4-A716-446655440000');
  });
});
