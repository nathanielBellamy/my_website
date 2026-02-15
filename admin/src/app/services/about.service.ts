import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { AboutContent, FilterOptions, PaginatedResponse } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class AboutService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/api/admin/about'; // Adjust as per your backend URL

  async getAllAboutContent(options: Partial<FilterOptions> = {}): Promise<PaginatedResponse<AboutContent>> {
    const params: any = {
      page: options.page || 1,
      limit: options.limit || 10,
    };
    if (options.status) params.status = options.status;
    if (options.sortField) params.sort = options.sortField;
    if (options.sortOrder) params.order = options.sortOrder;

    return await firstValueFrom(this.http.get<PaginatedResponse<AboutContent>>(this.apiUrl, { params }));
  }

  async getAboutContentById(id: string): Promise<AboutContent> {
    return await firstValueFrom(this.http.get<AboutContent>(`${this.apiUrl}/${id}`));
  }

  async createAboutContent(content: AboutContent): Promise<AboutContent> {
    return await firstValueFrom(this.http.post<AboutContent>(this.apiUrl, content));
  }

  async updateAboutContent(content: AboutContent): Promise<AboutContent> {
    return await firstValueFrom(this.http.put<AboutContent>(`${this.apiUrl}/${content.id}`, content));
  }

  async deleteAboutContent(id: string): Promise<void> {
    await firstValueFrom(this.http.delete<void>(`${this.apiUrl}/${id}`));
  }
}
