import { Component, inject, OnInit, signal } from '@angular/core';
import { AboutService } from '../../../services/about.service';
import { AboutContent } from '../../../models/data-models';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-about-content-list',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './about-content-list.component.html',
  styleUrl: './about-content-list.component.css',
})
export class AboutContentListComponent implements OnInit {
  private readonly aboutService = inject(AboutService);
  aboutContent = signal<AboutContent[]>([]);

  ngOnInit() {
    this.fetchAboutContent();
  }

  fetchAboutContent() {
    this.aboutService.getAllAboutContent().then((content) => {
      this.aboutContent.set(content);
    }).catch((error) => {
      console.error('Error fetching about content:', error);
    });
  }

  deleteContent(id: string) {
    this.aboutService.deleteAboutContent(id).then(() => {
      this.fetchAboutContent(); // Refresh the list after deletion
    }).catch((error) => {
      console.error('Error deleting about content:', error);
    });
  }
}
