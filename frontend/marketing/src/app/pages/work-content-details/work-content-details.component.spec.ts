import { render, screen } from '@testing-library/angular';
import { WorkContentDetailsComponent } from './work-content-details.component';
import { WorkService } from '../../services/work.service';
import { ActivatedRoute } from '@angular/router';
import { WorkContent } from '../../models/work.model';
import { provideMarkdown } from 'ngx-markdown';

describe('WorkContentDetailsComponent', () => {
  const mockWorkContent: WorkContent = {
    id: '550e8400-e29b-41d4-a716-446655440000',
    title: 'Test Work Title',
    content: 'This is test work content.',
    order: 1,
  };

  const mockWorkService = {
    getById: jest.fn().mockResolvedValue(mockWorkContent),
  };

  const mockActivatedRoute = {
    snapshot: {
      paramMap: {
        get: jest.fn().mockReturnValue('123'),
      },
    },
  };

  it('should render work content details', async () => {
    await render(WorkContentDetailsComponent, {
      componentProviders: [
        { provide: WorkService, useValue: mockWorkService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('Test Work Title')).toBeTruthy();
    expect(await screen.findByText('This is test work content.')).toBeTruthy();
  });

  it('should show loading throbber if no id provided', async () => {
    const routeNoId = {
      snapshot: {
        paramMap: {
          get: jest.fn().mockReturnValue(null),
        },
      },
    };

    await render(WorkContentDetailsComponent, {
      componentProviders: [
        { provide: WorkService, useValue: mockWorkService },
        { provide: ActivatedRoute, useValue: routeNoId },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByRole('status', { name: 'Loading content' })).toBeTruthy();
    expect(screen.queryByText('No work content ID provided')).toBeNull();
  });

  it('should show loading throbber if fetch fails', async () => {
    const mockErrorService = {
      getById: jest.fn().mockRejectedValue(new Error('Fetch failed')),
    };

    await render(WorkContentDetailsComponent, {
      componentProviders: [
        { provide: WorkService, useValue: mockErrorService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByRole('status', { name: 'Loading content' })).toBeTruthy();
    expect(screen.queryByText('Failed to load work content')).toBeNull();
  });

  it('should decode a hex-encoded UUID from the URL before calling the service', async () => {
    const hexId = '550e8400e29b41d4a716446655440000';
    const expectedUuid = '550e8400-e29b-41d4-a716-446655440000';

    const hexIdRoute = {
      snapshot: {
        paramMap: {
          get: jest.fn().mockReturnValue(hexId),
        },
      },
    };

    const trackingService = {
      getById: jest.fn().mockResolvedValue(mockWorkContent),
    };

    await render(WorkContentDetailsComponent, {
      componentProviders: [
        { provide: WorkService, useValue: trackingService },
        { provide: ActivatedRoute, useValue: hexIdRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(trackingService.getById).toHaveBeenCalledWith(expectedUuid);
    expect(await screen.findByText('Test Work Title')).toBeTruthy();
  });
});
