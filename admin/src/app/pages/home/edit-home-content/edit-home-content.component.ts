import { Component, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HomeService } from '../../../services/home.service';
import { ActivatedRoute, Router } from '@angular/router';
import { HomeContent } from '../../../models/data-models';

@Component({
  selector: 'app-edit-home-content',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './edit-home-content.component.html',
  styleUrl: './edit-home-content.component.css',
})
export class EditHomeContentComponent implements OnInit {
  private readonly homeService = inject(HomeService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  homeContent = signal<HomeContent | undefined>(undefined);
  activatedAtInput: string = '';
  deactivatedAtInput: string = '';

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.homeService.getHomeContentById(id).then(content => {
          this.homeContent.set(content);
          this.activatedAtInput = this.formatDateForInput(content.activatedAt);
          this.deactivatedAtInput = this.formatDateForInput(content.deactivatedAt);
        }).catch(error => {
          console.error('Error fetching home content:', error);
        });
      }
    });
  }

  async updateContent() {
    if (this.homeContent()) {
      const content = { ...this.homeContent()! };
      content.activatedAt = this.activatedAtInput ? new Date(this.activatedAtInput).toISOString() : null;
      content.deactivatedAt = this.deactivatedAtInput ? new Date(this.deactivatedAtInput).toISOString() : null;
      try {
        await this.homeService.updateHomeContent(content);
        this.router.navigate(['/home']); // Navigate back to the list after update
      } catch (error) {
        console.error('Error updating home content:', error);
      }
    }
  }

  goBack() {
    this.router.navigate(['/home']);
  }

  private formatDateForInput(dateStr?: string | null): string {
    if (!dateStr) return '';
    const d = new Date(dateStr);
    d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
    return d.toISOString().slice(0, 16);
  }
}
