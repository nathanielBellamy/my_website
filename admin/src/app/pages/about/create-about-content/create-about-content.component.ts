import { Component, inject } from '@angular/core';
import { AboutService } from '../../../services/about.service';
import { Router } from '@angular/router';
import { AboutContent } from '../../../models/data-models';
import { AboutFormComponent } from '../../../components/about-form/about-form.component';

@Component({
  selector: 'app-create-about-content',
  standalone: true,
  imports: [AboutFormComponent],
  templateUrl: './create-about-content.component.html',
  styleUrl: './create-about-content.component.css',
})
export class CreateAboutContentComponent {
  private readonly aboutService = inject(AboutService);
  private readonly router = inject(Router);

  async createContent(content: AboutContent) {
    try {
      await this.aboutService.createAboutContent(content);
      this.router.navigate(['/about']);
    } catch (error) {
      console.error('Error creating about content:', error);
    }
  }
}
