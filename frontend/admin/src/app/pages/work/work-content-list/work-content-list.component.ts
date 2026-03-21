import { Component, inject, OnInit, signal, computed } from '@angular/core';
import { WorkService } from '../../../services/work.service';
import { WorkContent, FilterOptions } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';
import { CsvControlsComponent } from '../../../components/csv-controls/csv-controls.component';

@Component({
  selector: 'app-work-content-list',
  standalone: true,
  imports: [RouterLink, CommonModule, CsvControlsComponent],
  templateUrl: './work-content-list.component.html',
  styleUrl: './work-content-list.component.css',
})
export class WorkContentListComponent implements OnInit {
  private readonly workService = inject(WorkService);

  // State
  workContent = signal<WorkContent[]>([]);
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

    this.workService.getAllWorkContent(options).then((response) => {
      this.workContent.set(response.data);
      this.total.set(response.total);
    }).catch((error) => {
      console.error('Error fetching work content:', error);
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
        this.workService.deleteWorkContent(id).then(() => {
        this.fetchWorkContent();
        }).catch((error) => {
        console.error('Error deleting work content:', error);
        });
    }
  }
}
