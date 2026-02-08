import { render, screen, fireEvent, waitFor } from '@testing-library/angular';
import { AboutContentListComponent } from './about-content-list.component';
import { AboutService } from 'app/services/about.service';
import { AboutContent } from '../../models/data-models';


describe('AboutContentListComponent', () => {
  let mockAboutService: Partial<AboutService>;

  const mockAboutContent: AboutContent[] = [
    { id: '1', title: 'About 1', content: 'Content 1' },
    { id: '2', title: 'About 2', content: 'Content 2' },
  ];

  beforeEach(() => {
    mockAboutService = {
      getAllAboutContent: jest.fn().mockReturnValue(Promise.resolve(mockAboutContent)),
      deleteAboutContent: jest.fn().mockReturnValue(Promise.resolve()),
    };
  });

  it('should create', async () => {
    await render(AboutContentListComponent, {
      providers: [{ provide: AboutService, useValue: mockAboutService }],

    });
    expect(screen.getByText('About Content')).toBeInTheDocument();
  });

  it('should fetch about content on ngOnInit', async () => {
    await render(AboutContentListComponent, {
      providers: [{ provide: AboutService, useValue: mockAboutService }],

    });

    expect(mockAboutService.getAllAboutContent).toHaveBeenCalled();
    await waitFor(() => {
      expect(screen.getByText('About 1')).toBeInTheDocument();
      expect(screen.getByText('About 2')).toBeInTheDocument();
    });
  });

  it('should delete about content and refresh the list', async () => {
    await render(AboutContentListComponent, {
      providers: [{ provide: AboutService, useValue: mockAboutService }],

    });

    await waitFor(() => {
      expect(screen.getByText('About 1')).toBeInTheDocument();
    });

    const deleteButton = screen.getByTestId('delete-button-1');
    await fireEvent.click(deleteButton);

    expect(mockAboutService.deleteAboutContent).toHaveBeenCalledWith('1');
    await waitFor(() => {
      expect(mockAboutService.getAllAboutContent).toHaveBeenCalledTimes(2);
    });
  });
});