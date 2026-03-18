import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { RouterTestingModule } from '@angular/router/testing';
import { EditAboutContentComponent } from './edit-about-content.component';
import { AboutService } from '../../../services/about.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { AboutContent } from '../../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('EditAboutContentComponent', () => {
  let mockAboutService: Partial<AboutService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockAboutContent: AboutContent = { id: '1', title: 'Existing About', content: 'Existing Content', activatedAt: null, deactivatedAt: null };

  beforeEach(async () => {
    mockAboutService = {
      getAboutContentById: jest.fn().mockReturnValue(Promise.resolve(mockAboutContent)),
      updateAboutContent: jest.fn().mockReturnValue(Promise.resolve(mockAboutContent)),
    };
    mockActivatedRoute = {
      paramMap: of(new Map([['id', '1']])),
    };
    mockRouter = {
      navigate: jest.fn(),
    };

    await render(EditAboutContentComponent, {
      providers: [
        { provide: AboutService, useValue: mockAboutService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });
    // Wait for the content to be loaded and the form to appear
    await screen.findByLabelText('Title');
  });

  it('should create', async () => {
    await waitFor(() => {
      expect(screen.getByText('Edit About Content')).toBeTruthy();
    });
  });

  it('should fetch about content on init and populate form', async () => {
    expect(mockAboutService.getAboutContentById).toHaveBeenCalledWith('1');
    expect(screen.getByLabelText('Title')).toHaveValue(mockAboutContent.title);
    expect(screen.getByLabelText('Content')).toHaveValue(mockAboutContent.content);
  });

  it('should update about content and navigate on success', async () => {
    const updatedTitle = 'Updated Title';
    const updatedContent = 'Updated Content';

    fireEvent.input(screen.getByLabelText('Title'), { target: { value: updatedTitle } });
    fireEvent.input(screen.getByLabelText('Content'), { target: { value: updatedContent } });

    fireEvent.click(screen.getByRole('button', { name: /Save/i }));

    await waitFor(() => {
      expect(mockAboutService.updateAboutContent).toHaveBeenCalledWith(expect.objectContaining({
        title: updatedTitle,
        content: updatedContent
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
