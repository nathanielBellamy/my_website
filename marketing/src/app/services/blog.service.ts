import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
import { environment } from '../../environments/environment';
import { BlogPost } from '../models/blog-post.model';

@Injectable({
  providedIn: 'root',
})
export class BlogService {
  private readonly apiUrl = `${environment.API_BASE_URL}/marketing/blog`;
  private readonly http = inject(HttpClient);

  constructor() {}

  // TODO: add route to marketing controller in backend/go/marketing
  getById(id: string): Promise<BlogPost> {
    return firstValueFrom(
      this.http.get<any>(`${this.apiUrl}/${id}`).pipe(
        map((item) => ({
          id: item.id,
          title: item.title,
          content: item.content,
          author: item.author,
          tags: item.tags,
          createdAt: item.createdAt,
          updatedAt: item.updatedAt,
        }))
      )
    );
  }

  getAll(page: number, limit: number): Promise<BlogPost[]> {
    return firstValueFrom(
      this.http.get<BlogPost[]>(`${this.apiUrl}?page=${page}&limit=${limit}`)
    );
  }

  // TODO: add route to marketing controller in backend/go/marketing
  getByTag(tag: string): Promise<BlogPost[]> {
    return firstValueFrom(
      this.http.get<BlogPost[]>(`${this.apiUrl}?tag=${tag}`)
    );
  }
  
  // TODO: add route to marketing controller in backend/go/marketing
  getByDate(start: Date, end: Date): Promise<BlogPost[]> {
    return firstValueFrom(
      this.http.get<BlogPost[]>(`${this.apiUrl}?start=${start.toUTCString()}&end=${end}`)
    );
  }
}