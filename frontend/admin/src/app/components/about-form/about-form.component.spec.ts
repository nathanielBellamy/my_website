import { render, screen, fireEvent } from '@testing-library/angular';
import { AboutFormComponent } from './about-form.component';
import { AboutContent } from '../../models/data-models';

describe('AboutFormComponent', () => {
  const mockContent: AboutContent = {
    id: '1',
    title: 'Test Title',
    content: 'Test Content',
    order: 1,
    activatedAt: '2023-01-01T10:00:00.000Z',
    deactivatedAt: '2023-01-02T10:00:00.000Z',
  };

  it('should create', async () => {
    await render(AboutFormComponent);
    expect(screen.getByText('About Content')).toBeInTheDocument();
  });

  it('should populate form when contentData is provided', async () => {
    await render(AboutFormComponent, {
      componentInputs: {
        contentData: mockContent,
      },
    });

    expect(screen.getByLabelText('Title')).toHaveValue(mockContent.title);
    expect(screen.getByLabelText('Content')).toHaveValue(mockContent.content);
  });

  it('should emit submitForm when valid form is submitted', async () => {
    const submitSpy = jest.fn();
    await render(AboutFormComponent, {
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
    await render(AboutFormComponent, {
      on: {
        cancel: cancelSpy,
      },
    });

    await fireEvent.click(screen.getByRole('button', { name: /Cancel/i }));
    expect(cancelSpy).toHaveBeenCalled();
  });
});
