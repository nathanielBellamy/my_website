import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BlogPost } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class BlogService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/admin/blog'; // Adjust as per your backend URL

  getAllBlogPosts(): Promise<BlogPost[]> {
    return this.http.get<BlogPost[]>(this.apiUrl).toPromise() as Promise<BlogPost[]>;
  }

  getBlogPostById(id: string): Promise<BlogPost> {
    return this.http.get<BlogPost>(`${this.apiUrl}/${id}`).toPromise() as Promise<BlogPost>;
  }

  createBlogPost(post: BlogPost): Promise<BlogPost> {
    return this.http.post<BlogPost>(this.apiUrl, post).toPromise() as Promise<BlogPost>;
  }

  updateBlogPost(post: BlogPost): Promise<BlogPost> {
    return this.http.put<BlogPost>(`${this.apiUrl}/${post.id}`, post).toPromise() as Promise<BlogPost>;
  }

  deleteBlogPost(id: string): Promise<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).toPromise() as Promise<void>;
  }
}
