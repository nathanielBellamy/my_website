import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CsvService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = '/v1/api/admin/csv';

  exportCsv(entity: string) {
    return this.http.get(`${this.apiUrl}/${entity}`, {
      responseType: 'blob',
      observe: 'response'
    });
  }

  importCsv(entity: string, file: File) {
    const formData = new FormData();
    formData.append('file', file);
    return this.http.post(`${this.apiUrl}/${entity}`, formData);
  }
}
