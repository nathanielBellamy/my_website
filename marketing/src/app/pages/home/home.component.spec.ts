import { render, screen, waitFor } from '@testing-library/angular';
import { provideMarkdown } from 'ngx-markdown';
import { HomeComponent } from './home.component';
import { CardComponent } from '../../components/card/card.component';
import { HomeStore } from './home.store';
import { signal, WritableSignal } from '@angular/core';
import { HomeContent } from '../../models/home.model';

const mockHomeContent: HomeContent[] = [
  { id: '1', title: 'Title 1', content: 'Body 1' },
  { id: '2', title: 'Title 2', content: 'Body 2' },
];

describe('HomeComponent', () => {
  let contentSignal: WritableSignal<HomeContent[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let loadMoreMock: jest.Mock;

  beforeEach(async () => {
    contentSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    loadMoreMock = jest.fn();

    await render(HomeComponent, {
      imports: [CardComponent],
      providers: [
        provideMarkdown(),
        {
          provide: HomeStore,
          useFactory: () => ({
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            allLoaded: signal(false),
            loadMore: loadMoreMock,
          }),
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
