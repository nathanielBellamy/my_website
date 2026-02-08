import { render, screen, fireEvent, RenderResult } from '@testing-library/angular';
import { CreateHomeContentComponent } from './create-home-content.component';
import { HomeService } from 'app/services/home.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { HomeContent } from '../../models/data-models';

describe('CreateHomeContentComponent', () => {
  let mockHomeService: Partial<HomeService>;
  let mockRouter: Partial<Router>;

  beforeEach(() => {
    mockHomeService = {
      createHomeContent: jest.fn().mockReturnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jest.fn(),
    };
  });

  it('should create', async () => {
    const { fixture } = await render(CreateHomeContentComponent, {
      imports: [FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: Router, useValue: mockRouter },
      ],
    });
    expect(screen.getByText('Create New Home Content')).toBeInTheDocument();
  });

  it('should create home content and navigate on success', async () => {
    const { fixture } = await render(CreateHomeContentComponent, {
      imports: [FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: Router, useValue: mockRouter },
      ],
    });

    // Set the component's homeContent directly
    fixture.componentInstance.homeContent.title = 'Test Title';
    fixture.componentInstance.homeContent.content = 'Test Content';
    fixture.detectChanges(); // Trigger change detection

    const createButton = screen.getByRole('button', { name: /Create/i });
    // Instead of clicking the button, directly call the component's method
    await fixture.componentInstance.createContent();

    const newContent: HomeContent = { id: '', title: 'Test Title', content: 'Test Content' };
    expect(mockHomeService.createHomeContent).toHaveBeenCalledWith(newContent);
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });

  it('should navigate back to list on goBack', async () => {
    const { fixture } = await render(CreateHomeContentComponent, {
      imports: [FormsModule],
      providers: [
        { provide: HomeService, useValue: mockHomeService },
        { provide: Router, useValue: mockRouter },
      ],
    });

    const backButton = screen.getByRole('button', { name: /Back to List/i });
    await fireEvent.click(backButton);

    expect(mockRouter.navigate).toHaveBeenCalledWith(['/home']);
  });
});