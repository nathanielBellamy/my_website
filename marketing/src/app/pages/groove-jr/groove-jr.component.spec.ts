import { render, screen, waitFor } from '@testing-library/angular';
import { GrooveJrComponent } from './groove-jr.component';
import { GrooveJrStore } from './groove-jr.store';
import { signal, WritableSignal } from '@angular/core';
import { GrooveJrContent } from '../../models/groove-jr.model';

const mockGrooveJrContent: GrooveJrContent[] = [
  { id: 1, title: 'Title 1', body: 'Body 1', url: 'http://test.com' },
  { id: 2, title: 'Title 2', body: 'Body 2', url: 'http://test2.com' },
];

describe('GrooveJrComponent', () => {
  let contentSignal: WritableSignal<GrooveJrContent[]>;
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
    await render(GrooveJrComponent, {
      providers: [
        {
          provide: GrooveJrStore,
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
    contentSignal.set(mockGrooveJrContent);
    await render(GrooveJrComponent, {
      providers: [
        {
          provide: GrooveJrStore,
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
    expect(screen.getAllByText('Learn more').length).toBe(2);
  });

  it('should show loading indicator when loading', async () => {
    loadingSignal.set(true);
    await render(GrooveJrComponent, {
      providers: [
        {
          provide: GrooveJrStore,
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
    await render(GrooveJrComponent, {
      providers: [
        {
          provide: GrooveJrStore,
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
