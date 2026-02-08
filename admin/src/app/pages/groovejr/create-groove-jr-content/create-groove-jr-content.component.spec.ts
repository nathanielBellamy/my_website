import { render, screen, fireEvent } from '@testing-library/angular';
import { RouterTestingModule } from '@angular/router/testing';
import { CreateGrooveJrContentComponent } from './create-groove-jr-content.component';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { GrooveJrContent } from '../../../models/data-models';

describe('CreateGrooveJrContentComponent', () => {

  let mockGrooveJrService: Partial<GrooveJrService>;
  let mockRouter: Partial<Router>;

  beforeEach(async () => {
    mockGrooveJrService = {
      createGrooveJrContent: jest.fn().mockReturnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jest.fn(),
    };

    await render(CreateGrooveJrContentComponent, {
      imports: [FormsModule, RouterTestingModule],
      providers: [
        { provide: GrooveJrService, useValue: mockGrooveJrService },
        { provide: Router, useValue: mockRouter },
      ],
    });
  });

  it('should create', async () => {
    expect(screen.getByText('Create New GrooveJr Content')).toBeTruthy();
  });

  it('should create GrooveJr content and navigate on success', async () => {
    const newContent: GrooveJrContent = { id: '', title: 'Test Title', content: 'Test Content' };

    fireEvent.input(screen.getByLabelText('Title:'), { target: { value: newContent.title } });
    fireEvent.input(screen.getByLabelText('Content:'), { target: { value: newContent.content } });

    fireEvent.click(screen.getByRole('button', { name: 'Create' }));

    // Await for the promise in createContent to resolve
    await Promise.resolve();

    expect(mockGrooveJrService.createGrooveJrContent).toHaveBeenCalledWith(newContent);
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/groovejr']);
  });

  it('should navigate back to list on goBack', async () => {
    fireEvent.click(screen.getByRole('button', { name: 'Back to List' }));
    await Promise.resolve(); // Allow promise to resolve for navigation
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/groovejr']);
  });
});
