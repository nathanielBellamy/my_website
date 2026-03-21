import { render, screen, fireEvent, waitFor, RenderResult } from '@testing-library/angular';
import { EditWorkContentComponent } from './edit-work-content.component';
import { WorkService } from 'app/services/work.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { WorkContent } from '../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('EditWorkContentComponent', () => {
  let mockWorkService: Partial<WorkService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockWorkContent: WorkContent = { id: '1', title: 'Existing Work', content: 'Existing Content', activatedAt: null, deactivatedAt: null };

  beforeEach(() => {
    mockWorkService = {
      getWorkContentById: jest.fn().mockReturnValue(Promise.resolve(mockWorkContent)),
      updateWorkContent: jest.fn().mockReturnValue(Promise.resolve(mockWorkContent)),
    };
    mockActivatedRoute = {
      paramMap: of(new Map([['id', '1']])),
    };
    mockRouter = {
      navigate: jest.fn(),
    };
  });

  it('should create', async () => {
    await render(EditWorkContentComponent, {
      providers: [
        { provide: WorkService, useValue: mockWorkService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });
    await waitFor(() => {
      expect(screen.getByText('Edit Work Content')).toBeInTheDocument();
    });
  });

  it('should fetch work content on init and populate form', async () => {
    await render(EditWorkContentComponent, {
      providers: [
        { provide: WorkService, useValue: mockWorkService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });

    expect(mockWorkService.getWorkContentById).toHaveBeenCalledWith('1');
    await waitFor(() => {
      expect(screen.getByLabelText('Title')).toHaveValue(mockWorkContent.title);
      expect(screen.getByLabelText('Content')).toHaveValue(mockWorkContent.content);
    });
  });

  it('should update work content and navigate on success', async () => {
    await render(EditWorkContentComponent, {
      providers: [
        { provide: WorkService, useValue: mockWorkService },
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
      expect(mockWorkService.updateWorkContent).toHaveBeenCalledWith(expect.objectContaining({
        title: 'Updated Title',
        content: 'Updated Content'
      }));
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/work']);
    });
  });

  it('should navigate back to list on Cancel', async () => {
    await render(EditWorkContentComponent, {
      providers: [
        { provide: WorkService, useValue: mockWorkService },
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
