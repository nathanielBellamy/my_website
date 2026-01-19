import { Injectable } from '@angular/core';
import { GrooveJrContent } from '../models/groove-jr.model';
import { environment } from '../../environments/environment';
import { PaginatedResponse } from '../models/pagination.model';

@Injectable({
  providedIn: 'root'
})
export class GrooveJrService {
  private readonly API_URL = `${environment.API_BASE_URL}/groove-jr`;

  async getAll(page = 1, limit = 10): Promise<PaginatedResponse<GrooveJrContent>> {
    const response = await fetch(`${this.API_URL}?page=${page}&limit=${limit}`);
    if (!response.ok) {
      throw new Error('Failed to fetch groove-jr content');
    }
    return response.json();
  }

  async getById(id: number): Promise<GrooveJrContent> {
    const response = await fetch(`${this.API_URL}/${id}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch groove-jr content with id ${id}`);
    }
    return response.json();
  }
}
