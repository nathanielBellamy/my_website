import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GrooveJrContent } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class GrooveJrService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/admin/groovejr'; // Adjust as per your backend URL

  getAllGrooveJrContent(): Promise<GrooveJrContent[]> {
    return this.http.get<GrooveJrContent[]>(this.apiUrl).toPromise() as Promise<GrooveJrContent[]>;
  }

  getGrooveJrContentById(id: string): Promise<GrooveJrContent> {
    return this.http.get<GrooveJrContent>(`${this.apiUrl}/${id}`).toPromise() as Promise<GrooveJrContent>;
  }

  createGrooveJrContent(content: GrooveJrContent): Promise<GrooveJrContent> {
    return this.http.post<GrooveJrContent>(this.apiUrl, content).toPromise() as Promise<GrooveJrContent>;
  }

  updateGrooveJrContent(content: GrooveJrContent): Promise<GrooveJrContent> {
    return this.http.put<GrooveJrContent>(`${this.apiUrl}/${content.id}`, content).toPromise() as Promise<GrooveJrContent>;
  }

  deleteGrooveJrContent(id: string): Promise<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).toPromise() as Promise<void>;
  }
}
