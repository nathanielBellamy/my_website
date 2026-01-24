import { HomeService } from './home.service';
import { HomeContent } from '../models/home.model';
import { environment } from '../../environments/environment.localhost';
import { PaginatedResponse } from '../models/pagination.model';

const mockHomeContent: HomeContent[] = [
  { id: 'home-1', title: 'Welcome', body: 'Welcome to the site.' },
  { id: 'home-2', title: 'About', body: 'This is a section about me.' },
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

      const content = await service.getAll(1, 10);
      expect(content).toEqual(mockPaginatedResponse);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}?page=1&limit=10`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getAll(1, 10)).rejects.toThrow('Failed to fetch home content');
    });
  });
});
