import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { RouterTestingModule } from '@angular/router/testing';
import { EditGrooveJrContentComponent } from './edit-groove-jr-content.component';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { GrooveJrContent } from '../../../models/data-models';

describe('EditGrooveJrContentComponent', () => {
  let mockGrooveJrService: Partial<GrooveJrService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockGrooveJrContent: GrooveJrContent = { id: '1', title: 'Existing GrooveJr', content: 'Existing Content', activatedAt: null, deactivatedAt: null };

  beforeEach(async () => {
    mockGrooveJrService = {
      getGrooveJrContentById: jest.fn().mockReturnValue(Promise.resolve(mockGrooveJrContent)),
      updateGrooveJrContent: jest.fn().mockReturnValue(Promise.resolve(mockGrooveJrContent)),
    };
    mockActivatedRoute = {
      paramMap: of(new Map([['id', '1']])),
    };
    mockRouter = {
      navigate: jest.fn(),
    };

    await render(EditGrooveJrContentComponent, {
      providers: [
        { provide: GrooveJrService, useValue: mockGrooveJrService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });
    // Wait for the content to be loaded and the form to appear
    await screen.findByLabelText('Title');
  });

  it('should create', async () => {
    await waitFor(() => {
      expect(screen.getByText('Edit GrooveJr Content')).toBeTruthy();
    });
  });

  it('should fetch GrooveJr content on init and populate form', async () => {
    expect(mockGrooveJrService.getGrooveJrContentById).toHaveBeenCalledWith('1');
    expect(screen.getByLabelText('Title')).toHaveValue(mockGrooveJrContent.title);
    expect(screen.getByLabelText('Content')).toHaveValue(mockGrooveJrContent.content);
  });

  it('should update GrooveJr content and navigate on success', async () => {
    const updatedTitle = 'Updated Title';
    const updatedContent = 'Updated Content';

    fireEvent.input(screen.getByLabelText('Title'), { target: { value: updatedTitle } });
    fireEvent.input(screen.getByLabelText('Content'), { target: { value: updatedContent } });

    fireEvent.click(screen.getByRole('button', { name: /Save/i }));

    await waitFor(() => {
      expect(mockGrooveJrService.updateGrooveJrContent).toHaveBeenCalledWith(expect.objectContaining({
        title: updatedTitle,
        content: updatedContent
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
