import { Component, inject, OnInit, signal, computed } from '@angular/core';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { GrooveJrContent, FilterOptions } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-groove-jr-content-list',
  standalone: true,
  imports: [RouterLink, CommonModule],
  templateUrl: './groove-jr-content-list.component.html',
  styleUrl: './groove-jr-content-list.component.css',
})
export class GrooveJrContentListComponent implements OnInit {
  private readonly grooveJrService = inject(GrooveJrService);
  
  // State
  grooveJrContent = signal<GrooveJrContent[]>([]);
  total = signal<number>(0);
  page = signal<number>(1);
  limit = signal<number>(10);
  status = signal<'current' | 'inactive' | 'past' | 'future'>('current');
  sortField = signal<string>('ordering');
  sortOrder = signal<'asc' | 'desc'>('asc');

  totalPages = computed(() => Math.ceil(this.total() / this.limit()));

  ngOnInit() {
    this.fetchGrooveJrContent();
  }

  fetchGrooveJrContent() {
    const options: Partial<FilterOptions> = {
      page: this.page(),
      limit: this.limit(),
      status: this.status(),
      sortField: this.sortField(),
      sortOrder: this.sortOrder(),
    };

    this.grooveJrService.getAllGrooveJrContent(options).then((response) => {
      this.grooveJrContent.set(response.data);
      this.total.set(response.total);
    }).catch((error) => {
      console.error('Error fetching GrooveJr content:', error);
    });
  }

  onPageChange(newPage: number) {
    if (newPage >= 1 && newPage <= this.totalPages()) {
      this.page.set(newPage);
      this.fetchGrooveJrContent();
    }
  }

  setStatus(newStatus: 'current' | 'inactive' | 'past' | 'future') {
    this.status.set(newStatus);
    this.page.set(1);
    this.fetchGrooveJrContent();
  }

  onSort(field: string) {
    if (this.sortField() === field) {
      this.sortOrder.update(o => o === 'asc' ? 'desc' : 'asc');
    } else {
      this.sortField.set(field);
      this.sortOrder.set('asc');
    }
    this.fetchGrooveJrContent();
  }

  deleteContent(id: string) {
    if(confirm('Are you sure you want to delete this content?')) {
        this.grooveJrService.deleteGrooveJrContent(id).then(() => {
        this.fetchGrooveJrContent();
        }).catch((error) => {
        console.error('Error deleting GrooveJr content:', error);
        });
    }
  }
}
