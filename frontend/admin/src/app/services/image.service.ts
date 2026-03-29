import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { Image } from '../models/data-models';

@Injectable({
  providedIn: 'root',
})
export class ImageService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = '/v1/api/admin';

  async listImages(): Promise<Image[]> {
    return await firstValueFrom(this.http.get<Image[]>(`${this.apiUrl}/images`));
  }

  async uploadImage(file: File, altText: string): Promise<Image> {
    const formData = new FormData();
    formData.append('image', file);
    formData.append('altText', altText);

    return await firstValueFrom(this.http.post<Image>(`${this.apiUrl}/upload`, formData));
  }

  async deleteImage(id: string): Promise<void> {
    return await firstValueFrom(this.http.delete<void>(`${this.apiUrl}/images/${id}`));
  }
}
