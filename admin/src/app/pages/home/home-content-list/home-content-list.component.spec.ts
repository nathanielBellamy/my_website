import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { HomeContentListComponent } from './home-content-list.component';
import { HomeService } from 'app/services/home.service';
import { of } from 'rxjs';
import { HomeContent } from '../../models/data-models';


describe('HomeContentListComponent', () => {
  let mockHomeService: Partial<HomeService>;
  const mockHomeContent: HomeContent[] = [
    { id: '1', title: 'Home 1', content: 'Content 1' },
    { id: '2', title: 'Home 2', content: 'Content 2' },
  ];
  beforeEach(() => {
    mockHomeService = {
      getAllHomeContent: jest.fn().mockReturnValue(Promise.resolve(mockHomeContent)),
      deleteHomeContent: jest.fn().mockReturnValue(Promise.resolve()),
    };
  });

  it('should create', async () => {
    await render(HomeContentListComponent, {
      providers: [{ provide: HomeService, useValue: mockHomeService }],

    });
    expect(screen.getByText('Home Content')).toBeInTheDocument();
  });

  it('should fetch home content on ngOnInit', async () => {
    await render(HomeContentListComponent, {
      providers: [{ provide: HomeService, useValue: mockHomeService }],

    });

    expect(mockHomeService.getAllHomeContent).toHaveBeenCalled();
    // Use waitFor to wait for the content to appear
    await waitFor(() => {
      expect(screen.getByText('Home 1')).toBeInTheDocument();
      expect(screen.getByText('Home 2')).toBeInTheDocument();
    });
  });

  it('should delete home content and refresh the list', async () => {
    await render(HomeContentListComponent, {
      providers: [{ provide: HomeService, useValue: mockHomeService }],

    });

    // Wait for initial content to be present
    await waitFor(() => {
      expect(screen.getByText('Home 1')).toBeInTheDocument();
    });

    // Simulate clicking the delete button for 'Home 1'
    const deleteButton = screen.getByTestId('delete-button-1');

    await fireEvent.click(deleteButton);

    expect(mockHomeService.deleteHomeContent).toHaveBeenCalledWith('1');
    // Wait for the list to refresh (getAllHomeContent called again)
    await waitFor(() => {
      expect(mockHomeService.getAllHomeContent).toHaveBeenCalledTimes(2); // Initial fetch + refresh
    });
  });
});
