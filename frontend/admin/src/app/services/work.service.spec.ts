import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { WorkService } from './work.service';
import { WorkContent } from '../models/data-models';

describe('WorkService', () => {
  let service: WorkService;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [WorkService]
    });
    service = TestBed.inject(WorkService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should retrieve all work content', async () => {
    const mockContent: WorkContent[] = [{ id: '1', title: 'Test Work 1', content: 'Content 1', order: 1 }];

    const promise = service.getAllWorkContent();

    const req = httpTestingController.expectOne(req => req.url.startsWith('/v1/api/admin/work'));
    expect(req.request.method).toBe('GET');
    req.flush(mockContent);

    const result = await promise;
    expect(result).toEqual(mockContent);
  });

  it('should retrieve work content by ID', async () => {
    const mockContent: WorkContent = { id: '1', title: 'Test Work 1', content: 'Content 1', order: 1 };

    const promise = service.getWorkContentById('1');

    const req = httpTestingController.expectOne('/v1/api/admin/work/1');
    expect(req.request.method).toBe('GET');
    req.flush(mockContent);

    const result = await promise;
    expect(result).toEqual(mockContent);
  });

  it('should create work content', async () => {
    const newContent: WorkContent = { id: '', title: 'New Work', content: 'New Content', order: 1 };
    const mockResponse: WorkContent = { ...newContent, id: '2' };

    const promise = service.createWorkContent(newContent);

    const req = httpTestingController.expectOne('/v1/api/admin/work');
    expect(req.request.method).toBe('POST');
    req.flush(mockResponse);

    const result = await promise;
    expect(result).toEqual(mockResponse);
  });

  it('should update work content', async () => {
    const updatedContent: WorkContent = { id: '1', title: 'Updated Work', content: 'Updated Content', order: 1 };
    const mockResponse: WorkContent = { ...updatedContent };

    const promise = service.updateWorkContent(updatedContent);

    const req = httpTestingController.expectOne('/v1/api/admin/work/1');
    expect(req.request.method).toBe('PUT');
    req.flush(mockResponse);

    const result = await promise;
    expect(result).toEqual(mockResponse);
  });

  it('should delete work content', async () => {
    const promise = service.deleteWorkContent('1');

    const req = httpTestingController.expectOne('/v1/api/admin/work/1');
    expect(req.request.method).toBe('DELETE');
    req.flush(null);

    await promise;
  });
});
