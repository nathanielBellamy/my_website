import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { BlogService } from './blog.service';
import { BlogPost } from '../models/data-models';

describe('BlogService', () => {
  let service: BlogService;
  let httpTestingController: HttpTestingController;

  const mockPost: BlogPost = {
    id: '1',
    title: 'Test Post',
    content: 'Test Content',
    author: { id: '1', name: 'Author' },
    tags: [],
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  };

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [BlogService]
    });
    service = TestBed.inject(BlogService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should retrieve all blog posts', async () => {
    const mockResponse = { data: [mockPost], total: 1 };
    const promise = service.getAllBlogPosts();

    const req = httpTestingController.expectOne(req => req.url.startsWith('/api/admin/blog'));
    expect(req.request.method).toEqual('GET');
    req.flush(mockResponse);

    const result = await promise;
    expect(result.data).toEqual([mockPost]);
    expect(result.total).toBe(1);
  });

  it('should retrieve blog post by ID', async () => {
    const promise = service.getBlogPostById('1');

    const req = httpTestingController.expectOne('/api/admin/blog/1');
    expect(req.request.method).toEqual('GET');
    req.flush(mockPost);

    await expect(promise).resolves.toEqual(mockPost);
  });

  it('should create blog post', async () => {
    const promise = service.createBlogPost(mockPost);

    const req = httpTestingController.expectOne('/api/admin/blog');
    expect(req.request.method).toEqual('POST');
    expect(req.request.body).toEqual(mockPost);
    req.flush(mockPost);

    await expect(promise).resolves.toEqual(mockPost);
  });

  it('should update blog post', async () => {
    const promise = service.updateBlogPost(mockPost);

    const req = httpTestingController.expectOne('/api/admin/blog/1');
    expect(req.request.method).toEqual('PUT');
    expect(req.request.body).toEqual(mockPost);
    req.flush(mockPost);

    await expect(promise).resolves.toEqual(mockPost);
  });

  it('should delete blog post', async () => {
    const promise = service.deleteBlogPost('1');

    const req = httpTestingController.expectOne('/api/admin/blog/1');
    expect(req.request.method).toEqual('DELETE');
    req.flush(null);

    await expect(promise).resolves.toBeNull();
  });
});
