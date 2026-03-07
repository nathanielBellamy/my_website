import { render, screen } from '@testing-library/angular';
import { GrooveJrContentDetailsComponent } from './groove-jr-content-details.component';
import { GrooveJrService } from '../../services/groove-jr.service';
import { ActivatedRoute } from '@angular/router';
import { GrooveJrContent } from '../../models/groove-jr.model';
import { provideMarkdown } from 'ngx-markdown';

describe('GrooveJrContentDetailsComponent', () => {
  const mockGrooveJrContent: GrooveJrContent = {
    id: '123',
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

  it('should show error if no id provided', async () => {
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

    expect(await screen.findByText('No GrooveJr content ID provided')).toBeTruthy();
  });

  it('should show error if fetch fails', async () => {
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

    expect(await screen.findByText('Failed to load GrooveJr content')).toBeTruthy();
  });
});
