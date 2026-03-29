import { render, screen } from '@testing-library/angular';
import { GrooveJrContentDetailsComponent } from './groove-jr-content-details.component';
import { GrooveJrService } from '../../services/groove-jr.service';
import { ActivatedRoute } from '@angular/router';
import { GrooveJrContent } from '../../models/groove-jr.model';
import { provideMarkdown } from 'ngx-markdown';

describe('GrooveJrContentDetailsComponent', () => {
  const mockGrooveJrContent: GrooveJrContent = {
    id: '550e8400-e29b-41d4-a716-446655440000',
    title: 'Test GrooveJr Title',
    content: 'This is test GrooveJr content.',
    order: 1,
  };

  const mockGrooveJrService = {
    getById: jest.fn().mockResolvedValue(mockGrooveJrContent),
  };

  const mockActivatedRoute = {
    snapshot: {
      paramMap: {
        get: jest.fn().mockReturnValue('123'),
      },
    },
  };

  it('should render GrooveJr content details', async () => {
    await render(GrooveJrContentDetailsComponent, {
      componentProviders: [
        { provide: GrooveJrService, useValue: mockGrooveJrService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('Test GrooveJr Title')).toBeTruthy();
    expect(await screen.findByText('This is test GrooveJr content.')).toBeTruthy();
  });

  it('should show loading throbber if no id provided', async () => {
    const routeNoId = {
      snapshot: {
        paramMap: {
          get: jest.fn().mockReturnValue(null),
        },
      },
    };

    await render(GrooveJrContentDetailsComponent, {
      componentProviders: [
        { provide: GrooveJrService, useValue: mockGrooveJrService },
        { provide: ActivatedRoute, useValue: routeNoId },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByRole('status', { name: 'Loading content' })).toBeTruthy();
    expect(screen.queryByText('No GrooveJr content ID provided')).toBeNull();
  });

  it('should show loading throbber if fetch fails', async () => {
    const mockErrorService = {
      getById: jest.fn().mockRejectedValue(new Error('Fetch failed')),
    };

    await render(GrooveJrContentDetailsComponent, {
      componentProviders: [
        { provide: GrooveJrService, useValue: mockErrorService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByRole('status', { name: 'Loading content' })).toBeTruthy();
    expect(screen.queryByText('Failed to load GrooveJr content')).toBeNull();
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
      getById: jest.fn().mockResolvedValue(mockGrooveJrContent),
    };

    await render(GrooveJrContentDetailsComponent, {
      componentProviders: [
        { provide: GrooveJrService, useValue: trackingService },
        { provide: ActivatedRoute, useValue: hexIdRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(trackingService.getById).toHaveBeenCalledWith(expectedUuid);
    expect(await screen.findByText('Test GrooveJr Title')).toBeTruthy();
  });
});
