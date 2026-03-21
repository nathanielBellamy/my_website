import { Component, inject, OnInit, signal, computed } from '@angular/core';
import { HomeService } from '../../../services/home.service';
import { WorkContent, FilterOptions } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';
import { CsvControlsComponent } from '../../../components/csv-controls/csv-controls.component';

@Component({
  selector: 'app-home-content-list',
  standalone: true,
  imports: [RouterLink, CommonModule, CsvControlsComponent],
  templateUrl: './home-content-list.component.html',
  styleUrl: './home-content-list.component.css',
})
export class WorkContentListComponent implements OnInit {
  private readonly homeService = inject(HomeService);
  
  // State
  homeContent = signal<WorkContent[]>([]);
  total = signal<number>(0);
  page = signal<number>(1);
  limit = signal<number>(10);
  status = signal<'current' | 'inactive' | 'past' | 'future'>('current');
  sortField = signal<string>('ordering');
  sortOrder = signal<'asc' | 'desc'>('asc');

  totalPages = computed(() => Math.ceil(this.total() / this.limit()));

  ngOnInit() {
    this.fetchWorkContent();
  }

  fetchWorkContent() {
    const options: Partial<FilterOptions> = {
      page: this.page(),
      limit: this.limit(),
      status: this.status(),
      sortField: this.sortField(),
      sortOrder: this.sortOrder(),
    };

    this.homeService.getAllWorkContent(options).then((response) => {
      this.homeContent.set(response.data);
      this.total.set(response.total);
    }).catch((error) => {
      console.error('Error fetching home content:', error);
    });
  }

  onPageChange(newPage: number) {
    if (newPage >= 1 && newPage <= this.totalPages()) {
      this.page.set(newPage);
      this.fetchWorkContent();
    }
  }

  setStatus(newStatus: 'current' | 'inactive' | 'past' | 'future') {
    this.status.set(newStatus);
    this.page.set(1);
    this.fetchWorkContent();
  }

  onSort(field: string) {
    if (this.sortField() === field) {
      this.sortOrder.update(o => o === 'asc' ? 'desc' : 'asc');
    } else {
      this.sortField.set(field);
      this.sortOrder.set('asc');
    }
    this.fetchWorkContent();
  }

  deleteContent(id: string) {
    if(confirm('Are you sure you want to delete this content?')) {
        this.homeService.deleteWorkContent(id).then(() => {
        this.fetchWorkContent();
        }).catch((error) => {
        console.error('Error deleting home content:', error);
        });
    }
  }
}
