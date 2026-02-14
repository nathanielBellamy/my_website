import { Component, input, output, OnInit, inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule, AbstractControl, ValidationErrors } from '@angular/forms';
import { HomeContent } from '../../models/data-models';
import { MarkdownComponent } from 'ngx-markdown';

@Component({
  selector: 'app-home-form',
  standalone: true,
  imports: [ReactiveFormsModule, MarkdownComponent],
  templateUrl: './home-form.component.html',
  styleUrl: './home-form.component.css',
})
export class HomeFormComponent implements OnInit {
  contentData = input<HomeContent | undefined>();
  submitForm = output<HomeContent>();

  private readonly fb = inject(FormBuilder);
  form!: FormGroup;

  ngOnInit() {
    this.form = this.fb.group({
      id: [this.contentData()?.id || ''],
      title: [this.contentData()?.title || '', Validators.required],
      content: [this.contentData()?.content || '', Validators.required],
      activatedAt: [this.formatDateForInput(this.contentData()?.activatedAt)],
      deactivatedAt: [this.formatDateForInput(this.contentData()?.deactivatedAt)],
    }, { validators: this.dateRangeValidator });
  }

  saveForm() {
    if (this.form.valid) {
      const formValue = this.form.value;
      const content: HomeContent = {
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
}
