import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { GrooveJrService } from './groove-jr.service';
import { GrooveJrContent } from '../models/groove-jr.model';
import { environment } from '../../environments/environment';

const mockGrooveJrContent: GrooveJrContent[] = [
  { id: '1', title: 'Groove Jr.', content: 'A music app.' },
  { id: '2', title: 'Features', content: 'It has many features.' },
];

describe('GrooveJrService', () => {
  let service: GrooveJrService;
  let httpMock: HttpTestingController;
  const API_URL = `${environment.API_BASE_URL}/marketing/groovejr`;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [GrooveJrService],
    });
    service = TestBed.inject(GrooveJrService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  describe('getAll', () => {
    it('should return paginated groove-jr content', async () => {
      const promise = service.getAll(1, 10);
      const req = httpMock.expectOne(`${API_URL}?page=1&limit=10`);
      expect(req.request.method).toBe('GET');
      req.flush(mockGrooveJrContent);
      const content = await promise;
      expect(content).toEqual(mockGrooveJrContent);
    });
  });

  describe('getById', () => {
    it('should return a single content item', async () => {
      const content = mockGrooveJrContent[0];
      const promise = service.getById('1');
      const req = httpMock.expectOne(`${API_URL}/1`);
      expect(req.request.method).toBe('GET');
      req.flush(content);
      const result = await promise;
      expect(result).toEqual(content);
    });
  });
});
