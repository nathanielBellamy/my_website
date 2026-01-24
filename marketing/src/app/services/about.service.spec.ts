import { AboutService } from './about.service';
import { AboutContent } from '../models/about.model';
import { environment } from '../../environments/environment';
import { PaginatedResponse } from '../models/pagination.model';
import { HttpClient } from '@angular/common/http';

const mockAboutContent: AboutContent[] = [
  { id: 1, title: 'About Me', body: 'I am a software engineer.' },
  { id: 2, title: 'My Hobbies', body: 'I like to code.' },
];

const mockPaginatedResponse: PaginatedResponse<AboutContent> = {
  data: mockAboutContent,
  total: 2,
  page: 1,
  limit: 10,
};

describe('AboutService', () => {
  let service: AboutService;
  const API_URL = `${environment.API_BASE_URL}/about`;

  beforeEach(() => {
    service = new AboutService();
    global.fetch = jest.fn();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  describe('getAll', () => {
    it('should return paginated about content', async () => {
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
      await expect(service.getAll(1, 10)).rejects.toThrow('Failed to fetch about content');
    });
  });

  describe('getById', () => {
    it('should return a single content item', async () => {
      const content = mockAboutContent[0];
      (global.fetch as jest.Mock).mockResolvedValue({
        ok: true,
        json: () => Promise.resolve(content),
      });

      const result = await service.getById('about-1');
      expect(result).toEqual(content);
      expect(fetch).toHaveBeenCalledWith(`${API_URL}/1`);
    });

    it('should throw an error if the request fails', async () => {
      (global.fetch as jest.Mock).mockResolvedValue({ ok: false });
      await expect(service.getById('myAbout')).rejects.toThrow('Failed to fetch about content.');
    });
  });
});
