import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { WorkContent, FilterOptions, PaginatedResponse } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = '/v1/api/admin/home'; // Adjust as per your backend URL

  async getAllWorkContent(options: Partial<FilterOptions> = {}): Promise<PaginatedResponse<WorkContent>> {
    const params: any = {
      page: options.page || 1,
      limit: options.limit || 10,
    };
    if (options.status) params.status = options.status;
    if (options.sortField) params.sort = options.sortField;
    if (options.sortOrder) params.order = options.sortOrder;

    return await firstValueFrom(this.http.get<PaginatedResponse<WorkContent>>(this.apiUrl, { params }));
  }

  async getWorkContentById(id: string): Promise<WorkContent> {
    return await firstValueFrom(this.http.get<WorkContent>(`${this.apiUrl}/${id}`));
  }

  async createWorkContent(content: WorkContent): Promise<WorkContent> {
    return await firstValueFrom(this.http.post<WorkContent>(this.apiUrl, content));
  }

  async updateWorkContent(content: WorkContent): Promise<WorkContent> {
    return await firstValueFrom(this.http.put<WorkContent>(`${this.apiUrl}/${content.id}`, content));
  }

  async deleteWorkContent(id: string): Promise<void> {
    return await firstValueFrom(this.http.delete<void>(`${this.apiUrl}/${id}`));
  }
}
