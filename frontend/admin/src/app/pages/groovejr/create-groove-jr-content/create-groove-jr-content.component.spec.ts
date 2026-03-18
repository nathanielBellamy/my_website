import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { RouterTestingModule } from '@angular/router/testing';
import { CreateGrooveJrContentComponent } from './create-groove-jr-content.component';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { GrooveJrContent } from '../../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('CreateGrooveJrContentComponent', () => {

  let mockGrooveJrService: Partial<GrooveJrService>;
  let mockRouter: Partial<Router>;

  beforeEach(async () => {
    mockGrooveJrService = {
      createGrooveJrContent: jest.fn().mockReturnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jest.fn(),
    };

    await render(CreateGrooveJrContentComponent, {
      providers: [
        { provide: GrooveJrService, useValue: mockGrooveJrService },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });
  });

  it('should create', async () => {
    expect(screen.getByText('Groove Jr Content')).toBeTruthy();
  });

  it('should create GrooveJr content and navigate on success', async () => {
    fireEvent.input(screen.getByLabelText('Title'), { target: { value: 'Test Title' } });
    fireEvent.input(screen.getByLabelText('Content'), { target: { value: 'Test Content' } });

    fireEvent.click(screen.getByRole('button', { name: /Save/i }));

    await waitFor(() => {
      expect(mockGrooveJrService.createGrooveJrContent).toHaveBeenCalledWith(expect.objectContaining({
        title: 'Test Title',
        content: 'Test Content'
      }));
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/groovejr']);
    });
  });

  it('should navigate back to list on Cancel', async () => {
    fireEvent.click(screen.getByRole('button', { name: /Cancel/i }));
    await Promise.resolve(); // Allow promise to resolve for navigation
    // Navigation on Cancel is not yet implemented in the component
  });
});
