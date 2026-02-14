import { Component, inject, OnInit, signal, computed } from '@angular/core';
import { HomeService } from '../../../services/home.service';
import { HomeContent, FilterOptions } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-home-content-list',
  standalone: true,
  imports: [RouterLink, CommonModule],
  templateUrl: './home-content-list.component.html',
  styleUrl: './home-content-list.component.css',
})
export class HomeContentListComponent implements OnInit {
  private readonly homeService = inject(HomeService);
  
  // State
  homeContent = signal<HomeContent[]>([]);
  total = signal<number>(0);
  page = signal<number>(1);
  limit = signal<number>(10);
  status = signal<'current' | 'inactive' | 'past' | 'future'>('current');
  sortField = signal<string>('ordering');
  sortOrder = signal<'asc' | 'desc'>('asc');

  totalPages = computed(() => Math.ceil(this.total() / this.limit()));

  ngOnInit() {
    this.fetchHomeContent();
  }

  fetchHomeContent() {
    const options: Partial<FilterOptions> = {
      page: this.page(),
      limit: this.limit(),
      status: this.status(),
      sortField: this.sortField(),
      sortOrder: this.sortOrder(),
    };

    this.homeService.getAllHomeContent(options).then((response) => {
      this.homeContent.set(response.data);
      this.total.set(response.total);
    }).catch((error) => {
      console.error('Error fetching home content:', error);
    });
  }

  onPageChange(newPage: number) {
    if (newPage >= 1 && newPage <= this.totalPages()) {
      this.page.set(newPage);
      this.fetchHomeContent();
    }
  }

  setStatus(newStatus: 'current' | 'inactive' | 'past' | 'future') {
    this.status.set(newStatus);
    this.page.set(1);
    this.fetchHomeContent();
  }

  onSort(field: string) {
    if (this.sortField() === field) {
      this.sortOrder.update(o => o === 'asc' ? 'desc' : 'asc');
    } else {
      this.sortField.set(field);
      this.sortOrder.set('asc');
    }
    this.fetchHomeContent();
  }

  deleteContent(id: string) {
    if(confirm('Are you sure you want to delete this content?')) {
        this.homeService.deleteHomeContent(id).then(() => {
        this.fetchHomeContent();
        }).catch((error) => {
        console.error('Error deleting home content:', error);
        });
    }
  }
}
