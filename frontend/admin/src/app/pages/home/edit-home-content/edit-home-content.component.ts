import { Component, inject, OnInit, signal } from '@angular/core';
import { HomeService } from '../../../services/home.service';
import { ActivatedRoute, Router } from '@angular/router';
import { WorkContent } from '../../../models/data-models';
import { HomeFormComponent } from '../../../components/home-form/home-form.component';

@Component({
  selector: 'app-edit-home-content',
  standalone: true,
  imports: [HomeFormComponent],
  templateUrl: './edit-home-content.component.html',
  styleUrl: './edit-home-content.component.css',
})
export class EditWorkContentComponent implements OnInit {
  private readonly homeService = inject(HomeService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  homeContent = signal<WorkContent | undefined>(undefined);

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.homeService.getWorkContentById(id).then(content => {
          this.homeContent.set(content);
        }).catch(error => {
          console.error('Error fetching home content:', error);
        });
      }
    });
  }

  async updateContent(content: WorkContent) {
    try {
      await this.homeService.updateWorkContent(content);
      this.router.navigate(['/home']);
    } catch (error) {
      console.error('Error updating home content:', error);
    }
  }

  goBack() {
    this.router.navigate(['/home']);
  }
}
