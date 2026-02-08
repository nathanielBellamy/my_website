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

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.aboutService.getAboutContentById(id).then(content => {
          this.aboutContent.set(content);
        }).catch(error => {
          console.error('Error fetching about content:', error);
        });
      }
    });
  }

  async updateContent() {
    if (this.aboutContent()) {
      try {
        await this.aboutService.updateAboutContent(this.aboutContent() as AboutContent);
        this.router.navigate(['/about']); // Navigate back to the list after update
      } catch (error) {
        console.error('Error updating about content:', error);
      }
    }
  }

  goBack() {
    this.router.navigate(['/about']);
  }
}
