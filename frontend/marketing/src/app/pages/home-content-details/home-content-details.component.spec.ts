import { render, screen } from '@testing-library/angular';
import { HomeContentDetailsComponent } from './home-content-details.component';
import { HomeService } from '../../services/home.service';
import { ActivatedRoute } from '@angular/router';
import { HomeContent } from '../../models/home.model';
import { provideMarkdown } from 'ngx-markdown';

describe('HomeContentDetailsComponent', () => {
  const mockHomeContent: HomeContent = {
    id: '550e8400-e29b-41d4-a716-446655440000',
    title: 'Test Home Title',
    content: 'This is test home content.',
    order: 1,
  };

  const mockHomeService = {
    getById: jest.fn().mockResolvedValue(mockHomeContent),
  };

  const mockActivatedRoute = {
    snapshot: {
      paramMap: {
        get: jest.fn().mockReturnValue('123'),
      },
    },
  };

  it('should render home content details', async () => {
    await render(HomeContentDetailsComponent, {
      componentProviders: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('Test Home Title')).toBeTruthy();
    expect(await screen.findByText('This is test home content.')).toBeTruthy();
  });

  it('should show error if no id provided', async () => {
    const routeNoId = {
      snapshot: {
        paramMap: {
          get: jest.fn().mockReturnValue(null),
        },
      },
    };

    await render(HomeContentDetailsComponent, {
      componentProviders: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: routeNoId },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('No home content ID provided')).toBeTruthy();
  });

  it('should show error if fetch fails', async () => {
    const mockErrorService = {
      getById: jest.fn().mockRejectedValue(new Error('Fetch failed')),
    };

    await render(HomeContentDetailsComponent, {
      componentProviders: [
        { provide: HomeService, useValue: mockErrorService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('Failed to load home content')).toBeTruthy();
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
      getById: jest.fn().mockResolvedValue(mockHomeContent),
    };

    await render(HomeContentDetailsComponent, {
      componentProviders: [
        { provide: HomeService, useValue: trackingService },
        { provide: ActivatedRoute, useValue: hexIdRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(trackingService.getById).toHaveBeenCalledWith(expectedUuid);
    expect(await screen.findByText('Test Home Title')).toBeTruthy();
  });
});
