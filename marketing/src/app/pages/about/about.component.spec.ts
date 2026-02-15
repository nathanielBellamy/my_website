import { render, screen, waitFor } from '@testing-library/angular';
import { provideMarkdown } from 'ngx-markdown';
import { AboutComponent } from './about.component';
import { AboutStore } from './about.store';
import { signal, WritableSignal } from '@angular/core';
import { AboutContent } from '../../models/about.model';

const mockAboutContent: AboutContent[] = [
  { id: '1', title: 'Title 1', content: 'Body 1' },
  { id: '2', title: 'Title 2', content: 'Body 2' },
];

describe('AboutComponent', () => {
  let contentSignal: WritableSignal<AboutContent[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let loadMoreMock: jest.Mock;

  beforeEach(async () => {
    contentSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    loadMoreMock = jest.fn();

    await render(AboutComponent, {
      providers: [
        provideMarkdown(),
        {
          provide: AboutStore,
          useValue: {
            content: contentSignal,
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

  it('should render content based on store content', async () => {
    contentSignal.set(mockAboutContent);

    await waitFor(() => {
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