import { render, screen, waitFor } from '@testing-library/angular';
import { DashboardComponent } from './dashboard.component';
import { LogsService } from '../../services/logs.service';
import { HealthInfo } from '../../models/log-models';

describe('DashboardComponent', () => {
  let mockLogsService: Partial<LogsService>;
  const mockHealth: HealthInfo = {
    uptime: '1d 2h 30m 0s',
    uptimeSeconds: 95400,
    goRoutines: 12,
    memAllocMb: 45.3,
    memSysMb: 120.5,
    numGc: 42,
    dbConnected: true,
    goVersion: 'go1.24.0',
    numCpu: 4,
  };

  beforeEach(() => {
    mockLogsService = {
      getHealth: jest.fn().mockReturnValue(Promise.resolve(mockHealth)),
      getLogHistory: jest.fn().mockReturnValue(Promise.resolve({ data: [], total: 0, page: 1, limit: 10 })),
    };
  });

  it('should create and show title', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });
    expect(screen.getByTestId('dashboard-title')).toBeInTheDocument();
    expect(screen.getByText('System Dashboard')).toBeInTheDocument();
  });

  it('should fetch health data on init', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(mockLogsService.getHealth).toHaveBeenCalled();

    await waitFor(() => {
      expect(screen.getByTestId('dashboard-uptime-card')).toBeInTheDocument();
      expect(screen.getByText('1d 2h 30m 0s')).toBeInTheDocument();
    });
  });

  it('should display memory info', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    await waitFor(() => {
      expect(screen.getByTestId('dashboard-memory-card')).toBeInTheDocument();
      expect(screen.getByText('45.3 MB')).toBeInTheDocument();
    });
  });

  it('should display goroutines count', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    await waitFor(() => {
      expect(screen.getByTestId('dashboard-goroutines-card')).toBeInTheDocument();
      expect(screen.getByText('12')).toBeInTheDocument();
    });
  });

  it('should display database connected status', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    await waitFor(() => {
      expect(screen.getByTestId('dashboard-db-card')).toBeInTheDocument();
      expect(screen.getByText('Connected')).toBeInTheDocument();
    });
  });

  it('should display system info', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    await waitFor(() => {
      expect(screen.getByTestId('dashboard-system-info')).toBeInTheDocument();
      expect(screen.getByText('go1.24.0')).toBeInTheDocument();
      expect(screen.getByText('4')).toBeInTheDocument();
    });
  });

  it('should show no recent errors message', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    await waitFor(() => {
      expect(screen.getByText(/No recent errors/)).toBeInTheDocument();
    });
  });

  it('should display recent errors when present', async () => {
    mockLogsService.getLogHistory = jest.fn().mockReturnValue(
      Promise.resolve({
        data: [
          { level: 'error', time: '2026-03-21T10:00:00Z', message: 'DB connection failed' },
        ],
        total: 1,
        page: 1,
        limit: 10,
      })
    );

    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    await waitFor(() => {
      expect(screen.getByText('DB connection failed')).toBeInTheDocument();
    });
  });

  it('should show quick links', async () => {
    await render(DashboardComponent, {
      providers: [{ provide: LogsService, useValue: mockLogsService }],
    });

    expect(screen.getByTestId('dashboard-link-logs')).toBeInTheDocument();
    expect(screen.getByTestId('dashboard-link-work')).toBeInTheDocument();
    expect(screen.getByTestId('dashboard-link-blog')).toBeInTheDocument();
  });
});
