import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { GrooveJrContent, FilterOptions, PaginatedResponse } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class GrooveJrService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/api/admin/groovejr'; // Adjust as per your backend URL

  async getAllGrooveJrContent(options: Partial<FilterOptions> = {}): Promise<PaginatedResponse<GrooveJrContent>> {
    const params: any = {
      page: options.page || 1,
      limit: options.limit || 10,
    };
    if (options.status) params.status = options.status;
    if (options.sortField) params.sort = options.sortField;
    if (options.sortOrder) params.order = options.sortOrder;

    return await firstValueFrom(this.http.get<PaginatedResponse<GrooveJrContent>>(this.apiUrl, { params }));
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
