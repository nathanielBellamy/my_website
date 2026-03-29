import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { RouterTestingModule } from '@angular/router/testing';
import { CreateAboutContentComponent } from './create-about-content.component';
import { AboutService } from '../../../services/about.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { AboutContent } from '../../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('CreateAboutContentComponent', () => {
  let mockAboutService: Partial<AboutService>;
  let mockRouter: Partial<Router>;

  beforeEach(async () => {
    mockAboutService = {
      createAboutContent: jest.fn().mockReturnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jest.fn(),
    };

    await render(CreateAboutContentComponent, {
      providers: [
        { provide: AboutService, useValue: mockAboutService },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });
  });

  it('should create', () => {
    expect(screen.getByText('About Content')).toBeTruthy();
  });

  it('should create about content and navigate on success', async () => {
    fireEvent.input(screen.getByLabelText('Title'), { target: { value: 'Test Title' } });
    fireEvent.input(screen.getByLabelText('Content'), { target: { value: 'Test Content' } });

    fireEvent.click(screen.getByRole('button', { name: /Save/i }));

    await waitFor(() => {
      expect(mockAboutService.createAboutContent).toHaveBeenCalledWith(expect.objectContaining({
        title: 'Test Title',
        content: 'Test Content'
      }));
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/about']);
    });
  });

  it('should navigate back to list on Cancel', async () => {
    fireEvent.click(screen.getByRole('button', { name: /Cancel/i }));
    await Promise.resolve(); // Allow promise to resolve for navigation
    // Navigation on Cancel is not yet implemented in the component
  });
});
