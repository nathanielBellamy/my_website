import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { HomeService } from './home.service';
import { HomeContent } from '../models/home.model';
import { environment } from '../../environments/environment';

const mockHomeContent: HomeContent[] = [
  { id: 'home-1', title: 'Welcome', content: 'Welcome to the site.', order: 1 },
  { id: 'home-2', title: 'About', content: 'This is a section about me.', order: 2 },
];

describe('HomeService', () => {
  let service: HomeService;
  let httpMock: HttpTestingController;
  const API_URL = `${environment.API_BASE_URL}/marketing/home`;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [HomeService],
    });
    service = TestBed.inject(HomeService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  describe('getAll', () => {
    it('should return paginated home content', async () => {
      const promise = service.getAll(1, 10);
      const req = httpMock.expectOne(`${API_URL}?page=1&limit=10`);
      expect(req.request.method).toBe('GET');
      req.flush(mockHomeContent);
      const content = await promise;
      expect(content).toEqual(mockHomeContent);
    });
  });
});
