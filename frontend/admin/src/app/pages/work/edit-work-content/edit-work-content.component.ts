import { Component, inject, OnInit, signal } from '@angular/core';
import { WorkService } from '../../../services/work.service';
import { ActivatedRoute, Router } from '@angular/router';
import { WorkContent } from '../../../models/data-models';
import { WorkFormComponent } from '../../../components/work-form/work-form.component';

@Component({
  selector: 'app-edit-work-content',
  standalone: true,
  imports: [WorkFormComponent],
  templateUrl: './edit-work-content.component.html',
  styleUrl: './edit-work-content.component.css',
})
export class EditWorkContentComponent implements OnInit {
  private readonly workService = inject(WorkService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  workContent = signal<WorkContent | undefined>(undefined);

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.workService.getWorkContentById(id).then(content => {
          this.workContent.set(content);
        }).catch(error => {
          console.error('Error fetching work content:', error);
        });
      }
    });
  }

  async updateContent(content: WorkContent) {
    try {
      await this.workService.updateWorkContent(content);
      this.router.navigate(['/work']);
    } catch (error) {
      console.error('Error updating work content:', error);
    }
  }

  goBack() {
    this.router.navigate(['/work']);
  }
}
