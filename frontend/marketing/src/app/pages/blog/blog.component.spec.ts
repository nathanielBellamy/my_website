import { render, screen, waitFor } from '@testing-library/angular';
import { provideMarkdown } from 'ngx-markdown';
import { BlogComponent } from './blog.component';
import { CardComponent } from '../../components/card/card.component';
import { BlogStore } from './blog.store';
import { signal, WritableSignal } from '@angular/core';
import { BlogPost, Tag } from '../../models/blog-post.model';
import { provideRouter } from '@angular/router';

const mockBlogPosts: BlogPost[] = [
  { id: '1', title: 'Title 1', content: 'Body 1', author: { id: '1', name: 'Author 1' }, tags: [], createdAt: '', updatedAt: '', order: 1 },
  { id: '2', title: 'Title 2', content: 'Body 2', author: { id: '2', name: 'Author 2' }, tags: [], createdAt: '', updatedAt: '', order: 2 },
];

describe('BlogComponent', () => {
  let postsSignal: WritableSignal<BlogPost[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let availableTagsSignal: WritableSignal<Tag[]>;
  let selectedTagsSignal: WritableSignal<string[]>;
  let loadMoreMock: jest.Mock;
  let loadTagsMock: jest.Mock;
  let searchTagsMock: jest.Mock;
  let toggleTagMock: jest.Mock;

  beforeEach(async () => {
    postsSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    availableTagsSignal = signal([]);
    selectedTagsSignal = signal([]);
    loadMoreMock = jest.fn();
    loadTagsMock = jest.fn();
    searchTagsMock = jest.fn();
    toggleTagMock = jest.fn();

    await render(BlogComponent, {
      imports: [CardComponent],
      providers: [
        provideMarkdown(),
        provideRouter([]),
        {
          provide: BlogStore,
          useValue: {
            posts: postsSignal,
            loading: loadingSignal,
            error: errorSignal,
            allLoaded: signal(false),
            availableTags: availableTagsSignal,
            selectedTags: selectedTagsSignal,
            loadMore: loadMoreMock,
            loadTags: loadTagsMock,
            searchTags: searchTagsMock,
            toggleTag: toggleTagMock,
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
      screen.getByText('Body 1');
    });
  });

  it('should show loading indicator when loading', async () => {
    loadingSignal.set(true);
    await waitFor(() => {
      screen.getByText('Loading...');
    });
  });

  it.skip('should show error message on error', async () => {
    errorSignal.set('Test Error');
    await waitFor(() => {
      screen.getByText('Error: Test Error');
    });
  });
});