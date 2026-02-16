import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { GrooveJrContentListComponent } from './groove-jr-content-list.component';
import { GrooveJrService } from 'app/services/groove-jr.service';
import { GrooveJrContent } from '../../models/data-models';


describe('GrooveJrContentListComponent', () => {
  let mockGrooveJrService: Partial<GrooveJrService>;

  const mockGrooveJrContent: GrooveJrContent[] = [
    { id: '1', title: 'GrooveJr 1', content: 'Content 1' },
    { id: '2', title: 'GrooveJr 2', content: 'Content 2' },
  ];

  beforeEach(() => {
    mockGrooveJrService = {
      getAllGrooveJrContent: jest.fn().mockReturnValue(Promise.resolve({ data: mockGrooveJrContent, total: mockGrooveJrContent.length })),
      deleteGrooveJrContent: jest.fn().mockReturnValue(Promise.resolve()),
    };
  });

  it('should create', async () => {
    await render(GrooveJrContentListComponent, {
      providers: [{ provide: GrooveJrService, useValue: mockGrooveJrService }],

    });
    expect(screen.getByText('GrooveJr Content')).toBeInTheDocument();
  });

  it('should fetch GrooveJr content on ngOnInit', async () => {
    await render(GrooveJrContentListComponent, {
      providers: [{ provide: GrooveJrService, useValue: mockGrooveJrService }],

    });

    expect(mockGrooveJrService.getAllGrooveJrContent).toHaveBeenCalled();
    await waitFor(() => {
      expect(screen.getByText('GrooveJr 1')).toBeInTheDocument();
      expect(screen.getByText('GrooveJr 2')).toBeInTheDocument();
    });
  });

  it('should delete GrooveJr content and refresh the list', async () => {
    await render(GrooveJrContentListComponent, {
      providers: [{ provide: GrooveJrService, useValue: mockGrooveJrService }],

    });

    await waitFor(() => {
      expect(screen.getByText('GrooveJr 1')).toBeInTheDocument();
    });

    const deleteButton = screen.getByTestId('delete-button-1');
    await fireEvent.click(deleteButton);

    expect(mockGrooveJrService.deleteGrooveJrContent).toHaveBeenCalledWith('1');
    await waitFor(() => {
      expect(mockGrooveJrService.getAllGrooveJrContent).toHaveBeenCalledTimes(2);
    });
  });
});