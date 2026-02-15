import { TestBed } from '@angular/core/testing';
import { AboutService } from './about.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { AboutContent } from '../models/data-models';

describe('AboutService', () => {
  let service: AboutService;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [AboutService]
    });
    service = TestBed.inject(AboutService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should retrieve all About content', async () => {
    const mockContent: AboutContent[] = [{ id: '1', title: 'Test About 1', content: 'Content 1' }];

    const promise = service.getAllAboutContent();

    const req = httpTestingController.expectOne(req => req.url.startsWith('http://localhost:8080/api/admin/about'));
    expect(req.request.method).toEqual('GET');
    req.flush(mockContent);

    await expect(promise).resolves.toEqual(mockContent);
  });

  it('should retrieve About content by ID', async () => {
    const mockContent: AboutContent = { id: '1', title: 'Test About 1', content: 'Content 1' };

    const promise = service.getAboutContentById('1');

    const req = httpTestingController.expectOne('http://localhost:8080/api/admin/about/1');
    expect(req.request.method).toEqual('GET');
    req.flush(mockContent);

    await expect(promise).resolves.toEqual(mockContent);
  });

  it('should create About content', async () => {
    const newContent: AboutContent = { id: '', title: 'New About', content: 'New Content' };
    const mockResponse: AboutContent = { ...newContent, id: '2' };

    const promise = service.createAboutContent(newContent);

    const req = httpTestingController.expectOne('http://localhost:8080/api/admin/about');
    expect(req.request.method).toEqual('POST');
    expect(req.request.body).toEqual(newContent);
    req.flush(mockResponse);

    await expect(promise).resolves.toEqual(mockResponse);
  });

  it('should update About content', async () => {
    const updatedContent: AboutContent = { id: '1', title: 'Updated About', content: 'Updated Content' };
    const mockResponse: AboutContent = { ...updatedContent };

    const promise = service.updateAboutContent(updatedContent);

    const req = httpTestingController.expectOne('http://localhost:8080/api/admin/about/1');
    expect(req.request.method).toEqual('PUT');
    expect(req.request.body).toEqual(updatedContent);
    req.flush(mockResponse);

    await expect(promise).resolves.toEqual(mockResponse);
  });

  it('should delete About content', async () => {
    const promise = service.deleteAboutContent('1');

    const req = httpTestingController.expectOne('http://localhost:8080/api/admin/about/1');
    expect(req.request.method).toEqual('DELETE');
    req.flush(null);

    await expect(promise).resolves.toBeNull();
  });
});
