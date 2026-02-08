import { Component, inject, OnInit, signal } from '@angular/core';
import { HomeService } from '../../../services/home.service';
import { HomeContent } from '../../../models/data-models';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-home-content-list',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './home-content-list.component.html',
  styleUrl: './home-content-list.component.css',
})
export class HomeContentListComponent implements OnInit {
  private readonly homeService = inject(HomeService);
  homeContent = signal<HomeContent[]>([]);

  ngOnInit() {
    this.fetchHomeContent();
  }

  fetchHomeContent() {
    this.homeService.getAllHomeContent().then((content) => {
      this.homeContent.set(content);
    }).catch((error) => {
      console.error('Error fetching home content:', error);
    });
  }

  deleteContent(id: string) {
    this.homeService.deleteHomeContent(id).then(() => {
      this.fetchHomeContent(); // Refresh the list after deletion
    }).catch((error) => {
      console.error('Error deleting home content:', error);
    });
  }
}
