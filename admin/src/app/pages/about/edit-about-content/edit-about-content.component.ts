import { Component, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AboutService } from '../../../services/about.service';
import { ActivatedRoute, Router } from '@angular/router';
import { AboutContent } from '../../../models/data-models';

@Component({
  selector: 'app-edit-about-content',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './edit-about-content.component.html',
  styleUrl: './edit-about-content.component.css',
})
export class EditAboutContentComponent implements OnInit {
  private readonly aboutService = inject(AboutService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  aboutContent = signal<AboutContent | undefined>(undefined);
  activatedAtInput: string = '';
  deactivatedAtInput: string = '';

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.aboutService.getAboutContentById(id).then(content => {
          this.aboutContent.set(content);
          this.activatedAtInput = this.formatDateForInput(content.activatedAt);
          this.deactivatedAtInput = this.formatDateForInput(content.deactivatedAt);
        }).catch(error => {
          console.error('Error fetching about content:', error);
        });
      }
    });
  }

  async updateContent() {
    if (this.aboutContent()) {
      const content = { ...this.aboutContent()! };
      content.activatedAt = this.activatedAtInput ? new Date(this.activatedAtInput).toISOString() : null;
      content.deactivatedAt = this.deactivatedAtInput ? new Date(this.deactivatedAtInput).toISOString() : null;
      try {
        await this.aboutService.updateAboutContent(content);
        this.router.navigate(['/about']); // Navigate back to the list after update
      } catch (error) {
        console.error('Error updating about content:', error);
      }
    }
  }

  goBack() {
    this.router.navigate(['/about']);
  }

  private formatDateForInput(dateStr?: string | null): string {
    if (!dateStr) return '';
    const d = new Date(dateStr);
    d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
    return d.toISOString().slice(0, 16);
  }
}
