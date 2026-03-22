import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { LogsService } from './logs.service';
import { HealthInfo, PaginatedLogResponse, LogFilesResponse } from '../models/log-models';

describe('LogsService', () => {
  let service: LogsService;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [LogsService],
    });
    service = TestBed.inject(LogsService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should retrieve health info', async () => {
    const mockHealth: HealthInfo = {
      uptime: '2h 30m 0s',
      uptimeSeconds: 9000,
      goRoutines: 8,
      memAllocMb: 32.5,
      memSysMb: 100.0,
      numGc: 20,
      dbConnected: true,
      goVersion: 'go1.24.0',
      numCpu: 4,
    };

    const promise = service.getHealth();

    const req = httpTestingController.expectOne('/v1/api/admin/health');
    expect(req.request.method).toBe('GET');
    req.flush(mockHealth);

    const result = await promise;
    expect(result).toEqual(mockHealth);
  });

  it('should retrieve log history with default params', async () => {
    const mockResponse: PaginatedLogResponse = {
      data: [
        { level: 'info', time: '2026-03-21T10:00:00Z', message: 'Server started' },
      ],
      total: 1,
      page: 1,
      limit: 50,
    };

    const promise = service.getLogHistory();

    const req = httpTestingController.expectOne(
      (r) => r.url === '/v1/api/admin/logs/history' && r.params.get('page') === '1' && r.params.get('limit') === '50'
    );
    expect(req.request.method).toBe('GET');
    req.flush(mockResponse);

    const result = await promise;
    expect(result).toEqual(mockResponse);
  });

  it('should retrieve log history with filters', async () => {
    const mockResponse: PaginatedLogResponse = {
      data: [],
      total: 0,
      page: 2,
      limit: 100,
    };

    const promise = service.getLogHistory({
      page: 2,
      limit: 100,
      level: 'error',
      search: 'timeout',
      date: '2026-03-21',
    });

    const req = httpTestingController.expectOne((r) => {
      return r.url === '/v1/api/admin/logs/history'
        && r.params.get('page') === '2'
        && r.params.get('limit') === '100'
        && r.params.get('level') === 'error'
        && r.params.get('search') === 'timeout'
        && r.params.get('date') === '2026-03-21';
    });
    expect(req.request.method).toBe('GET');
    req.flush(mockResponse);

    const result = await promise;
    expect(result).toEqual(mockResponse);
  });

  it('should retrieve log files', async () => {
    const mockResponse: LogFilesResponse = {
      files: [
        { path: '2026/03/2026-03-21T10-00-00Z-log.txt', date: '2026-03-21', size: 1024 },
      ],
    };

    const promise = service.getLogFiles();

    const req = httpTestingController.expectOne('/v1/api/admin/logs/files');
    expect(req.request.method).toBe('GET');
    req.flush(mockResponse);

    const result = await promise;
    expect(result).toEqual(mockResponse);
  });

  it('should create EventSource with correct URL for streamLogs', () => {
    // streamLogs creates an EventSource directly (not HTTP), so we test URL construction
    // We mock EventSource globally for this test
    const mockES = { close: jest.fn() } as unknown as EventSource;
    const originalEventSource = globalThis.EventSource;
    let capturedUrl = '';
    globalThis.EventSource = jest.fn((url: string) => {
      capturedUrl = url;
      return mockES;
    }) as unknown as typeof EventSource;

    const es = service.streamLogs({ level: 'error', search: 'db', lines: 50 });

    expect(capturedUrl).toBe('/v1/api/admin/logs/stream?level=error&search=db&lines=50');
    expect(es).toBe(mockES);

    globalThis.EventSource = originalEventSource;
  });

  it('should create EventSource with no query params when options are empty', () => {
    const mockES = { close: jest.fn() } as unknown as EventSource;
    const originalEventSource = globalThis.EventSource;
    let capturedUrl = '';
    globalThis.EventSource = jest.fn((url: string) => {
      capturedUrl = url;
      return mockES;
    }) as unknown as typeof EventSource;

    const es = service.streamLogs();

    expect(capturedUrl).toBe('/v1/api/admin/logs/stream');
    expect(es).toBe(mockES);

    globalThis.EventSource = originalEventSource;
  });
});
