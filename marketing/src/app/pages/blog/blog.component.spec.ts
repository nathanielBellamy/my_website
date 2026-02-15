import { render, screen, waitFor } from '@testing-library/angular';
import { provideMarkdown } from 'ngx-markdown';
import { BlogComponent } from './blog.component';
import { CardComponent } from '../../components/card/card.component';
import { BlogStore } from './blog.store';
import { signal, WritableSignal } from '@angular/core';
import { BlogPost } from '../../models/blog-post.model';

const mockBlogPosts: BlogPost[] = [
  { id: '1', title: 'Title 1', content: 'Body 1', author: 'Author 1', tags: [], createdAt: '', updatedAt: '' },
  { id: '2', title: 'Title 2', content: 'Body 2', author: 'Author 2', tags: [], createdAt: '', updatedAt: '' },
];

describe('BlogComponent', () => {
  let postsSignal: WritableSignal<BlogPost[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let loadMoreMock: jest.Mock;

  beforeEach(async () => {
    postsSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    loadMoreMock = jest.fn();

    await render(BlogComponent, {
      imports: [CardComponent],
      providers: [
        provideMarkdown(),
        {
          provide: BlogStore,
          useValue: {
            posts: postsSignal,
            loading: loadingSignal,
            error: errorSignal,
            allLoaded: signal(false),
            loadMore: loadMoreMock,
          },
        },
      ],
    });
  });

  it('should call loadMore on init', () => {
    expect(loadMoreMock).toHaveBeenCalled();
  });

  it('should render cards based on store posts', async () => {
    postsSignal.set(mockBlogPosts);

    await waitFor(() => {
      expect(screen.getAllByRole('article').length).toBe(2);
      screen.getByText('Title 1');
      screen.getByText('Body 2');
    });
  });

  it('should show loading indicator when loading', async () => {
    loadingSignal.set(true);
    await waitFor(() => {
      screen.getByText('Loading...');
    });
  });

  it('should show error message on error', async () => {
    errorSignal.set('Test Error');
    await waitFor(() => {
      screen.getByText('Error: Test Error');
    });
  });
});