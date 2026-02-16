import { render, screen, fireEvent } from '@testing-library/angular';
import { BlogFormComponent } from './blog-form.component';
import { BlogPost } from '../../models/data-models';

describe('BlogFormComponent', () => {
  const mockPost: BlogPost = {
    id: '1',
    title: 'Test Post',
    content: 'Test Content',
    author: { id: '1', name: 'Author' },
    tags: [{ id: '1', name: 'tag1' }],
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  };

  it('should create', async () => {
    await render(BlogFormComponent);
    expect(screen.getByText('Blog Post')).toBeInTheDocument();
  });

  it('should populate form when post is provided', async () => {
    await render(BlogFormComponent, {
      componentInputs: {
        post: mockPost,
      },
    });

    expect(screen.getByLabelText('Title')).toHaveValue(mockPost.title);
    expect(screen.getByLabelText('Content')).toHaveValue(mockPost.content);
    expect(screen.getByLabelText('Author')).toHaveValue(mockPost.author.name);
    expect(screen.getByLabelText(/Tags/i)).toHaveValue('tag1');
  });

  it('should emit submitForm when valid form is submitted', async () => {
    const submitSpy = jest.fn();
    await render(BlogFormComponent, {
      on: {
        submitForm: submitSpy,
      },
    });

    await fireEvent.input(screen.getByLabelText('Title'), { target: { value: 'New Post' } });
    await fireEvent.input(screen.getByLabelText('Content'), { target: { value: 'New Content' } });
    await fireEvent.input(screen.getByLabelText('Author'), { target: { value: 'New Author' } });
    await fireEvent.input(screen.getByLabelText(/Tags/i), { target: { value: 't1, t2' } });

    await fireEvent.click(screen.getByRole('button', { name: /Save/i }));

    expect(submitSpy).toHaveBeenCalledWith(expect.objectContaining({
      title: 'New Post',
      content: 'New Content',
      author: expect.objectContaining({ name: 'New Author' }),
      tags: expect.arrayContaining([
        expect.objectContaining({ name: 't1' }),
        expect.objectContaining({ name: 't2' }),
      ]),
    }));
  });

  it('should emit cancel when Cancel button is clicked', async () => {
    const cancelSpy = jest.fn();
    await render(BlogFormComponent, {
      on: {
        cancel: cancelSpy,
      },
    });

    await fireEvent.click(screen.getByRole('button', { name: /Cancel/i }));
    expect(cancelSpy).toHaveBeenCalled();
  });
});
