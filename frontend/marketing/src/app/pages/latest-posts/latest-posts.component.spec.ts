import { render, screen, waitFor } from '@testing-library/angular';
import { provideMarkdown } from 'ngx-markdown';
import { LatestPostsComponent } from './latest-posts.component';
import { CardComponent } from '../../components/card/card.component';
import { LatestPostsStore } from './latest-posts.store';
import { signal, WritableSignal } from '@angular/core';
import { HomeContent } from '../../models/home.model';
import { provideRouter } from '@angular/router';

const mockHomeContent: HomeContent[] = [
  { id: '1', title: 'Title 1', content: 'Body 1', order: 1 },
  { id: '2', title: 'Title 2', content: 'Body 2', order: 2 },
];

describe('LatestPostsComponent', () => {
  let contentSignal: WritableSignal<HomeContent[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let allLoadedSignal: WritableSignal<boolean>;
  let loadMoreMock: jest.Mock;

  beforeEach(async () => {
    contentSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    allLoadedSignal = signal(false);
    loadMoreMock = jest.fn();

    await render(LatestPostsComponent, {
      imports: [CardComponent],
      providers: [
        provideMarkdown(),
        provideRouter([]),
        {
          provide: LatestPostsStore,
          useValue: {
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            allLoaded: allLoadedSignal,
            loadMore: loadMoreMock,
          },
        },
      ],
    });
  });

  it('should call loadMore on init', () => {
    expect(loadMoreMock).toHaveBeenCalled();
  });

  it('should render cards based on store content', async () => {
    contentSignal.set(mockHomeContent);

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
