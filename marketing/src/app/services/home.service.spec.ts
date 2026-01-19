import { HomeService } from './home.service';
import { HomeContent } from '../models/home.model';
import { environment } from '../../environments/environment';
import { PaginatedResponse } from '../models/pagination.model';

const mockHomeContent: HomeContent[] = [
  { id: 1, title: 'Welcome', body: 'Welcome to the site.' },
  { id: 2, title: 'About', body: 'This is a section about me.' },
];

const mockPaginatedResponse: PaginatedResponse<HomeContent> = {
  data: mockHomeContent,
  total: 2,
  page: 1,
  limit: 10,
};

describe('HomeService', () => {
  let service: HomeService;
  const API_URL = `${environment.API_BASE_URL}/home`;

  beforeEach(() => {
    service = new HomeService();
    global.fetch = jest.fn();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  describe('getAll', () => {
    it('should return paginated home content', async () => {
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
      await expect(service.getAll()).rejects.toThrow('Failed to fetch home content');
    });
  });

  describe('getById', () => {
    it('should return a single content item', async () => {
      const content = mockHomeContent[0];
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
      await expect(service.getById(1)).rejects.toThrow('Failed to fetch home content with id 1');
    });
  });
});
