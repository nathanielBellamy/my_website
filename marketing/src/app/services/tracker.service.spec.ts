import { TrackerService } from './tracker.service';
import { environment } from '../../environments/environment.localhost';

describe('TrackerService', () => {
  let service: TrackerService;
  const API_URL = `${environment.API_BASE_URL}/tracker`;

  beforeEach(() => {
    service = new TrackerService();
    global.fetch = jest.fn();
    localStorage.clear();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('should track ip address if not already tracked', async () => {
    (global.fetch as jest.Mock)
      .mockResolvedValueOnce({
        ok: true,
        json: () => Promise.resolve({ ip: '127.0.0.1' }),
      })
      .mockResolvedValueOnce({
        ok: true,
      });

    await service.trackIp();

    expect(fetch).toHaveBeenCalledWith('https://api.ipify.org?format=json');
    expect(fetch).toHaveBeenCalledWith(API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ ip: '127.0.0.1' }),
    });
    expect(localStorage.getItem('ip_tracked')).toBe('true');
  });

  it('should not track ip address if already tracked', async () => {
    localStorage.setItem('ip_tracked', 'true');
    await service.trackIp();
    expect(fetch).not.toHaveBeenCalled();
  });

  it('should handle failure to get ip address', async () => {
    (global.fetch as jest.Mock).mockResolvedValueOnce({ ok: false });
    await service.trackIp();
    expect(localStorage.getItem('ip_tracked')).toBeNull();
  });

  it('should handle failure to send ip address', async () => {
    (global.fetch as jest.Mock)
      .mockResolvedValueOnce({
        ok: true,
        json: () => Promise.resolve({ ip: '127.0.0.1' }),
      })
      .mockResolvedValueOnce({
        ok: false,
      });

    await service.trackIp();
    expect(localStorage.getItem('ip_tracked')).toBeNull();
  });
});
