import { Component, inject } from '@angular/core';
import { HomeService } from '../../../services/home.service';
import { Router } from '@angular/router';
import { HomeContent } from '../../../models/data-models';
import { HomeFormComponent } from '../../../components/home-form/home-form.component';

@Component({
  selector: 'app-create-home-content',
  standalone: true,
  imports: [HomeFormComponent],
  templateUrl: './create-home-content.component.html',
  styleUrl: './create-home-content.component.css',
})
export class CreateHomeContentComponent {
  private readonly homeService = inject(HomeService);
  private readonly router = inject(Router);

  async createContent(content: HomeContent) {
    try {
      await this.homeService.createHomeContent(content);
      this.router.navigate(['/home']);
    } catch (error) {
      console.error('Error creating home content:', error);
    }
  }
}
