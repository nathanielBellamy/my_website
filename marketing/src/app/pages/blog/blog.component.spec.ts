import { render, screen, waitFor } from '@testing-library/angular';
import { BlogComponent } from './blog.component';
import { CardComponent } from '../../components/card/card.component';
import { BlogStore } from './blog.store';
import { signal, WritableSignal } from '@angular/core';
import { BlogPost } from '../../models/blog.model';

const mockBlogPosts: BlogPost[] = [
  { id: 1, title: 'Title 1', content: 'Body 1', author: 'Author 1', date: '2026-01-18' },
  { id: 2, title: 'Title 2', content: 'Body 2', author: 'Author 2', date: '2026-01-18' },
];

describe('BlogComponent', () => {
  let postsSignal: WritableSignal<BlogPost[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let loadPostsMock: jest.Mock;

  beforeEach(() => {
    postsSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    loadPostsMock = jest.fn();
  });

  it('should call loadPosts on init', async () => {
    await render(BlogComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: BlogStore,
          useValue: {
            posts: postsSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadPosts: loadPostsMock,
          },
        },
      ],
    });
    expect(loadPostsMock).toHaveBeenCalled();
  });

  it('should render cards based on store posts', async () => {
    postsSignal.set(mockBlogPosts);
    await render(BlogComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: BlogStore,
          useValue: {
            posts: postsSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadPosts: loadPostsMock,
          },
        },
      ],
    });

    expect(screen.getAllByRole('article').length).toBe(2);
    screen.getByText('Title 1');
    screen.getByText('Body 2');
  });

  it('should show loading indicator when loading', async () => {
    loadingSignal.set(true);
    await render(BlogComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: BlogStore,
          useValue: {
            posts: postsSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadPosts: loadPostsMock,
          },
        },
      ],
    });
    screen.getByText('Loading...');
  });

  it('should show error message on error', async () => {
    errorSignal.set('Test Error');
    await render(BlogComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: BlogStore,
          useValue: {
            posts: postsSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadPosts: loadPostsMock,
          },
        },
      ],
    });
    screen.getByText('Error: Test Error');
  });
});