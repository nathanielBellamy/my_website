import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HomeContent } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/admin/home'; // Adjust as per your backend URL

  getAllHomeContent(): Promise<HomeContent[]> {
    return this.http.get<HomeContent[]>(this.apiUrl).toPromise() as Promise<HomeContent[]>;
  }

  getHomeContentById(id: string): Promise<HomeContent> {
    return this.http.get<HomeContent>(`${this.apiUrl}/${id}`).toPromise() as Promise<HomeContent>;
  }

  createHomeContent(content: HomeContent): Promise<HomeContent> {
    return this.http.post<HomeContent>(this.apiUrl, content).toPromise() as Promise<HomeContent>;
  }

  updateHomeContent(content: HomeContent): Promise<HomeContent> {
    return this.http.put<HomeContent>(`${this.apiUrl}/${content.id}`, content).toPromise() as Promise<HomeContent>;
  }

  deleteHomeContent(id: string): Promise<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).toPromise() as Promise<void>;
  }
}
