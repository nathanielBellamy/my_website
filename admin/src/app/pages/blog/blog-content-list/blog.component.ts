import { Component, inject, OnInit, signal, computed } from '@angular/core';
import { BlogService } from '../../../services/blog.service';
import { BlogPost, Tag, FilterOptions } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [RouterLink, CommonModule],
  templateUrl: './blog.component.html',
  styleUrl: './blog.component.css',
})
export class BlogComponent implements OnInit {
  private readonly blogService = inject(BlogService);
  
  // State
  blogPosts = signal<BlogPost[]>([]);
  total = signal<number>(0);
  page = signal<number>(1);
  limit = signal<number>(10);
  status = signal<'current' | 'inactive' | 'past' | 'future'>('current');
  sortField = signal<string>('ordering');
  sortOrder = signal<'asc' | 'desc'>('asc');

  totalPages = computed(() => Math.ceil(this.total() / this.limit()));

  ngOnInit() {
    this.fetchBlogPosts();
  }

  fetchBlogPosts() {
    const options: Partial<FilterOptions> = {
      page: this.page(),
      limit: this.limit(),
      status: this.status(),
      sortField: this.sortField(),
      sortOrder: this.sortOrder(),
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
