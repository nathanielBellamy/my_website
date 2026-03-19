import { TestBed } from '@angular/core/testing';
import { GrooveJrService } from './groove-jr.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { GrooveJrContent } from '../models/data-models';

describe('GrooveJrService', () => {
  let service: GrooveJrService;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [GrooveJrService]
    });
    service = TestBed.inject(GrooveJrService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should retrieve all GrooveJr content', async () => {
    const mockContent: GrooveJrContent[] = [{ id: '1', title: 'Test GrooveJr 1', content: 'Content 1' }];

    const promise = service.getAllGrooveJrContent();

    const req = httpTestingController.expectOne(req => req.url.startsWith('/v1/api/admin/groovejr'));
    expect(req.request.method).toEqual('GET');
    req.flush(mockContent);

    await expect(promise).resolves.toEqual(mockContent);
  });

  it('should retrieve GrooveJr content by ID', async () => {
    const mockContent: GrooveJrContent = { id: '1', title: 'Test GrooveJr 1', content: 'Content 1' };

    const promise = service.getGrooveJrContentById('1');

    const req = httpTestingController.expectOne('/v1/api/admin/groovejr/1');
    expect(req.request.method).toEqual('GET');
    req.flush(mockContent);

    await expect(promise).resolves.toEqual(mockContent);
  });

  it('should create GrooveJr content', async () => {
    const newContent: GrooveJrContent = { id: '', title: 'New GrooveJr', content: 'New Content' };
    const mockResponse: GrooveJrContent = { ...newContent, id: '2' };

    const promise = service.createGrooveJrContent(newContent);

    const req = httpTestingController.expectOne('/v1/api/admin/groovejr');
    expect(req.request.method).toEqual('POST');
    expect(req.request.body).toEqual(newContent);
    req.flush(mockResponse);

    await expect(promise).resolves.toEqual(mockResponse);
  });

  it('should update GrooveJr content', async () => {
    const updatedContent: GrooveJrContent = { id: '1', title: 'Updated GrooveJr', content: 'Updated Content' };
    const mockResponse: GrooveJrContent = { ...updatedContent };

    const promise = service.updateGrooveJrContent(updatedContent);

    const req = httpTestingController.expectOne('/v1/api/admin/groovejr/1');
    expect(req.request.method).toEqual('PUT');
    expect(req.request.body).toEqual(updatedContent);
    req.flush(mockResponse);

    await expect(promise).resolves.toEqual(mockResponse);
  });

  it('should delete GrooveJr content', async () => {
    const promise = service.deleteGrooveJrContent('1');

    const req = httpTestingController.expectOne('/v1/api/admin/groovejr/1');
    expect(req.request.method).toEqual('DELETE');
    req.flush(null);

    await expect(promise).resolves.toBeNull();
});
