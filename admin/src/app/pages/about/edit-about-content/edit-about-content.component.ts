import { Component, inject, OnInit, signal } from '@angular/core';
import { AboutService } from '../../../services/about.service';
import { ActivatedRoute, Router } from '@angular/router';
import { AboutContent } from '../../../models/data-models';
import { AboutFormComponent } from '../../../components/about-form/about-form.component';

@Component({
  selector: 'app-edit-about-content',
  standalone: true,
  imports: [AboutFormComponent],
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

  async updateContent(content: AboutContent) {
    try {
      await this.aboutService.updateAboutContent(content);
      this.router.navigate(['/about']);
    } catch (error) {
      console.error('Error updating about content:', error);
    }
  }
}
