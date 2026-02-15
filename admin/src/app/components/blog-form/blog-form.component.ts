import { Component, input, output, EventEmitter, OnInit, inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule, AbstractControl, ValidationErrors } from '@angular/forms';
import { BlogPost } from '../../models/data-models';
import { MarkdownComponent } from 'ngx-markdown';

@Component({
  selector: 'app-blog-form',
  standalone: true,
  imports: [ReactiveFormsModule, MarkdownComponent], // Add MarkdownComponent
  templateUrl: './blog-form.component.html',
  styleUrl: './blog-form.component.css',
})
export class BlogFormComponent implements OnInit {
  post = input<BlogPost | undefined>();
  submitForm = output<BlogPost>();
  cancel = output<void>();

  private readonly fb = inject(FormBuilder);
  blogForm!: FormGroup;

  ngOnInit() {
    // Default activatedAt to now for new posts
    const initialActivatedAt = this.post()?.activatedAt || (this.post() ? null : new Date().toISOString());

    this.blogForm = this.fb.group({
      id: [this.post()?.id || ''],
      title: [this.post()?.title || '', Validators.required],
      order: [this.post()?.order || 0, Validators.required],
      content: [this.post()?.content || '', Validators.required],
      author: this.fb.group({
        id: [this.post()?.author?.id || ''],
        name: [this.post()?.author?.name || '', Validators.required],
      }),
      tags: [this.post()?.tags?.map(tag => tag.name).join(', ') || ''], // Assuming tags are comma-separated strings for input
      createdAt: [this.post()?.createdAt || ''],
      updatedAt: [this.post()?.updatedAt || ''],
      activatedAt: [this.formatDateForInput(initialActivatedAt)],
      deactivatedAt: [this.formatDateForInput(this.post()?.deactivatedAt)],
    }, { validators: this.dateRangeValidator });
  }

  saveForm() {
    if (this.blogForm.valid) {
      const formValue = this.blogForm.value;
      const blogPost: BlogPost = {
        ...formValue,
        tags: formValue.tags.split(',').map((name: string) => ({ id: '', name: name.trim() })), // Convert back to Tag array
        activatedAt: formValue.activatedAt ? new Date(formValue.activatedAt).toISOString() : null,
        deactivatedAt: formValue.deactivatedAt ? new Date(formValue.deactivatedAt).toISOString() : null,
      };
      this.submitForm.emit(blogPost);
    }
  }

  private formatDateForInput(dateStr?: string | null): string {
    if (!dateStr) return '';
    const d = new Date(dateStr);
    d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
    return d.toISOString().slice(0, 16);
  }

  private dateRangeValidator(group: AbstractControl): ValidationErrors | null {
    const start = group.get('activatedAt')?.value;
    const end = group.get('deactivatedAt')?.value;
    if (start && end && new Date(start) >= new Date(end)) {
      return { dateRangeInvalid: true };
    }
    return null;
  }

  onCancel() {
    this.cancel.emit();
  }
}
