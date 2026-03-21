import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
import { environment } from '../../environments/environment';
import { WorkContent } from '../models/work.model';

@Injectable({
  providedIn: 'root',
})
export class WorkService {
  private readonly apiUrl = `${environment.BASE_URL}/v1/api/marketing/work`;
  private readonly http = inject(HttpClient);

  getAll(page: number, limit: number): Promise<WorkContent[]> {
    return firstValueFrom(
      this.http.get<WorkContent[]>(
        `${this.apiUrl}?page=${page}&limit=${limit}`
      )
    );
  }

  getById(id: string): Promise<WorkContent> {
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
