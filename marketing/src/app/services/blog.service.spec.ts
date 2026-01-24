import { BlogService } from './blog.service';
import { BlogPost } from '../models/blog.model';
import { environment } from '../../environments/environment.localhost';
import { PaginatedResponse } from '../models/pagination.model';

const mockBlogPosts: BlogPost[] = [
  { id: 1, title: 'Post 1', content: 'Content 1', author: 'Author 1', date: '2026-01-18' },
  { id: 2, title: 'Post 2', content: 'Content 2', author: 'Author 2', date: '2026-01-18' },
];

const mockPaginatedResponse: PaginatedResponse<BlogPost> = {
  data: mockBlogPosts,
  total: 2,
  page: 1,
  limit: 10,
};

describe('BlogService', () => {
  let service: BlogService;
  const API_URL = `${environment.API_BASE_URL}/blog`;

  beforeEach(() => {
    service = new BlogService();
    global.fetch = jest.fn();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  describe('getAll', () => {
    it('should return paginated blog posts', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({
        ok: true,
        json: () => Promise.resolve(mockPaginatedResponse),
      });

      const posts = await service.getAll(1, 10);
      expect(posts).toEqual(mockPaginatedResponse);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}?page=1&limit=10`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getAll(1, 10)).rejects.toThrow('Failed to fetch blog posts');
    });
  });

  describe('getById', () => {
    it('should return a single blog post', async () => {
      const post = mockBlogPosts[0];
      (global.fetch as jest.Mock).mockResolvedValue({
        ok: true,
        json: () => Promise.resolve(post),
      });

      const result = await service.getById('blog-1');
      expect(result).toEqual(post);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}/1`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getById('blog-uuid')).rejects.toThrow('Failed to fetch blog post.');
    });
  });

  describe('getByTag', () => {
    it('should return paginated blog posts for a tag', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({
        ok: true,
        json: () => Promise.resolve(mockPaginatedResponse),
      });

      const posts = await service.getByTag('testing');
      expect(posts).toEqual(mockPaginatedResponse);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}/tag/testing?page=1&limit=10`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getByTag('testing')).rejects.toThrow('Failed to fetch blog posts with tag testing');
    });
  });

  describe('getByDate', () => {
    it('should return paginated blog posts for a date', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({
        ok: true,
        json: () => Promise.resolve(mockPaginatedResponse),
      });

      const posts = await service.getByDate(new Date(), new Date());
      expect(posts).toEqual(mockPaginatedResponse);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}/date/2026-01-18?page=1&limit=10`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getByDate(new Date(), new Date())).rejects.toThrow('Failed to fetch blog posts.');
    });
  });
});
