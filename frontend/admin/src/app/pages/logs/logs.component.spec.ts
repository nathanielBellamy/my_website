import { render, screen, fireEvent } from '@testing-library/angular';
import { LogsComponent } from './logs.component';
import { LogsService } from '../../services/logs.service';

describe('LogsComponent', () => {
  let mockLogsService: Partial<LogsService>;
  let mockEventSource: {
    addEventListener: jest.Mock;
    removeEventListener: jest.Mock;
    close: jest.Mock;
    onerror: ((event: Event) => void) | null;
    onopen: ((event: Event) => void) | null;
  };

  beforeEach(() => {
    mockEventSource = {
      addEventListener: jest.fn(),
      removeEventListener: jest.fn(),
      close: jest.fn(),
      onerror: null,
      onopen: null,
    };

    mockLogsService = {
      streamLogs: jest.fn().mockReturnValue(mockEventSource as unknown as EventSource),
      getLogHistory: jest.fn().mockReturnValue(
        Promise.resolve({ data: [], total: 0, page: 1, limit: 100 })
      ),
      getLogFiles: jest.fn().mockReturnValue(
        Promise.resolve({ files: [] })
      ),
    };
  });

  it('should create and show title', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });
    expect(screen.getByTestId('logs-page-title')).toBeInTheDocument();
    expect(screen.getByText('System Logs')).toBeInTheDocument();
  });

  it('should connect to SSE on init', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(mockLogsService.streamLogs).toHaveBeenCalledWith({ lines: 200 });
  });

  it('should register log and backfill-complete event listeners', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const eventNames = mockEventSource.addEventListener.mock.calls.map(
      (call: unknown[]) => call[0]
    );
    expect(eventNames).toContain('log');
    expect(eventNames).toContain('backfill-complete');
  });

  it('should show live mode controls', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(screen.getByTestId('logs-live-btn')).toBeInTheDocument();
    expect(screen.getByTestId('logs-history-btn')).toBeInTheDocument();
    expect(screen.getByTestId('logs-autoscroll-btn')).toBeInTheDocument();
    expect(screen.getByTestId('logs-clear-btn')).toBeInTheDocument();
  });

  it('should show level filter chips', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(screen.getByTestId('logs-level-chip-info')).toBeInTheDocument();
    expect(screen.getByTestId('logs-level-chip-warn')).toBeInTheDocument();
    expect(screen.getByTestId('logs-level-chip-error')).toBeInTheDocument();
    expect(screen.getByTestId('logs-level-chip-debug')).toBeInTheDocument();
    expect(screen.getByTestId('logs-level-chip-fatal')).toBeInTheDocument();
  });

  it('should show search input', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(screen.getByTestId('logs-search-input')).toBeInTheDocument();
  });

  it('should toggle auto-scroll on button click', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const btn = screen.getByTestId('logs-autoscroll-btn');
    expect(btn.textContent).toContain('ON');

    await fireEvent.click(btn);
    expect(btn.textContent).toContain('OFF');
  });

  it('should show empty state when no logs', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(screen.getByTestId('logs-container')).toBeInTheDocument();
  });

  it('should close EventSource on destroy', async () => {
    const { fixture } = await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    fixture.destroy();
    expect(mockEventSource.close).toHaveBeenCalled();
  });

  it('should switch to history mode when history button clicked', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const historyBtn = screen.getByTestId('logs-history-btn');
    await fireEvent.click(historyBtn);

    expect(mockLogsService.getLogHistory).toHaveBeenCalled();
    expect(mockEventSource.close).toHaveBeenCalled();
  });

  it('should show time range filter buttons with 1h selected by default', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(screen.getByTestId('logs-time-range-group')).toBeInTheDocument();
    expect(screen.getByTestId('logs-time-range-1h')).toBeInTheDocument();
    expect(screen.getByTestId('logs-time-range-6h')).toBeInTheDocument();
    expect(screen.getByTestId('logs-time-range-24h')).toBeInTheDocument();
    expect(screen.getByTestId('logs-time-range-all')).toBeInTheDocument();

    const oneHourBtn = screen.getByTestId('logs-time-range-1h');
    expect(oneHourBtn.getAttribute('aria-pressed')).toBe('true');
  });

  it('should toggle time range on click', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const allBtn = screen.getByTestId('logs-time-range-all');
    await fireEvent.click(allBtn);
    expect(allBtn.getAttribute('aria-pressed')).toBe('true');

    const oneHourBtn = screen.getByTestId('logs-time-range-1h');
    expect(oneHourBtn.getAttribute('aria-pressed')).toBe('false');
  });

  it('should show copy CSV button', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const copyBtn = screen.getByTestId('logs-copy-csv-btn');
    expect(copyBtn).toBeInTheDocument();
    expect(copyBtn.textContent).toContain('Copy CSV');
  });

  it('should copy CSV to clipboard when copy button clicked', async () => {
    const writeTextMock = jest.fn().mockResolvedValue(undefined);
    Object.assign(navigator, {
      clipboard: { writeText: writeTextMock },
    });

    const { fixture } = await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    // Simulate receiving log entries via SSE
    const logHandler = mockEventSource.addEventListener.mock.calls
      .find((call: unknown[]) => call[0] === 'log')?.[1] as (event: MessageEvent) => void;

    logHandler(new MessageEvent('log', {
      data: JSON.stringify({
        level: 'info',
        time: new Date().toISOString(),
        message: 'test message',
        fields: { path: '/api/test' },
      }),
    }));
    fixture.detectChanges();

    const copyBtn = screen.getByTestId('logs-copy-csv-btn');
    await fireEvent.click(copyBtn);

    expect(writeTextMock).toHaveBeenCalledTimes(1);
    const csv = writeTextMock.mock.calls[0][0] as string;
    expect(csv).toContain('level,time,message,path');
    expect(csv).toContain('info');
    expect(csv).toContain('test message');
    expect(csv).toContain('/api/test');
  });

  it('should expand entry fields on click and collapse on second click', async () => {
    const { fixture } = await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const logHandler = mockEventSource.addEventListener.mock.calls
      .find((call: unknown[]) => call[0] === 'log')?.[1] as (event: MessageEvent) => void;

    logHandler(new MessageEvent('log', {
      data: JSON.stringify({
        level: 'error',
        time: new Date().toISOString(),
        message: 'something failed',
        fields: { status: '500', path: '/api/health' },
      }),
    }));
    fixture.detectChanges();

    // Fields detail should not be visible initially
    expect(screen.queryByTestId('log-entry-fields-0')).not.toBeInTheDocument();

    // Click to expand
    const row = screen.getByTestId('log-entry-row-0');
    await fireEvent.click(row);
    expect(screen.getByTestId('log-entry-fields-0')).toBeInTheDocument();

    // Click again to collapse
    await fireEvent.click(row);
    expect(screen.queryByTestId('log-entry-fields-0')).not.toBeInTheDocument();
  });

  it('should show field count badge when entry has fields', async () => {
    const { fixture } = await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const logHandler = mockEventSource.addEventListener.mock.calls
      .find((call: unknown[]) => call[0] === 'log')?.[1] as (event: MessageEvent) => void;

    logHandler(new MessageEvent('log', {
      data: JSON.stringify({
        level: 'info',
        time: new Date().toISOString(),
        message: 'has fields',
        fields: { a: '1', b: '2', c: '3' },
      }),
    }));
    fixture.detectChanges();

    const badge = screen.getByTestId('log-entry-field-count-0');
    expect(badge).toBeInTheDocument();
    expect(badge.textContent).toContain('3 fields');
  });

  it('should filter entries by time range', async () => {
    const { fixture } = await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const logHandler = mockEventSource.addEventListener.mock.calls
      .find((call: unknown[]) => call[0] === 'log')?.[1] as (event: MessageEvent) => void;

    // Recent entry (should be visible with 1h filter)
    logHandler(new MessageEvent('log', {
      data: JSON.stringify({
        level: 'info',
        time: new Date().toISOString(),
        message: 'recent entry',
      }),
    }));

    // Old entry (should be filtered out with 1h filter)
    const oldDate = new Date();
    oldDate.setHours(oldDate.getHours() - 2);
    logHandler(new MessageEvent('log', {
      data: JSON.stringify({
        level: 'info',
        time: oldDate.toISOString(),
        message: 'old entry',
      }),
    }));
    fixture.detectChanges();

    // Default filter is 1h, so only recent entry should show
    expect(screen.getByText('recent entry')).toBeInTheDocument();
    expect(screen.queryByText('old entry')).not.toBeInTheDocument();

    // Switch to "All" time range, both should show
    const allBtn = screen.getByTestId('logs-time-range-all');
    await fireEvent.click(allBtn);

    expect(screen.getByText('recent entry')).toBeInTheDocument();
    expect(screen.getByText('old entry')).toBeInTheDocument();
  });

  it('should display CSV with proper escaping for fields containing commas', async () => {
    const writeTextMock = jest.fn().mockResolvedValue(undefined);
    Object.assign(navigator, {
      clipboard: { writeText: writeTextMock },
    });

    const { fixture } = await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    const logHandler = mockEventSource.addEventListener.mock.calls
      .find((call: unknown[]) => call[0] === 'log')?.[1] as (event: MessageEvent) => void;

    logHandler(new MessageEvent('log', {
      data: JSON.stringify({
        level: 'error',
        time: new Date().toISOString(),
        message: 'error, with comma',
        fields: { body: '{"key": "val"}' },
      }),
    }));
    fixture.detectChanges();

    // Switch to "All" so the entry is visible regardless of time
    const allBtn = screen.getByTestId('logs-time-range-all');
    await fireEvent.click(allBtn);

    const copyBtn = screen.getByTestId('logs-copy-csv-btn');
    await fireEvent.click(copyBtn);

    const csv = writeTextMock.mock.calls[0][0] as string;
    // Message with comma should be quoted
    expect(csv).toContain('"error, with comma"');
  });

  it('should indicate filtered state in footer when time range is not All', async () => {
    await render(LogsComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    // Default is 1h (not All), so "(filtered)" should appear
    expect(screen.getByText('(filtered)')).toBeInTheDocument();
  });
});
