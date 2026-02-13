import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { Router } from '@angular/router';
import { GrooveJrContent } from '../../../models/data-models';

@Component({
  selector: 'app-create-groove-jr-content',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './create-groove-jr-content.component.html',
  styleUrl: './create-groove-jr-content.component.css',
})
export class CreateGrooveJrContentComponent {
  private readonly grooveJrService = inject(GrooveJrService);
  private readonly router = inject(Router);

  grooveJrContent: GrooveJrContent = {
    id: '',
    title: '',
    content: '',
  };

  activatedAtInput: string = '';
  deactivatedAtInput: string = '';

  async createContent() {
    try {
      if (this.activatedAtInput) {
        this.grooveJrContent.activatedAt = new Date(this.activatedAtInput).toISOString();
      }
      if (this.deactivatedAtInput) {
        this.grooveJrContent.deactivatedAt = new Date(this.deactivatedAtInput).toISOString();
      }
      await this.grooveJrService.createGrooveJrContent(this.grooveJrContent);
      this.router.navigate(['/groovejr']); // Navigate back to the list after creation
    } catch (error) {
      console.error('Error creating GrooveJr content:', error);
    }
  }

  goBack() {
    this.router.navigate(['/groovejr']);
  }
}
