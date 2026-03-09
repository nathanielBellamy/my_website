import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
import { environment } from '../../environments/environment';
import { AboutContent } from '../models/about.model';

@Injectable({
  providedIn: 'root',
})
export class AboutService {
  private readonly apiUrl = `${environment.BASE_URL}/api/marketing/about`;
  private readonly http = inject(HttpClient);

  getAll(page: number, limit: number): Promise<AboutContent[]> {
    return firstValueFrom(
      this.http.get<any[]>(`${this.apiUrl}?page=${page}&limit=${limit}`).pipe(
        map((items) =>
          items.map((item) => ({
            id: item.id,
            title: item.title,
            content: item.content,
            order: item.order,
          }))
        )
      )
    );
  }

  getById(id: string): Promise<AboutContent> {
    return firstValueFrom(
      this.http.get<any>(`${this.apiUrl}/${id}`).pipe(
        map((item) => ({
          id: item.id,
          title: item.title,
          content: item.content,
          order: item.order,
        }))
      )
    );
  }
}