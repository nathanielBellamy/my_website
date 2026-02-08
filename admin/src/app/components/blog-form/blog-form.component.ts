import { Component, input, output, EventEmitter, OnInit, inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { BlogPost } from '../../models/data-models';
import { JsonPipe } from '@angular/common'; // Import JsonPipe for debugging

@Component({
  selector: 'app-blog-form',
  standalone: true,
  imports: [ReactiveFormsModule, JsonPipe], // Add ReactiveFormsModule and JsonPipe
  templateUrl: './blog-form.component.html',
  styleUrl: './blog-form.component.css',
})
export class BlogFormComponent implements OnInit {
  post = input<BlogPost | undefined>();
  submitForm = output<BlogPost>();

  private readonly fb = inject(FormBuilder);
  blogForm!: FormGroup;

  ngOnInit() {
    this.blogForm = this.fb.group({
      id: [this.post()?.id || ''],
      title: [this.post()?.title || '', Validators.required],
      content: [this.post()?.content || '', Validators.required],
      author: this.fb.group({
        id: [this.post()?.author?.id || ''],
        name: [this.post()?.author?.name || '', Validators.required],
      }),
      tags: [this.post()?.tags?.map(tag => tag.name).join(', ') || ''], // Assuming tags are comma-separated strings for input
      createdAt: [this.post()?.createdAt || ''],
      updatedAt: [this.post()?.updatedAt || ''],
    });
  }

  saveForm() {
    if (this.blogForm.valid) {
      const formValue = this.blogForm.value;
      const blogPost: BlogPost = {
        ...formValue,
        tags: formValue.tags.split(',').map((name: string) => ({ id: '', name: name.trim() })), // Convert back to Tag array
      };
      this.submitForm.emit(blogPost);
    }
  }
}
