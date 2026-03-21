import {
  Component,
  OnInit,
  OnDestroy,
  signal,
  inject,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { LogsService } from '../../services/logs.service';
import { HealthInfo, LogEntry } from '../../models/log-models';

@Component({
  selector: 'app-dashboard',
  imports: [CommonModule, RouterLink],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
})
export class DashboardComponent implements OnInit, OnDestroy {
  private readonly logsService: LogsService = inject(LogsService);
  private refreshInterval: ReturnType<typeof setInterval> | null = null;

  protected readonly health = signal<HealthInfo | null>(null);
  protected readonly recentErrors = signal<LogEntry[]>([]);
  protected readonly loading = signal(true);
  protected readonly lastRefresh = signal<Date | null>(null);

  ngOnInit(): void {
    this.refresh();
    this.refreshInterval = setInterval(() => this.refresh(), 15000);
  }

  ngOnDestroy(): void {
    if (this.refreshInterval) {
      clearInterval(this.refreshInterval);
    }
  }

  protected async refresh(): Promise<void> {
    this.loading.set(true);

    try {
      const [healthData, errorLogs] = await Promise.all([
        this.logsService.getHealth(),
        this.logsService.getLogHistory({ limit: 10, level: 'error' }),
      ]);

      this.health.set(healthData);
      this.recentErrors.set(errorLogs.data);
      this.lastRefresh.set(new Date());
    } catch (err) {
      console.error('Error loading dashboard data:', err);
    } finally {
      this.loading.set(false);
    }
  }

  protected formatMb(mb: number): string {
    if (mb >= 1024) {
      return (mb / 1024).toFixed(1) + ' GB';
    }
    return mb.toFixed(1) + ' MB';
  }

  protected formatTime(time: string): string {
    if (!time) return '';
    try {
      return new Date(time).toLocaleString();
    } catch {
      return time;
    }
  }
}
