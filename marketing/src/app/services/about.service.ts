import { Injectable } from '@angular/core';
import { AboutContent } from '../models/about.model';
import { environment } from '../../environments/environment.localhost';
import { PaginatedResponse } from '../models/pagination.model';

@Injectable({
  providedIn: 'root'
})
export class AboutService {
  private readonly API_URL = `${environment.API_BASE_URL}/about`;

  async getAll(page = 1, limit = 10): Promise<PaginatedResponse<AboutContent>> {
    const response = await fetch(`${this.API_URL}?page=${page}&limit=${limit}`);
    if (!response.ok) {
      throw new Error('Failed to fetch about content');
    }
    return response.json();
  }

  async getById(id: number): Promise<AboutContent> {
    const response = await fetch(`${this.API_URL}/${id}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch about content with id ${id}`);
    }
    return response.json();
  }
}
