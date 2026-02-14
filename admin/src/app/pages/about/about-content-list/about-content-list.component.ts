import { Component, inject, OnInit, signal, computed } from '@angular/core';
import { AboutService } from '../../../services/about.service';
import { AboutContent, FilterOptions } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-about-content-list',
  standalone: true,
  imports: [RouterLink, CommonModule],
  templateUrl: './about-content-list.component.html',
  styleUrl: './about-content-list.component.css',
})
export class AboutContentListComponent implements OnInit {
  private readonly aboutService = inject(AboutService);
  
  // State
  aboutContent = signal<AboutContent[]>([]);
  total = signal<number>(0);
  page = signal<number>(1);
  limit = signal<number>(10);
  showInactive = signal<boolean>(false);
  sortField = signal<string>('ordering');
  sortOrder = signal<'asc' | 'desc'>('asc');

  totalPages = computed(() => Math.ceil(this.total() / this.limit()));

  ngOnInit() {
    this.fetchAboutContent();
  }

  fetchAboutContent() {
    const options: Partial<FilterOptions> = {
      page: this.page(),
      limit: this.limit(),
      showInactive: this.showInactive(),
      sortField: this.sortField(),
      sortOrder: this.sortOrder(),
    };

    this.aboutService.getAllAboutContent(options).then((response) => {
      this.aboutContent.set(response.data);
      this.total.set(response.total);
    }).catch((error) => {
      console.error('Error fetching about content:', error);
    });
  }

  onPageChange(newPage: number) {
    if (newPage >= 1 && newPage <= this.totalPages()) {
      this.page.set(newPage);
      this.fetchAboutContent();
    }
  }

  toggleShowInactive() {
    this.showInactive.update(v => !v);
    this.page.set(1);
    this.fetchAboutContent();
  }

  onSort(field: string) {
    if (this.sortField() === field) {
      this.sortOrder.update(o => o === 'asc' ? 'desc' : 'asc');
    } else {
      this.sortField.set(field);
      this.sortOrder.set('asc');
    }
    this.fetchAboutContent();
  }

  deleteContent(id: string) {
    if(confirm('Are you sure you want to delete this content?')) {
        this.aboutService.deleteAboutContent(id).then(() => {
        this.fetchAboutContent();
        }).catch((error) => {
        console.error('Error deleting about content:', error);
        });
    }
  }
}
