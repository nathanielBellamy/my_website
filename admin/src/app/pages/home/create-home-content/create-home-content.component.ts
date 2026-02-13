import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HomeService } from '../../../services/home.service';
import { Router } from '@angular/router';
import { HomeContent } from '../../../models/data-models';

@Component({
  selector: 'app-create-home-content',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './create-home-content.component.html',
  styleUrl: './create-home-content.component.css',
})
export class CreateHomeContentComponent {
  private readonly homeService = inject(HomeService);
  private readonly router = inject(Router);

  homeContent: HomeContent = {
    id: '',
    title: '',
    content: '',
  };

  activatedAtInput: string = '';
  deactivatedAtInput: string = '';

  async createContent() {
    try {
      if (this.activatedAtInput) {
        this.homeContent.activatedAt = new Date(this.activatedAtInput).toISOString();
      }
      if (this.deactivatedAtInput) {
        this.homeContent.deactivatedAt = new Date(this.deactivatedAtInput).toISOString();
      }
      await this.homeService.createHomeContent(this.homeContent);
      this.router.navigate(['/home']); // Navigate back to the list after creation
    } catch (error) {
      console.error('Error creating home content:', error);
    }
  }

  goBack() {
    this.router.navigate(['/home']);
  }
}
