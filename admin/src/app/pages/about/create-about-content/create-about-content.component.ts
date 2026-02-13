import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AboutService } from '../../../services/about.service';
import { Router } from '@angular/router';
import { AboutContent } from '../../../models/data-models';

@Component({
  selector: 'app-create-about-content',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './create-about-content.component.html',
  styleUrl: './create-about-content.component.css',
})
export class CreateAboutContentComponent {
  private readonly aboutService = inject(AboutService);
  private readonly router = inject(Router);

  aboutContent: AboutContent = {
    id: '',
    title: '',
    content: '',
  };

  activatedAtInput: string = '';
  deactivatedAtInput: string = '';

  async createContent() {
    try {
      if (this.activatedAtInput) {
        this.aboutContent.activatedAt = new Date(this.activatedAtInput).toISOString();
      }
      if (this.deactivatedAtInput) {
        this.aboutContent.deactivatedAt = new Date(this.deactivatedAtInput).toISOString();
      }
      await this.aboutService.createAboutContent(this.aboutContent);
      this.router.navigate(['/about']); // Navigate back to the list after creation
    } catch (error) {
      console.error('Error creating about content:', error);
    }
  }

  goBack() {
    this.router.navigate(['/about']);
  }
}
