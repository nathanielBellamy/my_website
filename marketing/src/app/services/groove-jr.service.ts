import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { GrooveJrContent } from '../models/groove-jr.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class GrooveJrService {
  private apiUrl = `${environment.API_BASE_URL}/marketing/groovejr`;
  private readonly http = inject(HttpClient);

  constructor() {}

  getAll(page: number, limit: number): Promise<GrooveJrContent[]> {
    return firstValueFrom(
      this.http.get<GrooveJrContent[]>(`${this.apiUrl}?page=${page}&limit=${limit}`)
    );
  }

  getById(id: string): Promise<GrooveJrContent[]> {
    return firstValueFrom(
      this.http.get<GrooveJrContent[]>(`${this.apiUrl}/${id}`)
    );
  }
}