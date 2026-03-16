import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { AboutService } from './about.service';
import { AboutContent } from '../models/about.model';
import { environment } from '../../environments/environment';

const mockAboutContent: AboutContent[] = [
  { id: '1', title: 'About Me', content: 'I am a software engineer.', order: 1 },
  { id: '2', title: 'My Hobbies', content: 'I like to code.', order: 2 },
];

describe('AboutService', () => {
  let service: AboutService;
  let httpMock: HttpTestingController;
  const API_URL = `${environment.BASE_URL}/api/marketing/about`;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [AboutService],
    });
    service = TestBed.inject(AboutService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  describe('getAll', () => {
    it('should return paginated about content', async () => {
      const promise = service.getAll(1, 10);
      const req = httpMock.expectOne(`${API_URL}?page=1&limit=10`);
      expect(req.request.method).toBe('GET');
      req.flush(mockAboutContent);
      const content = await promise;
      expect(content).toEqual(mockAboutContent);
    });
  });

  describe('getById', () => {
    it('should return a single content item', async () => {
      const content = mockAboutContent[0];
      const promise = service.getById('1');
      const req = httpMock.expectOne(`${API_URL}/1`);
      expect(req.request.method).toBe('GET');
      req.flush(content);
      const result = await promise;
      expect(result).toEqual(content);
    });
  });
});
