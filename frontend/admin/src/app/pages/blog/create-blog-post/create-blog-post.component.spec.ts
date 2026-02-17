import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { CreateBlogPostComponent } from './create-blog-post.component';
import { BlogService } from '../../../services/blog.service';
import { Router } from '@angular/router';
import { BlogPost } from '../../../models/data-models';

describe('CreateBlogPostComponent', () => {
  let mockBlogService: Partial<BlogService>;
  let mockRouter: Partial<Router>;

  beforeEach(() => {
    mockBlogService = {
      createBlogPost: jest.fn().mockReturnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
      getTags: jest.fn().mockReturnValue(Promise.resolve([])),
    };
    mockRouter = {
      navigate: jest.fn(),
    };
  });

  it('should create', async () => {
    await render(CreateBlogPostComponent, {
      providers: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: Router, useValue: mockRouter },
      ],
    });
    expect(screen.getByText('Create New Blog Post')).toBeInTheDocument();
  });

  it('should create blog post and navigate on success', async () => {
    await render(CreateBlogPostComponent, {
      providers: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: Router, useValue: mockRouter },
      ],
    });

    const titleInput = screen.getByLabelText(/Title/i);
    const contentInput = screen.getByLabelText(/Content/i);
    const authorInput = screen.getByLabelText(/Author/i);
    const saveButton = screen.getByRole('button', { name: /Save/i });

    await fireEvent.input(titleInput, { target: { value: 'Test Title' } });
    await fireEvent.input(contentInput, { target: { value: 'Test Content' } });
    await fireEvent.input(authorInput, { target: { value: 'Test Author' } });
    await fireEvent.click(saveButton);

    await waitFor(() => {
      expect(mockBlogService.createBlogPost).toHaveBeenCalledWith(expect.objectContaining({
        title: 'Test Title',
        content: 'Test Content',
        author: expect.objectContaining({ name: 'Test Author' })
      }));
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/blog']);
    });
  });

  it('should navigate back to list on Cancel', async () => {
    await render(CreateBlogPostComponent, {
      providers: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: Router, useValue: mockRouter },
      ],
    });

    const cancelButton = screen.getByRole('button', { name: /Cancel/i });
    await fireEvent.click(cancelButton);

    await waitFor(() => {
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/blog']);
    });
  });
});
