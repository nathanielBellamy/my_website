import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { environment } from '../../environments/environment';
import { AboutContent } from '../models/about.model';

@Injectable({
  providedIn: 'root',
})
export class AboutService {
  private apiUrl = `${environment.API_BASE_URL}/marketing/about`;
  private readonly http = inject(HttpClient);

  constructor() {}

  getAll(page: number, limit: number): Promise<AboutContent[]> {
    return firstValueFrom(
      this.http.get<AboutContent[]>(`${this.apiUrl}?page=${page}&limit=${limit}`)
    );
  }

  getById(id: string): Promise<AboutContent[]> {
    return firstValueFrom(
      this.http.get<AboutContent[]>(`${this.apiUrl}/${id}`)
    );
  }
}