import { Component, input, output, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule, AbstractControl, ValidationErrors } from '@angular/forms';
import { WorkContent } from '../../models/data-models';
import { MarkdownComponent } from 'ngx-markdown';
import { ImageGalleryComponent } from '../image-gallery/image-gallery.component';

@Component({
  selector: 'app-work-form',
  standalone: true,
  imports: [ReactiveFormsModule, MarkdownComponent, ImageGalleryComponent],
  templateUrl: './work-form.component.html',
  styleUrl: './work-form.component.css',
})
export class WorkFormComponent implements OnInit {
  contentData = input<WorkContent | undefined>();
  submitForm = output<WorkContent>();
  cancel = output<void>();

  private readonly fb = inject(FormBuilder);
  form!: FormGroup;
  showGallery = signal<boolean>(false);

  ngOnInit() {
    this.form = this.fb.group({
      id: [this.contentData()?.id || ''],
      title: [this.contentData()?.title || '', Validators.required],
      order: [this.contentData()?.order || 0, Validators.required],
      content: [this.contentData()?.content || '', Validators.required],
      activatedAt: [this.formatDateForInput(this.contentData()?.activatedAt)],
      deactivatedAt: [this.formatDateForInput(this.contentData()?.deactivatedAt)],
    }, { validators: this.dateRangeValidator });
  }

  toggleGallery() {
    this.showGallery.update(show => !show);
  }

  saveForm() {
    if (this.form.valid) {
      const formValue = this.form.value;
      const content: WorkContent = {
        ...formValue,
        activatedAt: formValue.activatedAt ? new Date(formValue.activatedAt).toISOString() : null,
        deactivatedAt: formValue.deactivatedAt ? new Date(formValue.deactivatedAt).toISOString() : null,
      };
      this.submitForm.emit(content);
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
