import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
import { environment } from '../../environments/environment';
import { HomeContent } from '../models/home.model';

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  private readonly apiUrl = `${environment.BASE_URL_API}/marketing/home`;
  private readonly http = inject(HttpClient);

  getAll(page: number, limit: number): Promise<HomeContent[]> {
    return firstValueFrom(
      this.http.get<HomeContent[]>(
        `${this.apiUrl}?page=${page}&limit=${limit}`
      )
    );
  }
}