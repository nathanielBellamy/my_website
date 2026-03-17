import { render, screen } from '@testing-library/angular';
import { AboutContentDetailsComponent } from './about-content-details.component';
import { AboutService } from '../../services/about.service';
import { ActivatedRoute } from '@angular/router';
import { AboutContent } from '../../models/about.model';
import { provideMarkdown } from 'ngx-markdown';

describe('AboutContentDetailsComponent', () => {
  const mockAboutContent: AboutContent = {
    id: '550e8400-e29b-41d4-a716-446655440000',
    title: 'Test About Title',
    content: 'This is test about content.',
    order: 1,
  };

  const mockAboutService = {
    getById: jest.fn().mockResolvedValue(mockAboutContent),
  };

  const mockActivatedRoute = {
    snapshot: {
      paramMap: {
        get: jest.fn().mockReturnValue('123'),
      },
    },
  };

  it('should render about content details', async () => {
    await render(AboutContentDetailsComponent, {
      componentProviders: [
        { provide: AboutService, useValue: mockAboutService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('Test About Title')).toBeTruthy();
    expect(await screen.findByText('This is test about content.')).toBeTruthy();
  });

  it('should show error if no id provided', async () => {
    const routeNoId = {
      snapshot: {
        paramMap: {
          get: jest.fn().mockReturnValue(null),
        },
      },
    };

    await render(AboutContentDetailsComponent, {
      componentProviders: [
        { provide: AboutService, useValue: mockAboutService },
        { provide: ActivatedRoute, useValue: routeNoId },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('No about content ID provided')).toBeTruthy();
  });

  it('should show error if fetch fails', async () => {
    const mockErrorService = {
      getById: jest.fn().mockRejectedValue(new Error('Fetch failed')),
    };

    await render(AboutContentDetailsComponent, {
      componentProviders: [
        { provide: AboutService, useValue: mockErrorService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(await screen.findByText('Failed to load about content')).toBeTruthy();
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
      getById: jest.fn().mockResolvedValue(mockAboutContent),
    };

    await render(AboutContentDetailsComponent, {
      componentProviders: [
        { provide: AboutService, useValue: trackingService },
        { provide: ActivatedRoute, useValue: hexIdRoute },
      ],
      providers: [
        provideMarkdown(),
      ],
    });

    expect(trackingService.getById).toHaveBeenCalledWith(expectedUuid);
    expect(await screen.findByText('Test About Title')).toBeTruthy();
  });
});
