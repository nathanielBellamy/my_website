import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
import { environment } from '../../environments/environment';
import { HomeContent } from '../models/home.model';

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  private apiUrl = `${environment.API_BASE_URL}/marketing/home`;
  private readonly http = inject(HttpClient);

  constructor() {}

  getAll(page: number, limit: number): Promise<HomeContent[]> {
    return firstValueFrom(
      this.http.get<any[]>(`${this.apiUrl}?page=${page}&limit=${limit}`).pipe(
        map((items) =>
          items.map((item) => ({
            id: item.id,
            title: item.title,
            content: item.content,
          }))
        )
      )
    );
  }
}