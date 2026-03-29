import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { WorkService } from './work.service';
import { WorkContent } from '../models/work.model';
import { environment } from '../../environments/environment';

const mockWorkContent: WorkContent[] = [
  { id: 'work-1', title: 'Welcome', content: 'Welcome to the site.', order: 1 },
  { id: 'work-2', title: 'About', content: 'This is a section about me.', order: 2 },
];

describe('WorkService', () => {
  let service: WorkService;
  let httpMock: HttpTestingController;
  const API_URL = `${environment.BASE_URL}/v1/api/marketing/work`;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [WorkService],
    });
    service = TestBed.inject(WorkService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  describe('getAll', () => {
    it('should return paginated work content', async () => {
      const promise = service.getAll(1, 10);
      const req = httpMock.expectOne(`${API_URL}?page=1&limit=10`);
      expect(req.request.method).toBe('GET');
      req.flush(mockWorkContent);
      const content = await promise;
      expect(content).toEqual(mockWorkContent);
    });
  });
});
