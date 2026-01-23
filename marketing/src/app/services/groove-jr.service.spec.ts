import { GrooveJrService } from './groove-jr.service';
import { GrooveJrContent } from '../models/groove-jr.model';
import { environment } from '../../environments/environment.localhost';
import { PaginatedResponse } from '../models/pagination.model';

const mockGrooveJrContent: GrooveJrContent[] = [
  { id: 1, title: 'Groove Jr.', body: 'A music app.', url: 'http://someurl.com' },
  { id: 2, title: 'Features', body: 'It has many features.', url: 'http://someurl.com/features' },
];

const mockPaginatedResponse: PaginatedResponse<GrooveJrContent> = {
  data: mockGrooveJrContent,
  total: 2,
  page: 1,
  limit: 10,
};

describe('GrooveJrService', () => {
  let service: GrooveJrService;
  const API_URL = `${environment.API_BASE_URL}/groove-jr`;

  beforeEach(() => {
    service = new GrooveJrService();
    global.fetch = jest.fn();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  describe('getAll', () => {
    it('should return paginated groove-jr content', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({
        ok: true,
        json: () => Promise.resolve(mockPaginatedResponse),
      });

      const content = await service.getAll();
      expect(content).toEqual(mockPaginatedResponse);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}?page=1&limit=10`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getAll()).rejects.toThrow('Failed to fetch groove-jr content');
    });
  });

  describe('getById', () => {
    it('should return a single content item', async () => {
      const content = mockGrooveJrContent[0];
      (global.fetch as jest.Mock).mockResolvedValue({
        ok: true,
        json: () => Promise.resolve(content),
      });

      const result = await service.getById(1);
      expect(result).toEqual(content);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}/1`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getById(1)).rejects.toThrow('Failed to fetch groove-jr content with id 1');
    });
  });
});
