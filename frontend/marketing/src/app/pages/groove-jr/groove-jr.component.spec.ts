import { render, screen, waitFor } from '@testing-library/angular';
import { provideMarkdown } from 'ngx-markdown';
import { GrooveJrComponent } from './groove-jr.component';
import { GrooveJrStore } from './groove-jr.store';
import { signal, WritableSignal } from '@angular/core';
import { GrooveJrContent } from '../../models/groove-jr.model';
import { RenderResult } from '@testing-library/angular';
import { provideRouter } from '@angular/router';

const mockGrooveJrContent: GrooveJrContent[] = [
  { id: '1', title: 'Title 1', content: 'Body 1', order: 1 },
  { id: '2', title: 'Title 2', content: 'Body 2', order: 2 },
];

describe('GrooveJrComponent', () => {
  let contentSignal: WritableSignal<GrooveJrContent[]>;
  let loadingSignal: WritableSignal<boolean>;
  let errorSignal: WritableSignal<string | null>;
  let loadMoreMock: jest.Mock;
  let fixture: RenderResult<GrooveJrComponent>['fixture'];

  beforeEach(async () => {
    contentSignal = signal([]);
    loadingSignal = signal(false);
    errorSignal = signal(null);
    loadMoreMock = jest.fn();

    const renderResult = await render(GrooveJrComponent, {
      providers: [
        provideMarkdown(),
        provideRouter([]),
        {
          provide: GrooveJrStore,
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
    fixture = renderResult.fixture;
  });

  it('should call loadMore on init', () => {
    expect(loadMoreMock).toHaveBeenCalled();
  });

  it('should render content based on store content', async () => {
    contentSignal.set(mockGrooveJrContent);
    fixture.detectChanges();
    await waitFor(() => {
      screen.getByText('Title 1');
      screen.getByText('Body 2');
    });
  });

  it('should show loading indicator when loading', async () => {
    loadingSignal.set(true);
    fixture.detectChanges();
    await waitFor(() => {
      screen.getByText('Loading...');
    });
  });

  it('should show error message on error', async () => {
    errorSignal.set('Test Error');
    fixture.detectChanges();
    await waitFor(() => {
      screen.getByText('Error: Test Error');
    });
  });

  it('should render scroll indicator', () => {
    screen.getByTestId('scroll-to-old-site-preview');
  });
});
