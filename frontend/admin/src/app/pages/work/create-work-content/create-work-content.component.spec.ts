import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { CreateWorkContentComponent } from './create-work-content.component';
import { WorkService } from 'app/services/work.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { WorkContent } from '../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('CreateWorkContentComponent', () => {
  let mockWorkService: Partial<WorkService>;
  let mockRouter: Partial<Router>;

  beforeEach(() => {
    mockWorkService = {
      createWorkContent: jest.fn().mockReturnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test', order: 1 })),
    };
    mockRouter = {
      navigate: jest.fn(),
    };
  });

  it('should create', async () => {
    await render(CreateWorkContentComponent, {
      providers: [
        { provide: WorkService, useValue: mockWorkService },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });
    expect(screen.getByText('Work Content')).toBeInTheDocument();
  });

  it('should create work content and navigate on success', async () => {
    await render(CreateWorkContentComponent, {
      providers: [
        { provide: WorkService, useValue: mockWorkService },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });

    const titleInput = screen.getByLabelText('Title');
    const contentInput = screen.getByLabelText('Content');
    const saveButton = screen.getByRole('button', { name: /Save/i });

    await fireEvent.input(titleInput, { target: { value: 'Test Title' } });
    await fireEvent.input(contentInput, { target: { value: 'Test Content' } });
    await fireEvent.click(saveButton);

    await waitFor(() => {
      expect(mockWorkService.createWorkContent).toHaveBeenCalledWith(expect.objectContaining({
        title: 'Test Title',
        content: 'Test Content',
      }));
      expect(mockRouter.navigate).toHaveBeenCalledWith(['/work']);
    });
  });

  it('should navigate back to list on Cancel', async () => {
    await render(CreateWorkContentComponent, {
      providers: [
        { provide: WorkService, useValue: mockWorkService },
        { provide: Router, useValue: mockRouter },
        provideMarkdown(),
      ],
    });

    const cancelButton = screen.getByRole('button', { name: /Cancel/i });
    await fireEvent.click(cancelButton);
  });
});
