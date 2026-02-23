import { Component, inject, OnInit, signal, computed, OnDestroy } from '@angular/core';
import { BlogService } from '../../../services/blog.service';
import { BlogPost, Tag, FilterOptions } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';
import { Subject, debounceTime, distinctUntilChanged, takeUntil } from 'rxjs';
import { CsvControlsComponent } from '../../../components/csv-controls/csv-controls.component';

@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [RouterLink, CommonModule, CsvControlsComponent],
  templateUrl: './blog.component.html',
  styleUrl: './blog.component.css',
})
export class BlogComponent implements OnInit, OnDestroy {
  private readonly blogService = inject(BlogService);
  
  // State
  blogPosts = signal<BlogPost[]>([]);
  total = signal<number>(0);
  page = signal<number>(1);
  limit = signal<number>(10);
  status = signal<'current' | 'inactive' | 'past' | 'future'>('current');
  sortField = signal<string>('ordering');
  sortOrder = signal<'asc' | 'desc'>('asc');
  availableTags = signal<Tag[]>([]);
  selectedTags = signal<string[]>([]);
  tagSearch = signal<string>('');

  private readonly searchSubject = new Subject<string>();
  private readonly destroy$ = new Subject<void>();

  totalPages = computed(() => Math.ceil(this.total() / this.limit()));

  ngOnInit() {
    this.fetchTags();
    this.fetchBlogPosts();

    this.searchSubject.pipe(
      debounceTime(500),
      distinctUntilChanged(),
      takeUntil(this.destroy$)
    ).subscribe(query => {
      this.tagSearch.set(query);
      this.fetchTags();
    });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }

  fetchTags() {
    this.blogService.getTags(this.tagSearch()).then(tags => {
        this.availableTags.set(tags);
    });
  }

  onSearchTags(event: Event) {
      const input = event.target as HTMLInputElement;
      this.searchSubject.next(input.value);
  }

  onToggleTag(tagId: string) {
      this.selectedTags.update(tags => {
          if (tags.includes(tagId)) {
              return tags.filter(id => id !== tagId);
          } else {
              return [...tags, tagId];
          }
      });
      this.page.set(1);
      this.fetchBlogPosts();
  }

  fetchBlogPosts() {
    const options: Partial<FilterOptions> = {
      page: this.page(),
      limit: this.limit(),
      status: this.status(),
      sortField: this.sortField(),
      sortOrder: this.sortOrder(),
      tags: this.selectedTags(),
    };

    this.blogService.getAllBlogPosts(options).then((response) => {
      this.blogPosts.set(response.data);
      this.total.set(response.total);
    }).catch((error) => {
      console.error('Error fetching blog posts:', error);
    });
  }

  onPageChange(newPage: number) {
    if (newPage >= 1 && newPage <= this.totalPages()) {
      this.page.set(newPage);
      this.fetchBlogPosts();
    }
  }

  setStatus(newStatus: 'current' | 'inactive' | 'past' | 'future') {
    this.status.set(newStatus);
    this.page.set(1); // Reset to page 1 on filter change
    this.fetchBlogPosts();
  }

  onSort(field: string) {
    if (this.sortField() === field) {
      this.sortOrder.update(o => o === 'asc' ? 'desc' : 'asc');
    } else {
      this.sortField.set(field);
      this.sortOrder.set('asc');
    }
    this.fetchBlogPosts();
  }

  deletePost(id: string) {
    if(confirm('Are you sure you want to delete this post?')) {
        this.blogService.deleteBlogPost(id).then(() => {
        this.fetchBlogPosts();
        }).catch((error) => {
        console.error('Error deleting blog post:', error);
        });
    }
  }

  getTagsAsString(tags: Tag[] | undefined): string {
    return tags ? tags.map(tag => tag.name).join(', ') : '';
  }
}
