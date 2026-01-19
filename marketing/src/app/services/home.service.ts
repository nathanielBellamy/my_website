import { Injectable } from '@angular/core';
import { HomeContent } from '../models/home.model';
import { environment } from '../../environments/environment';
import { PaginatedResponse } from '../models/pagination.model';

@Injectable({
  providedIn: 'root'
})
export class HomeService {
  private readonly API_URL = `${environment.API_BASE_URL}/home`;

  async getAll(page = 1, limit = 10): Promise<PaginatedResponse<HomeContent>> {
    const response = await fetch(`${this.API_URL}?page=${page}&limit=${limit}`);
    if (!response.ok) {
      throw new Error('Failed to fetch home content');
    }
    return response.json();
  }

  async getById(id: number): Promise<HomeContent> {
    const response = await fetch(`${this.API_URL}/${id}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch home content with id ${id}`);
    }
    return response.json();
  }
}
