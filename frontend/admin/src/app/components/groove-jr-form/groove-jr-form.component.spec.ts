import { render, screen, fireEvent } from '@testing-library/angular';
import { GrooveJrFormComponent } from './groove-jr-form.component';
import { GrooveJrContent } from '../../models/data-models';
import { provideMarkdown } from 'ngx-markdown';

describe('GrooveJrFormComponent', () => {
  const mockContent: GrooveJrContent = {
    id: '1',
    title: 'Test Title',
    content: 'Test Content',
    order: 1,
    activatedAt: '2023-01-01T10:00:00.000Z',
    deactivatedAt: '2023-01-02T10:00:00.000Z',
  };

  it('should create', async () => {
    await render(GrooveJrFormComponent, {
      providers: [provideMarkdown()],
    });
    expect(screen.getByText('Groove Jr Content')).toBeInTheDocument();
  });

  it('should populate form when contentData is provided', async () => {
    await render(GrooveJrFormComponent, {
      providers: [provideMarkdown()],
      componentInputs: {
        contentData: mockContent,
      },
    });

    expect(screen.getByLabelText('Title')).toHaveValue(mockContent.title);
    expect(screen.getByLabelText('Content')).toHaveValue(mockContent.content);
  });

  it('should emit submitForm when valid form is submitted', async () => {
    const submitSpy = jest.fn();
    await render(GrooveJrFormComponent, {
      providers: [provideMarkdown()],
      on: {
        submitForm: submitSpy,
      },
    });

    await fireEvent.input(screen.getByLabelText('Title'), { target: { value: 'New Title' } });
    await fireEvent.input(screen.getByLabelText('Content'), { target: { value: 'New Content' } });

    await fireEvent.click(screen.getByRole('button', { name: /Save/i }));

    expect(submitSpy).toHaveBeenCalledWith(expect.objectContaining({
      title: 'New Title',
      content: 'New Content',
    }));
  });

  it('should emit cancel when Cancel button is clicked', async () => {
    const cancelSpy = jest.fn();
    await render(GrooveJrFormComponent, {
      providers: [provideMarkdown()],
      on: {
        cancel: cancelSpy,
      },
    });

    await fireEvent.click(screen.getByRole('button', { name: /Cancel/i }));
    expect(cancelSpy).toHaveBeenCalled();
  });
});
