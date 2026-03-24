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

enum TimeRange {
  ONE_HOUR = '1h',
  SIX_HOURS = '6h',
  TWENTY_FOUR_HOURS = '24h',
  ALL = 'all',
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
  private copiedTimeout: ReturnType<typeof setTimeout> | null = null;

  protected readonly LogLevel = LogLevel;
  protected readonly ConnectionStatus = ConnectionStatus;
  protected readonly TimeRange = TimeRange;

  protected readonly entries = signal<LogEntry[]>([]);
  protected readonly connectionStatus = signal<ConnectionStatus>(ConnectionStatus.DISCONNECTED);
  protected readonly autoScroll = signal(true);
  protected readonly searchText = signal('');
  protected readonly selectedLevel = signal<string>('');
  protected readonly selectedTimeRange = signal<TimeRange>(TimeRange.ONE_HOUR);
  protected readonly isHistoryMode = signal(false);
  protected readonly selectedDate = signal('');
  protected readonly historyPage = signal(1);
  protected readonly historyTotal = signal(0);
  protected readonly expandedEntries = signal<Set<number>>(new Set());
  protected readonly copiedFeedback = signal(false);

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
    const timeRange = this.selectedTimeRange();
    const cutoff = this.getTimeCutoff(timeRange);

    return this.entries().filter((entry) => {
      if (level && entry.level?.toLowerCase() !== level.toLowerCase()) {
        return false;
      }
      if (cutoff && entry.time) {
        try {
          const entryTime = new Date(entry.time).getTime();
          if (entryTime < cutoff) {
            return false;
          }
        } catch {
          // keep entries with unparseable times
        }
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
    if (this.copiedTimeout) {
      clearTimeout(this.copiedTimeout);
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

  protected setTimeRange(range: TimeRange): void {
    this.selectedTimeRange.set(range);
  }

  protected clearEntries(): void {
    this.entries.set([]);
  }

  protected toggleEntryExpanded(index: number): void {
    this.expandedEntries.update((prev) => {
      const next = new Set(prev);
      if (next.has(index)) {
        next.delete(index);
      } else {
        next.add(index);
      }
      return next;
    });
  }

  protected isEntryExpanded(index: number): boolean {
    return this.expandedEntries().has(index);
  }

  protected fieldCount(fields: Record<string, string> | undefined): number {
    if (!fields) return 0;
    return Object.keys(fields).length;
  }

  protected copyLogsToClipboard(): void {
    const entries = this.filteredEntries();
    const csv = this.buildCsv(entries);
    navigator.clipboard.writeText(csv).then(() => {
      this.copiedFeedback.set(true);
      if (this.copiedTimeout) {
        clearTimeout(this.copiedTimeout);
      }
      this.copiedTimeout = setTimeout(() => {
        this.copiedFeedback.set(false);
      }, 2000);
    });
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

  private getTimeCutoff(range: TimeRange): number | null {
    const now = Date.now();
    switch (range) {
      case TimeRange.ONE_HOUR:
        return now - 60 * 60 * 1000;
      case TimeRange.SIX_HOURS:
        return now - 6 * 60 * 60 * 1000;
      case TimeRange.TWENTY_FOUR_HOURS:
        return now - 24 * 60 * 60 * 1000;
      case TimeRange.ALL:
        return null;
    }
  }

  private buildCsv(entries: LogEntry[]): string {
    // Collect all unique field keys across entries
    const fieldKeys = new Set<string>();
    for (const entry of entries) {
      if (entry.fields) {
        for (const key of Object.keys(entry.fields)) {
          fieldKeys.add(key);
        }
      }
    }
    const sortedFieldKeys = [...fieldKeys].sort();

    const headers = ['level', 'time', 'message', ...sortedFieldKeys];
    const rows = [headers.map(this.csvEscape).join(',')];

    for (const entry of entries) {
      const row = [
        entry.level ?? '',
        entry.time ?? '',
        entry.message ?? '',
        ...sortedFieldKeys.map((key) => {
          const val = entry.fields?.[key];
          return val != null ? String(val) : '';
        }),
      ];
      rows.push(row.map(this.csvEscape).join(','));
    }

    return rows.join('\n');
  }

  private csvEscape(value: string): string {
    if (value.includes(',') || value.includes('"') || value.includes('\n')) {
      return '"' + value.replace(/"/g, '""') + '"';
    }
    return value;
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
