import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { BlogComponent } from './blog.component';
import { BlogService } from '../../../services/blog.service';
import { BlogPost } from '../../../models/data-models';
import { RouterTestingModule } from '@angular/router/testing';

describe('BlogComponent', () => {
  let mockBlogService: Partial<BlogService>;
  const mockBlogPosts: BlogPost[] = [
    { id: '1', title: 'Blog 1', content: 'Content 1', author: { name: 'Author 1' }, tags: [{ id: '1', name: 'tag1' }], createdAt: new Date().toISOString(), updatedAt: new Date().toISOString() },
    { id: '2', title: 'Blog 2', content: 'Content 2', author: { name: 'Author 2' }, tags: [{ id: '2', name: 'tag2' }], createdAt: new Date().toISOString(), updatedAt: new Date().toISOString() },
  ];

  beforeEach(() => {
    mockBlogService = {
      getAllBlogPosts: jest.fn().mockReturnValue(Promise.resolve(mockBlogPosts)),
      deleteBlogPost: jest.fn().mockReturnValue(Promise.resolve()),
    };
  });

  it('should create', async () => {
    await render(BlogComponent, {
      imports: [RouterTestingModule],
      providers: [{ provide: BlogService, useValue: mockBlogService }],
    });
    expect(screen.getByText('Blog Posts')).toBeInTheDocument();
  });

  it('should fetch blog posts on ngOnInit', async () => {
    await render(BlogComponent, {
      imports: [RouterTestingModule],
      providers: [{ provide: BlogService, useValue: mockBlogService }],
    });

    expect(mockBlogService.getAllBlogPosts).toHaveBeenCalled();
    await waitFor(() => {
      expect(screen.getByText('Blog 1')).toBeInTheDocument();
      expect(screen.getByText('Blog 2')).toBeInTheDocument();
    });
  });

  it('should delete blog post and refresh the list', async () => {
    await render(BlogComponent, {
      imports: [RouterTestingModule],
      providers: [{ provide: BlogService, useValue: mockBlogService }],
    });

    await waitFor(() => {
      expect(screen.getByText('Blog 1')).toBeInTheDocument();
    });

    const deleteButton = screen.getByTestId('delete-button-1');
    await fireEvent.click(deleteButton);

    expect(mockBlogService.deleteBlogPost).toHaveBeenCalledWith('1');
    await waitFor(() => {
      expect(mockBlogService.getAllBlogPosts).toHaveBeenCalledTimes(2);
    });
  });
});
