import { Component, inject } from '@angular/core';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { Router } from '@angular/router';
import { GrooveJrContent } from '../../../models/data-models';
import { GrooveJrFormComponent } from '../../../components/groove-jr-form/groove-jr-form.component';

@Component({
  selector: 'app-create-groove-jr-content',
  standalone: true,
  imports: [GrooveJrFormComponent],
  templateUrl: './create-groove-jr-content.component.html',
  styleUrl: './create-groove-jr-content.component.css',
})
export class CreateGrooveJrContentComponent {
  private readonly grooveJrService = inject(GrooveJrService);
  private readonly router = inject(Router);

  async createContent(content: GrooveJrContent) {
    try {
      await this.grooveJrService.createGrooveJrContent(content);
      this.router.navigate(['/groovejr']);
    } catch (error) {
      console.error('Error creating GrooveJr content:', error);
    }
  }
}
