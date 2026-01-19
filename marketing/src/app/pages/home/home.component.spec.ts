import { render, screen, waitFor } from '@testing-library/angular';
import { HomeComponent } from './home.component';
import { CardComponent } from '../../components/card/card.component';
import { HomeStore } from './home.store';
import { signal, WritableSignal } from '@angular/core';
import { HomeContent } from '../../models/home.model';

const mockHomeContent: HomeContent[] = [
  { id: 1, title: 'Title 1', body: 'Body 1' },
  { id: 2, title: 'Title 2', body: 'Body 2' },
];

describe('HomeComponent', () => {
  let contentSignal: WritableSignal<HomeContent[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let loadContentMock: jest.Mock;

  beforeEach(() => {
    contentSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    loadContentMock = jest.fn();
  });

  it('should call loadContent on init', async () => {
    await render(HomeComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: HomeStore,
          useFactory: () => ({
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          }),
        },
      ],
    });
    expect(loadContentMock).toHaveBeenCalled();
  });

  it('should render cards based on store content', async () => {
    contentSignal.set(mockHomeContent);
    await render(HomeComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: HomeStore,
          useFactory: () => ({
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          }),
        },
      ],
    });

    expect(screen.getAllByRole('article').length).toBe(2);
    screen.getByText('Title 1');
    screen.getByText('Body 2');
  });

  it('should show loading indicator when loading', async () => {
    loadingSignal.set(true);
    await render(HomeComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: HomeStore,
          useFactory: () => ({
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          }),
        },
      ],
    });
    screen.getByText('Loading...');
  });

  it('should show error message on error', async () => {
    errorSignal.set('Test Error');
    await render(HomeComponent, {
      imports: [CardComponent],
      providers: [
        {
          provide: HomeStore,
          useFactory: () => ({
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          }),
        },
      ],
    });
    screen.getByText('Error: Test Error');
  });
});
