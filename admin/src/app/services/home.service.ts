import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { HomeContent, FilterOptions, PaginatedResponse } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/api/admin/home'; // Adjust as per your backend URL

  async getAllHomeContent(options: Partial<FilterOptions> = {}): Promise<PaginatedResponse<HomeContent>> {
    const params: any = {
      page: options.page || 1,
      limit: options.limit || 10,
    };
    if (options.status) params.status = options.status;
    if (options.sortField) params.sort = options.sortField;
    if (options.sortOrder) params.order = options.sortOrder;

    return await firstValueFrom(this.http.get<PaginatedResponse<HomeContent>>(this.apiUrl, { params }));
  }

  async getHomeContentById(id: string): Promise<HomeContent> {
    return await firstValueFrom(this.http.get<HomeContent>(`${this.apiUrl}/${id}`));
  }

  async createHomeContent(content: HomeContent): Promise<HomeContent> {
    return await firstValueFrom(this.http.post<HomeContent>(this.apiUrl, content));
  }

  async updateHomeContent(content: HomeContent): Promise<HomeContent> {
    return await firstValueFrom(this.http.put<HomeContent>(`${this.apiUrl}/${content.id}`, content));
  }

  async deleteHomeContent(id: string): Promise<void> {
    return await firstValueFrom(this.http.delete<void>(`${this.apiUrl}/${id}`));
  }
}
