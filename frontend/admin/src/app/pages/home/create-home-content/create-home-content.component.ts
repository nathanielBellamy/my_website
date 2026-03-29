import { Component, inject } from '@angular/core';
import { HomeService } from '../../../services/home.service';
import { Router } from '@angular/router';
import { WorkContent } from '../../../models/data-models';
import { HomeFormComponent } from '../../../components/home-form/home-form.component';

@Component({
  selector: 'app-create-home-content',
  standalone: true,
  imports: [HomeFormComponent],
  templateUrl: './create-home-content.component.html',
  styleUrl: './create-home-content.component.css',
})
export class CreateWorkContentComponent {
  private readonly homeService = inject(HomeService);
  private readonly router = inject(Router);

  async createContent(content: WorkContent) {
    try {
      await this.homeService.createWorkContent(content);
      this.router.navigate(['/home']);
    } catch (error) {
      console.error('Error creating home content:', error);
    }
  }

  goBack() {
    this.router.navigate(['/home']);
  }
}
