import { render, screen, fireEvent, waitFor, RenderResult } from '@testing-library/angular';
import { EditHomeContentComponent } from './edit-home-content.component';
import { HomeService } from 'app/services/home.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { HomeContent } from '../../models/data-models';

describe('EditHomeContentComponent', () => {
  let mockHomeService: Partial<HomeService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockHomeContent: HomeContent = { id: '1', title: 'Existing Home', content: 'Existing Content' };

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
      imports: [FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });
    expect(screen.getByText('Edit Home Content')).toBeInTheDocument();
  });

  it('should fetch home content on init and populate form', async () => {
    const { fixture } = await render(EditHomeContentComponent, {
      imports: [FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });

    expect(mockHomeService.getHomeContentById).toHaveBeenCalledWith('1');
    await waitFor(() => {
      expect(screen.getByLabelText('Title:')).toHaveValue(mockHomeContent.title);
      expect(screen.getByLabelText('Content:')).toHaveValue(mockHomeContent.content);
    });
  });

  it('should update home content and navigate on success', async () => {
    const { fixture } = await render(EditHomeContentComponent, {
      imports: [FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });

    await waitFor(() => { // Wait for the form to be rendered
      expect(screen.getByLabelText('Title:')).toBeInTheDocument();
    });

    // Ensure homeContent is not undefined after initial fetch
    await waitFor(() => {
      expect(fixture.componentInstance.homeContent()).toBeDefined();
    });

    // Update the homeContent signal correctly
    const currentHomeContent = fixture.componentInstance.homeContent();
    if (currentHomeContent) {
      fixture.componentInstance.homeContent.set({
        ...currentHomeContent,
        title: 'Updated Title',
        content: 'Updated Content'
      });
    }
    fixture.detectChanges(); // Trigger change detection

    // Directly call the component's method
    await fixture.componentInstance.updateContent();

    const updatedHomeContent: HomeContent = { id: '1', title: 'Updated Title', content: 'Updated Content' };

    expect(mockHomeService.updateHomeContent).toHaveBeenCalledWith(updatedHomeContent);
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });

  it('should navigate back to list on goBack', async () => {
    const { fixture } = await render(EditHomeContentComponent, {
      imports: [FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    });

    await waitFor(() => { // Wait for the back button to be rendered
      expect(screen.getByRole('button', { name: /Back to List/i })).toBeInTheDocument();
    });
    const backButton = screen.getByRole('button', { name: /Back to List/i });
    await fireEvent.click(backButton);

    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });
});
