import { inject, Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { BlogPost, FilterOptions, PaginatedResponse } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class BlogService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/api/admin/blog'; // Adjust as per your backend URL

  async getAllBlogPosts(options: Partial<FilterOptions> = {}): Promise<PaginatedResponse<BlogPost>> {
    const params: any = {
      page: options.page || 1,
      limit: options.limit || 10,
    };
    if (options.status) params.status = options.status;
    if (options.sortField) params.sort = options.sortField;
    if (options.sortOrder) params.order = options.sortOrder;

    return await firstValueFrom(this.http.get<PaginatedResponse<BlogPost>>(this.apiUrl, { params }));
  }

  async getBlogPostById(id: string): Promise<BlogPost> {
    return await firstValueFrom(this.http.get<BlogPost>(`${this.apiUrl}/${id}`));
  }

  async createBlogPost(post: BlogPost): Promise<BlogPost> {
    return await firstValueFrom(this.http.post<BlogPost>(this.apiUrl, post));
  }

  async updateBlogPost(post: BlogPost): Promise<BlogPost> {
    return await firstValueFrom(this.http.put<BlogPost>(`${this.apiUrl}/${post.id}`, post));
  }

  async deleteBlogPost(id: string): Promise<void> {
    return await firstValueFrom(this.http.delete<void>(`${this.apiUrl}/${id}`));
  }
}
