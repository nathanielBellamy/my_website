import {
  Component,
  OnInit,
  OnDestroy,
  signal,
  computed,
  inject,
  ElementRef,
  ViewChild,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { LogsService } from '../../services/logs.service';
import { LogEntry, LogLevel } from '../../models/log-models';

enum ConnectionStatus {
  CONNECTED = 'connected',
  DISCONNECTED = 'disconnected',
  RECONNECTING = 'reconnecting',
}

@Component({
  selector: 'app-logs',
  imports: [CommonModule, FormsModule],
  templateUrl: './logs.component.html',
  styleUrl: './logs.component.css',
})
export class LogsComponent implements OnInit, OnDestroy {
  @ViewChild('logContainer') logContainer!: ElementRef<HTMLDivElement>;

  private readonly logsService: LogsService = inject(LogsService);
  private eventSource: EventSource | null = null;
  private reconnectTimeout: ReturnType<typeof setTimeout> | null = null;

  protected readonly LogLevel = LogLevel;
  protected readonly ConnectionStatus = ConnectionStatus;

  protected readonly entries = signal<LogEntry[]>([]);
  protected readonly connectionStatus = signal<ConnectionStatus>(ConnectionStatus.DISCONNECTED);
  protected readonly autoScroll = signal(true);
  protected readonly searchText = signal('');
  protected readonly selectedLevel = signal<string>('');
  protected readonly isHistoryMode = signal(false);
  protected readonly selectedDate = signal('');
  protected readonly historyPage = signal(1);
  protected readonly historyTotal = signal(0);

  protected readonly levelCounts = computed(() => {
    const counts: Record<string, number> = {
      info: 0,
      warn: 0,
      error: 0,
      debug: 0,
      fatal: 0,
    };
    for (const entry of this.entries()) {
      const lvl = entry.level?.toLowerCase() ?? 'unknown';
      if (lvl in counts) {
        counts[lvl]++;
      }
    }
    return counts;
  });

  protected readonly filteredEntries = computed(() => {
    const search = this.searchText().toLowerCase();
    const level = this.selectedLevel();
    return this.entries().filter((entry) => {
      if (level && entry.level?.toLowerCase() !== level.toLowerCase()) {
        return false;
      }
      if (search) {
        const text = `${entry.level} ${entry.time} ${entry.message} ${JSON.stringify(entry.fields ?? {})}`.toLowerCase();
        return text.includes(search);
      }
      return true;
    });
  });

  ngOnInit(): void {
    this.connectSSE();
  }

  ngOnDestroy(): void {
    this.disconnectSSE();
    if (this.reconnectTimeout) {
      clearTimeout(this.reconnectTimeout);
    }
  }

  protected connectSSE(): void {
    this.disconnectSSE();
    this.entries.set([]);
    this.connectionStatus.set(ConnectionStatus.RECONNECTING);

    this.eventSource = this.logsService.streamLogs({
      lines: 200,
    });

    this.eventSource.addEventListener('log', (event: MessageEvent) => {
      try {
        const entry: LogEntry = JSON.parse(event.data);
        this.entries.update((prev) => [...prev, entry]);
        if (this.autoScroll()) {
          this.scrollToBottom();
        }
      } catch {
        // raw text line, wrap it
        this.entries.update((prev) => [
          ...prev,
          { level: 'unknown', time: '', message: event.data },
        ]);
      }
    });

    this.eventSource.addEventListener('backfill-complete', () => {
      this.connectionStatus.set(ConnectionStatus.CONNECTED);
      if (this.autoScroll()) {
        this.scrollToBottom();
      }
    });

    this.eventSource.onerror = () => {
      this.connectionStatus.set(ConnectionStatus.DISCONNECTED);
      this.disconnectSSE();
      // Only auto-reconnect if we're still in live mode
      if (!this.isHistoryMode()) {
        this.reconnectTimeout = setTimeout(() => {
          this.connectionStatus.set(ConnectionStatus.RECONNECTING);
          this.connectSSE();
        }, 3000);
      }
    };

    this.eventSource.onopen = () => {
      this.connectionStatus.set(ConnectionStatus.CONNECTED);
    };
  }

  protected disconnectSSE(): void {
    if (this.reconnectTimeout) {
      clearTimeout(this.reconnectTimeout);
      this.reconnectTimeout = null;
    }
    if (this.eventSource) {
      this.eventSource.close();
      this.eventSource = null;
    }
  }

  protected toggleAutoScroll(): void {
    this.autoScroll.update((v) => !v);
    if (this.autoScroll()) {
      this.scrollToBottom();
    }
  }

  protected setLevelFilter(level: string): void {
    if (this.selectedLevel() === level) {
      this.selectedLevel.set('');
    } else {
      this.selectedLevel.set(level);
    }
  }

  protected clearEntries(): void {
    this.entries.set([]);
  }

  protected async loadHistory(): Promise<void> {
    this.disconnectSSE();
    this.isHistoryMode.set(true);
    this.connectionStatus.set(ConnectionStatus.DISCONNECTED);

    try {
      const result = await this.logsService.getLogHistory({
        page: this.historyPage(),
        limit: 100,
        date: this.selectedDate() || undefined,
        level: this.selectedLevel() || undefined,
        search: this.searchText() || undefined,
      });
      this.entries.set(result.data);
      this.historyTotal.set(result.total);
    } catch (err) {
      console.error('Error loading log history:', err);
    }
  }

  protected async historyPrev(): Promise<void> {
    if (this.historyPage() > 1) {
      this.historyPage.update((p) => p - 1);
      await this.loadHistory();
    }
  }

  protected async historyNext(): Promise<void> {
    const maxPage = Math.ceil(this.historyTotal() / 100);
    if (this.historyPage() < maxPage) {
      this.historyPage.update((p) => p + 1);
      await this.loadHistory();
    }
  }

  protected switchToLive(): void {
    this.isHistoryMode.set(false);
    this.historyPage.set(1);
    this.connectSSE();
  }

  protected getLevelClass(level: string): string {
    switch (level?.toLowerCase()) {
      case 'error':
      case 'fatal':
        return 'text-red-400';
      case 'warn':
        return 'text-yellow-400';
      case 'info':
        return 'text-blue-400';
      case 'debug':
        return 'text-gray-400';
      default:
        return 'text-gray-500';
    }
  }

  protected getLevelBgClass(level: string): string {
    switch (level?.toLowerCase()) {
      case 'error':
      case 'fatal':
        return 'bg-red-900/30 border-red-700';
      case 'warn':
        return 'bg-yellow-900/30 border-yellow-700';
      case 'info':
        return 'bg-blue-900/30 border-blue-700';
      case 'debug':
        return 'bg-gray-800/30 border-gray-600';
      default:
        return 'bg-gray-800/30 border-gray-600';
    }
  }

  protected formatTime(time: string): string {
    if (!time) return '';
    try {
      const d = new Date(time);
      return d.toLocaleTimeString('en-US', { hour12: false }) + '.' + d.getMilliseconds().toString().padStart(3, '0');
    } catch {
      return time;
    }
  }

  protected formatFields(fields: Record<string, string> | undefined): string {
    if (!fields || Object.keys(fields).length === 0) return '';
    return Object.entries(fields)
      .map(([k, v]) => `${k}=${v}`)
      .join(' ');
  }

  private scrollToBottom(): void {
    setTimeout(() => {
      if (this.logContainer?.nativeElement) {
        const el = this.logContainer.nativeElement;
        el.scrollTop = el.scrollHeight;
      }
    }, 0);
  }
}
