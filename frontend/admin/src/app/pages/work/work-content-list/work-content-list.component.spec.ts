import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { WorkContentListComponent } from './work-content-list.component';
import { WorkService } from 'app/services/work.service';
import { WorkContent } from '../../models/data-models';


describe('WorkContentListComponent', () => {
  let mockWorkService: Partial<WorkService>;
  const mockWorkContent: WorkContent[] = [
    { id: '1', title: 'Work 1', content: 'Content 1', order: 1 },
    { id: '2', title: 'Work 2', content: 'Content 2', order: 2 },
  ];
  beforeEach(() => {
    mockWorkService = {
      getAllWorkContent: jest.fn().mockReturnValue(Promise.resolve({ data: mockWorkContent, total: mockWorkContent.length })),
      deleteWorkContent: jest.fn().mockReturnValue(Promise.resolve()),
    };
  });

  it('should create', async () => {
    await render(WorkContentListComponent, {
      providers: [{ provide: WorkService, useValue: mockWorkService }],

    });
    expect(screen.getByText('Work Content')).toBeInTheDocument();
  });

  it('should fetch work content on ngOnInit', async () => {
    await render(WorkContentListComponent, {
      providers: [{ provide: WorkService, useValue: mockWorkService }],

    });

    expect(mockWorkService.getAllWorkContent).toHaveBeenCalled();
    // Use waitFor to wait for the content to appear
    await waitFor(() => {
      expect(screen.getByText('Work 1')).toBeInTheDocument();
      expect(screen.getByText('Work 2')).toBeInTheDocument();
    });
  });

  it('should delete work content and refresh the list', async () => {
    await render(WorkContentListComponent, {
      providers: [{ provide: WorkService, useValue: mockWorkService }],

    });

    // Wait for initial content to be present
    await waitFor(() => {
      expect(screen.getByText('Work 1')).toBeInTheDocument();
    });

    // Simulate clicking the delete button for 'Work 1'
    const deleteButton = screen.getByTestId('delete-button-1');

    await fireEvent.click(deleteButton);

    expect(mockWorkService.deleteWorkContent).toHaveBeenCalledWith('1');
    // Wait for the list to refresh (getAllWorkContent called again)
    await waitFor(() => {
      expect(mockWorkService.getAllWorkContent).toHaveBeenCalledTimes(2); // Initial fetch + refresh
    });
  });
});
