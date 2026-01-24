import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { BlogService } from './blog.service';
import { BlogPost } from '../models/blog-post.model';
import { environment } from '../../environments/environment';

const mockBlogPosts: BlogPost[] = [
  { id: '1', title: 'Post 1', content: 'Content 1', author: 'Author 1', tags: [], createdAt: '', updatedAt: '' },
  { id: '2', title: 'Post 2', content: 'Content 2', author: 'Author 2', tags: [], createdAt: '', updatedAt: '' },
];

describe('BlogService', () => {
  let service: BlogService;
  let httpMock: HttpTestingController;
  const API_URL = `${environment.API_BASE_URL}/marketing/blog`;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [BlogService],
    });
    service = TestBed.inject(BlogService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  describe('getAll', () => {
    it('should return paginated blog posts', async () => {
      const promise = service.getAll(1, 10);
      const req = httpMock.expectOne(`${API_URL}?page=1&limit=10`);
      expect(req.request.method).toBe('GET');
      req.flush(mockBlogPosts);
      const posts = await promise;
      expect(posts).toEqual(mockBlogPosts);
    });
  });

  describe('getById', () => {
    it('should return a single blog post', async () => {
      const post = mockBlogPosts[0];
      const promise = service.getById('1');
      const req = httpMock.expectOne(`${API_URL}/1`);
      expect(req.request.method).toBe('GET');
      req.flush(post);
      const result = await promise;
      expect(result).toEqual(post);
    });
  });
});
