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

  async getGrooveJrContentById(id: string): Promise<GrooveJrContent> {
    return await firstValueFrom(this.http.get<GrooveJrContent>(`${this.apiUrl}/${id}`));
  }

  async createGrooveJrContent(content: GrooveJrContent): Promise<GrooveJrContent> {
    return await firstValueFrom(this.http.post<GrooveJrContent>(this.apiUrl, content));
  }

  async updateGrooveJrContent(content: GrooveJrContent): Promise<GrooveJrContent> {
    return await firstValueFrom(this.http.put<GrooveJrContent>(`${this.apiUrl}/${content.id}`, content));
  }

  async deleteGrooveJrContent(id: string): Promise<void> {
    await firstValueFrom(this.http.delete<void>(`${this.apiUrl}/${id}`));
  }
}
