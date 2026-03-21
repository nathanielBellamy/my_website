import { render, screen, fireEvent } from '@testing-library/angular';
import { WorkFormComponent } from './work-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { WorkContent } from '../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('WorkFormComponent', () => {
  const mockContent: WorkContent = {
    id: '1',
    title: 'Test Title',
    content: 'Test Content',
    order: 1,
    activatedAt: '2023-01-01T10:00:00.000Z',
    deactivatedAt: '2023-01-02T10:00:00.000Z',
  };

  it('should create', async () => {
    await render(WorkFormComponent, {
      providers: [provideMarkdown()],
    });
    expect(screen.getByText('Work Content')).toBeInTheDocument();
  });

  it('should populate form when contentData is provided', async () => {
    await render(WorkFormComponent, {
      providers: [provideMarkdown()],
      componentInputs: {
        contentData: mockContent,
      },
    });

    expect(screen.getByLabelText('Title')).toHaveValue(mockContent.title);
    expect(screen.getByLabelText('Content')).toHaveValue(mockContent.content);
    expect(screen.getByLabelText('Order')).toHaveValue(mockContent.order);
  });

  it('should emit submitForm when valid form is submitted', async () => {
    const submitSpy = jest.fn();
    await render(WorkFormComponent, {
      providers: [provideMarkdown()],
      on: {
        submitForm: submitSpy,
      },
    });

    await fireEvent.input(screen.getByLabelText('Title'), { target: { value: 'New Title' } });
    await fireEvent.input(screen.getByLabelText('Content'), { target: { value: 'New Content' } });
    await fireEvent.input(screen.getByLabelText('Order'), { target: { value: '2' } });

    await fireEvent.click(screen.getByRole('button', { name: /Save/i }));

    expect(submitSpy).toHaveBeenCalledWith(expect.objectContaining({
      title: 'New Title',
      content: 'New Content',
      order: 2,
    }));
  });

  it('should emit cancel when Cancel button is clicked', async () => {
    const cancelSpy = jest.fn();
    await render(WorkFormComponent, {
      providers: [provideMarkdown()],
      on: {
        cancel: cancelSpy,
      },
    });

    await fireEvent.click(screen.getByRole('button', { name: /Cancel/i }));
    expect(cancelSpy).toHaveBeenCalled();
  });

  it('should show error if dates are invalid', async () => {
    await render(WorkFormComponent, {
      providers: [provideMarkdown()],
    });

    await fireEvent.input(screen.getByLabelText('Activation Time'), { target: { value: '2023-01-02T10:00' } });
    await fireEvent.input(screen.getByLabelText('Deactivation Time'), { target: { value: '2023-01-01T10:00' } });

    expect(screen.getByText(/Deactivation time must be after activation time/i)).toBeInTheDocument();
  });
});
