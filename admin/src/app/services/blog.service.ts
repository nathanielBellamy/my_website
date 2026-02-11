import { inject, Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { BlogPost } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class BlogService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/api/admin/blog'; // Adjust as per your backend URL

  async getAllBlogPosts(): Promise<BlogPost[]> {
    return await firstValueFrom(this.http.get<BlogPost[]>(this.apiUrl));
  }

  async getBlogPostById(id: string): Promise<BlogPost> {
    return await firstValueFrom(this.http.get<BlogPost>(`${this.apiUrl}/${id}`));
  }

  createBlogPost(post: BlogPost): Promise<BlogPost> {
    return this.http.post<BlogPost>(this.apiUrl, post).toPromise() as Promise<BlogPost>;
  }

  async updateBlogPost(post: BlogPost): Promise<BlogPost> {
    return await firstValueFrom(this.http.put<BlogPost>(`${this.apiUrl}/${post.id}`, post));
  }

  deleteBlogPost(id: string): Promise<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).toPromise() as Promise<void>;
  }
}
