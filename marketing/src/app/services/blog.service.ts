import { Injectable } from '@angular/core';
import { BlogPost } from '../models/blog.model';
import { environment } from '../../environments/environment';
import { PaginatedResponse } from '../models/pagination.model';

@Injectable({
  providedIn: 'root'
})
export class BlogService {
  private readonly API_URL = `${environment.API_BASE_URL}/blog`;

  async getAll(page = 1, limit = 10): Promise<PaginatedResponse<BlogPost>> {
    const response = await fetch(`${this.API_URL}?page=${page}&limit=${limit}`);
    if (!response.ok) {
      throw new Error('Failed to fetch blog posts');
    }
    return response.json();
  }

  async getById(id: number): Promise<BlogPost> {
    const response = await fetch(`${this.API_URL}/${id}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch blog post with id ${id}`);
    }
    return response.json();
  }

  async getByTag(tag: string, page = 1, limit = 10): Promise<PaginatedResponse<BlogPost>> {
    const response = await fetch(`${this.API_URL}/tag/${tag}?page=${page}&limit=${limit}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch blog posts with tag ${tag}`);
    }
    return response.json();
  }

  async getByDate(date: string, page = 1, limit = 10): Promise<PaginatedResponse<BlogPost>> {
    const response = await fetch(`${this.API_URL}/date/${date}?page=${page}&limit=${limit}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch blog posts with date ${date}`);
    }
    return response.json();
  }
}
