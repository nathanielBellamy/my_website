import { inject, Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import {
  HealthInfo,
  LogFilesResponse,
  PaginatedLogResponse,
} from '../models/log-models';

@Injectable({
  providedIn: 'root',
})
export class LogsService {
  private readonly http = inject(HttpClient);
  private readonly logsUrl = '/v1/api/admin/logs';
  private readonly healthUrl = '/v1/api/admin/health';

  streamLogs(options: { level?: string; search?: string; lines?: number } = {}): EventSource {
    const params = new URLSearchParams();
    if (options.level) params.set('level', options.level);
    if (options.search) params.set('search', options.search);
    if (options.lines) params.set('lines', options.lines.toString());

    const queryString = params.toString();
    const url = `${this.logsUrl}/stream${queryString ? '?' + queryString : ''}`;

    return new EventSource(url);
  }

  async getLogHistory(options: {
    page?: number;
    limit?: number;
    level?: string;
    search?: string;
    date?: string;
  } = {}): Promise<PaginatedLogResponse> {
    const params: Record<string, string | number> = {
      page: options.page ?? 1,
      limit: options.limit ?? 50,
    };
    if (options.level) params['level'] = options.level;
    if (options.search) params['search'] = options.search;
    if (options.date) params['date'] = options.date;

    return await firstValueFrom(
      this.http.get<PaginatedLogResponse>(`${this.logsUrl}/history`, { params })
    );
  }

  async getLogFiles(): Promise<LogFilesResponse> {
    return await firstValueFrom(
      this.http.get<LogFilesResponse>(`${this.logsUrl}/files`)
    );
  }

  async getHealth(): Promise<HealthInfo> {
    return await firstValueFrom(
      this.http.get<HealthInfo>(this.healthUrl)
    );
  }
}
