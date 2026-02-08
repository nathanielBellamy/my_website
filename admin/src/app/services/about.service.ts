import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AboutContent } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class AboutService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/admin/about'; // Adjust as per your backend URL

  getAllAboutContent(): Promise<AboutContent[]> {
    return this.http.get<AboutContent[]>(this.apiUrl).toPromise() as Promise<AboutContent[]>;
  }

  getAboutContentById(id: string): Promise<AboutContent> {
    return this.http.get<AboutContent>(`${this.apiUrl}/${id}`).toPromise() as Promise<AboutContent>;
  }

  createAboutContent(content: AboutContent): Promise<AboutContent> {
    return this.http.post<AboutContent>(this.apiUrl, content).toPromise() as Promise<AboutContent>;
  }

  updateAboutContent(content: AboutContent): Promise<AboutContent> {
    return this.http.put<AboutContent>(`${this.apiUrl}/${content.id}`, content).toPromise() as Promise<AboutContent>;
  }

  deleteAboutContent(id: string): Promise<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).toPromise() as Promise<void>;
  }
}
