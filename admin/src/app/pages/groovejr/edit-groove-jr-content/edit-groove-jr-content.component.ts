import { Component, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { ActivatedRoute, Router } from '@angular/router';
import { GrooveJrContent } from '../../../models/data-models';

@Component({
  selector: 'app-edit-groove-jr-content',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './edit-groove-jr-content.component.html',
  styleUrl: './edit-groove-jr-content.component.css',
})
export class EditGrooveJrContentComponent implements OnInit {
  private readonly grooveJrService = inject(GrooveJrService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  grooveJrContent = signal<GrooveJrContent | undefined>(undefined);
  activatedAtInput: string = '';
  deactivatedAtInput: string = '';

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.grooveJrService.getGrooveJrContentById(id).then(content => {
          this.grooveJrContent.set(content);
          this.activatedAtInput = this.formatDateForInput(content.activatedAt);
          this.deactivatedAtInput = this.formatDateForInput(content.deactivatedAt);
        }).catch(error => {
          console.error('Error fetching GrooveJr content:', error);
        });
      }
    });
  }

  async updateContent() {
    if (this.grooveJrContent()) {
      const content = { ...this.grooveJrContent()! };
      content.activatedAt = this.activatedAtInput ? new Date(this.activatedAtInput).toISOString() : null;
      content.deactivatedAt = this.deactivatedAtInput ? new Date(this.deactivatedAtInput).toISOString() : null;
      try {
        await this.grooveJrService.updateGrooveJrContent(content);
        this.router.navigate(['/groovejr']); // Navigate back to the list after update
      } catch (error) {
        console.error('Error updating GrooveJr content:', error);
      }
    }
  }

  goBack() {
    this.router.navigate(['/groovejr']);
  }

  private formatDateForInput(dateStr?: string | null): string {
    if (!dateStr) return '';
    const d = new Date(dateStr);
    d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
    return d.toISOString().slice(0, 16);
  }
}
