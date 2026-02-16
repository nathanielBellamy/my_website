import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
import { GrooveJrContent } from '../models/groove-jr.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class GrooveJrService {
  private readonly apiUrl = `${environment.API_BASE_URL}/marketing/groovejr`;
  private readonly http = inject(HttpClient);

  getAll(page: number, limit: number): Promise<GrooveJrContent[]> {
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

  getById(id: string): Promise<GrooveJrContent> {
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