import { render, screen, fireEvent } from '@testing-library/angular';
import { RouterTestingModule } from '@angular/router/testing';
import { CreateAboutContentComponent } from './create-about-content.component';
import { AboutService } from '../../../services/about.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { AboutContent } from '../../../models/data-models';

describe('CreateAboutContentComponent', () => {
describe('CreateAboutContentComponent', () => {
  let mockAboutService: Partial<AboutService>;
  let mockRouter: Partial<Router>;

  beforeEach(async () => {
    mockAboutService = {
      createAboutContent: jest.fn().mockReturnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jest.fn(),
    };

    await render(CreateAboutContentComponent, {
      imports: [FormsModule, RouterTestingModule],
      providers: [
        { provide: AboutService, useValue: mockAboutService },
        { provide: Router, useValue: mockRouter },
      ],
    });
  });

  it('should create', () => {
    expect(screen.getByText('Create New About Content')).toBeTruthy();
  });

  it('should create about content and navigate on success', async () => {
    const newContent: AboutContent = { id: '', title: 'Test Title', content: 'Test Content' };

    fireEvent.input(screen.getByLabelText('Title:'), { target: { value: newContent.title } });
    fireEvent.input(screen.getByLabelText('Content:'), { target: { value: newContent.content } });

    fireEvent.click(screen.getByRole('button', { name: 'Create' }));

    await Promise.resolve(); // Allow promise to resolve for content creation and navigation

    expect(mockAboutService.createAboutContent).toHaveBeenCalledWith(newContent);
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/about']);
  });

  it('should navigate back to list on goBack', async () => {
    fireEvent.click(screen.getByRole('button', { name: 'Back to List' }));
    await Promise.resolve(); // Allow promise to resolve for navigation
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/about']);
  });
});
