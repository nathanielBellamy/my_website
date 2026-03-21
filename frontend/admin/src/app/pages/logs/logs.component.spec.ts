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
});
