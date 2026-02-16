import { render, screen } from '@testing-library/angular';
import { BlogContentDetailsComponent } from './blog-content-details.component';
import { BlogService } from '../../services/blog.service';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { BlogPost } from '../../models/blog-post.model';
import { provideMarkdown } from 'ngx-markdown';

describe('BlogContentDetailsComponent', () => {
  const mockBlogPost: BlogPost = {
    id: '123',
    title: 'Test Post',
    content: 'This is a test post content.',
    author: { id: '1', name: 'Test Author' },
    tags: [{ id: '1', name: 'test' }, { id: '2', name: 'angular' }],
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
    order: 1
  };

  const mockBlogService = {
    getById: jest.fn().mockResolvedValue(mockBlogPost),
  };

  const mockActivatedRoute = {
    snapshot: {
      paramMap: {
        get: jest.fn().mockReturnValue('123'),
      },
    },
  };

  it('should render blog post details', async () => {
    await render(BlogContentDetailsComponent, {
      componentProviders: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown()
      ]
    });

    expect(await screen.findByText('Test Post')).toBeTruthy();
    expect(await screen.findByText('This is a test post content.')).toBeTruthy();
    expect(await screen.findByText('#test')).toBeTruthy();
    expect(await screen.findByText('#angular')).toBeTruthy();
  });

  it('should show error if no id provided', async () => {
    const routeNoId = {
      snapshot: {
        paramMap: {
          get: jest.fn().mockReturnValue(null),
        },
      },
    };

    await render(BlogContentDetailsComponent, {
      componentProviders: [
        { provide: BlogService, useValue: mockBlogService },
        { provide: ActivatedRoute, useValue: routeNoId },
      ],
      providers: [
        provideMarkdown()
      ]
    });

    expect(await screen.findByText('No blog post ID provided')).toBeTruthy();
  });

  it('should show error if fetch fails', async () => {
    const consoleSpy = jest.spyOn(console, 'error').mockImplementation(() => {});
    const mockErrorService = {
      getById: jest.fn().mockRejectedValue(new Error('Fetch failed')),
    };

    await render(BlogContentDetailsComponent, {
      componentProviders: [
        { provide: BlogService, useValue: mockErrorService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown()
      ]
    });

    expect(await screen.findByText('Failed to load blog post')).toBeTruthy();
    consoleSpy.mockRestore();
  });
});
