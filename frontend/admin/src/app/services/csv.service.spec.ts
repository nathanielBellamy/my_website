import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { CsvService } from './csv.service';

describe('CsvService', () => {
  let service: CsvService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [CsvService]
    });
    service = TestBed.inject(CsvService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should export csv', () => {
    const entity = 'blog';
    const mockBlob = new Blob(['title,content'], { type: 'text/csv' });

    service.exportCsv(entity).subscribe(response => {
      expect(response.body).toEqual(mockBlob);
    });

    const req = httpMock.expectOne(`/v1/api/admin/csv/${entity}`);
    expect(req.request.method).toBe('GET');
    expect(req.request.responseType).toBe('blob');
    req.flush(mockBlob);
  });

  it('should import csv', () => {
    const entity = 'blog';
    const file = new File(['title,content'], 'test.csv', { type: 'text/csv' });

    service.importCsv(entity, file).subscribe(response => {
      expect(response).toBeTruthy();
    });

    const req = httpMock.expectOne(`/v1/api/admin/csv/${entity}`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body instanceof FormData).toBe(true);
    expect(req.request.body.has('file')).toBe(true);
    req.flush({});
  });
});
