import { Component, inject } from '@angular/core';
import { WorkService } from '../../../services/work.service';
import { Router } from '@angular/router';
import { WorkContent } from '../../../models/data-models';
import { WorkFormComponent } from '../../../components/work-form/work-form.component';

@Component({
  selector: 'app-create-work-content',
  standalone: true,
  imports: [WorkFormComponent],
  templateUrl: './create-work-content.component.html',
  styleUrl: './create-work-content.component.css',
})
export class CreateWorkContentComponent {
  private readonly workService = inject(WorkService);
  private readonly router = inject(Router);

  async createContent(content: WorkContent) {
    try {
      await this.workService.createWorkContent(content);
      this.router.navigate(['/work']);
    } catch (error) {
      console.error('Error creating work content:', error);
    }
  }

  goBack() {
    this.router.navigate(['/work']);
  }
}
