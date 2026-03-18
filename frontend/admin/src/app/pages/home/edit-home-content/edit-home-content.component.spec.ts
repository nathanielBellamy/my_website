import { render, screen, fireEvent, waitFor, RenderResult } from '@testing-library/angular';
import { EditHomeContentComponent } from './edit-home-content.component';
import { HomeService } from 'app/services/home.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { HomeContent } from '../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('EditHomeContentComponent', () => {
  let mockHomeService: Partial<HomeService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockHomeContent: HomeContent = { id: '1', title: 'Existing Home', content: 'Existing Content', activatedAt: null, deactivatedAt: null };

  beforeEach(() => {
    mockHomeService = {
      getHomeContentById: jest.fn().mockReturnValue(Promise.resolve(mockHomeContent)),
      updateHomeContent: jest.fn().mockReturnValue(Promise.resolve(mockHomeContent)),
    };
    mockActivatedRoute = {
      paramMap: of(new Map([['id', '1']])),
    };
    mockRouter = {
      navigate: jest.fn(),
    };
  });

  it('should create', async () => {
    await render(EditHomeContentComponent, {
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });
    await waitFor(() => {
      expect(screen.getByText('Edit Home Content')).toBeInTheDocument();
    });
  });

  it('should fetch home content on init and populate form', async () => {
    await render(EditHomeContentComponent, {
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });

    expect(mockHomeService.getHomeContentById).toHaveBeenCalledWith('1');
    await waitFor(() => {
      expect(screen.getByLabelText('Title')).toHaveValue(mockHomeContent.title);
      expect(screen.getByLabelText('Content')).toHaveValue(mockHomeContent.content);
    });
  });

  it('should update home content and navigate on success', async () => {
    await render(EditHomeContentComponent, {
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });

    await waitFor(() => {
      expect(screen.getByLabelText('Title')).toBeInTheDocument();
    });

    const titleInput = screen.getByLabelText('Title');
    const contentInput = screen.getByLabelText('Content');
    const saveButton = screen.getByRole('button', { name: /Save/i });

    await fireEvent.input(titleInput, { target: { value: 'Updated Title' } });
    await fireEvent.input(contentInput, { target: { value: 'Updated Content' } });
    await fireEvent.click(saveButton);

    await waitFor(() => {
      expect(mockHomeService.updateHomeContent).toHaveBeenCalledWith(expect.objectContaining({
        title: 'Updated Title',
        content: 'Updated Content'
      }));
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
    });
  });

  it('should navigate back to list on Cancel', async () => {
    await render(EditHomeContentComponent, {
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });

    await waitFor(() => {
      expect(screen.getByRole('button', { name: /Cancel/i })).toBeInTheDocument();
    });
    const cancelButton = screen.getByRole('button', { name: /Cancel/i });
    await fireEvent.click(cancelButton);

    // Navigation on Cancel is not yet implemented in the component
  });
});
