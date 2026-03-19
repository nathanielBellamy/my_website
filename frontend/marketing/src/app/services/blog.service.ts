import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
import { environment } from '../../environments/environment';
import { BlogPost, Tag } from '../models/blog-post.model';

@Injectable({
  providedIn: 'root',
})
export class BlogService {
  private readonly apiUrl = `${environment.BASE_URL}/v1/api/marketing/blog`;
  private readonly tagsUrl = `${environment.BASE_URL}/v1/api/marketing/tags`;
  private readonly http = inject(HttpClient);

  // TODO: add route to marketing controller in backend/go/marketing
  getById(id: string): Promise<BlogPost> {
    return firstValueFrom(
      this.http.get<any>(`${this.apiUrl}/${id}`).pipe(
        map((item) => ({
          id: item.id,
          title: item.title,
          content: item.content,
          order: item.order,
          author: item.author,
          tags: item.tags,
          createdAt: item.createdAt,
          updatedAt: item.updatedAt,
        }))
      )
    );
  }

  getAll(page: number, limit: number, tags?: string[]): Promise<BlogPost[]> {
    let url = `${this.apiUrl}?page=${page}&limit=${limit}`;
    if (tags && tags.length > 0) {
      url += `&tags=${tags.join(',')}`;
    }
    return firstValueFrom(
      this.http.get<BlogPost[]>(url)
    );
  }

  getTags(search: string = '', limit: number = 20): Promise<Tag[]> {
    let url = `${this.tagsUrl}?limit=${limit}`;
    if (search) {
      url += `&search=${search}`;
    }
    return firstValueFrom(this.http.get<Tag[]>(url));
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