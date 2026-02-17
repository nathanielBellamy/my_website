import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { EditBlogPostComponent } from './edit-blog-post.component';
import { BlogService } from '../../../services/blog.service';
import { ActivatedRoute, Router } from '@angular/router';
import { of } from 'rxjs';
import { BlogPost } from '../../../models/data-models';

describe('EditBlogPostComponent', () => {
  let mockBlogService: Partial<BlogService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockPost: BlogPost = {
    id: '1',
    title: 'Existing Post',
    content: 'Existing Content',
    author: { id: '1', name: 'Author' },
    tags: [],
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  };

  beforeEach(() => {
    mockBlogService = {
      getBlogPostById: jest.fn().mockReturnValue(Promise.resolve(mockPost)),
      updateBlogPost: jest.fn().mockReturnValue(Promise.resolve(mockPost)),
      getTags: jest.fn().mockReturnValue(Promise.resolve([])),
    };
    mockActivatedRoute = {
      snapshot: {
        paramMap: {
          get: jest.fn().mockReturnValue('1'),
        },
      },
    } as any;
    mockRouter = {
      navigate: jest.fn(),
    };
  });

  it('should create', async () => {
    await render(EditBlogPostComponent, {
      providers: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });
    expect(screen.getByText('Edit Blog Post')).toBeInTheDocument();
  });

  it('should fetch blog post on init and populate form', async () => {
    await render(EditBlogPostComponent, {
      providers: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });

    await waitFor(() => {
      expect(mockBlogService.getBlogPostById).toHaveBeenCalledWith('1');
      expect(screen.getByLabelText(/Title/i)).toHaveValue(mockPost.title);
      expect(screen.getByLabelText(/Content/i)).toHaveValue(mockPost.content);
    });
  });

  it('should update blog post and navigate on success', async () => {
    await render(EditBlogPostComponent, {
      providers: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });

    await waitFor(() => {
      expect(screen.getByLabelText(/Title/i)).toBeInTheDocument();
    });

    const titleInput = screen.getByLabelText(/Title/i);
    const saveButton = screen.getByRole('button', { name: /Save/i });

    await fireEvent.input(titleInput, { target: { value: 'Updated Title' } });
    await fireEvent.click(saveButton);

    await waitFor(() => {
      expect(mockBlogService.updateBlogPost).toHaveBeenCalledWith(expect.objectContaining({
        title: 'Updated Title',
      }));
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/blog']);
    });
  });

  it('should navigate back to list on Cancel', async () => {
    await render(EditBlogPostComponent, {
      providers: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });

    await waitFor(() => {
      expect(screen.getByRole('button', { name: /Cancel/i })).toBeInTheDocument();
    });
    const cancelButton = screen.getByRole('button', { name: /Cancel/i });
    await fireEvent.click(cancelButton);

    await waitFor(() => {
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/blog']);
    });
  });
});
