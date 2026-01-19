import { render, screen, waitFor } from '@testing-library/angular';
import { AboutComponent } from './about.component';
import { AboutStore } from './about.store';
import { signal, WritableSignal } from '@angular/core';
import { AboutContent } from '../../models/about.model';

const mockAboutContent: AboutContent[] = [
  { id: 1, title: 'Title 1', body: 'Body 1' },
  { id: 2, title: 'Title 2', body: 'Body 2' },
];

describe('AboutComponent', () => {
  let contentSignal: WritableSignal<AboutContent[]>;
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
    await render(AboutComponent, {
      providers: [
        {
          provide: AboutStore,
          useValue: {
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          },
        },
      ],
    });
    expect(loadContentMock).toHaveBeenCalled();
  });

  it('should render content based on store content', async () => {
    contentSignal.set(mockAboutContent);
    await render(AboutComponent, {
      providers: [
        {
          provide: AboutStore,
          useValue: {
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          },
        },
      ],
    });

    screen.getByText('Title 1');
    screen.getByText('Body 2');
  });

  it('should show loading indicator when loading', async () => {
    loadingSignal.set(true);
    await render(AboutComponent, {
      providers: [
        {
          provide: AboutStore,
          useValue: {
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          },
        },
      ],
    });
    screen.getByText('Loading...');
  });

  it('should show error message on error', async () => {
    errorSignal.set('Test Error');
    await render(AboutComponent, {
      providers: [
        {
          provide: AboutStore,
          useValue: {
            content: contentSignal,
            loading: loadingSignal,
            error: errorSignal,
            loadContent: loadContentMock,
          },
        },
      ],
    });
    screen.getByText('Error: Test Error');
  });
});