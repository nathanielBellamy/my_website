import { TestBed } from '@angular/core/testing';
import { HomeService } from './home.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HomeContent } from '../models/data-models';

describe('HomeService', () => {
  let service: HomeService;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [HomeService]
    });
    service = TestBed.inject(HomeService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify(); // Ensure that there are no outstanding requests.
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should retrieve all home content', async () => {
    const mockContent: HomeContent[] = [{ id: '1', title: 'Test Home 1', content: 'Content 1' }];

    // Trigger the service method
    const promise = service.getAllHomeContent();

    // Expect a GET request to the API URL
    const req = httpTestingController.expectOne(req => req.url.startsWith('/api/admin/home'));
    expect(req.request.method).toEqual('GET');

    // Respond to the request with mock data
    req.flush(mockContent);

    // Await the promise and assert the result
    await expect(promise).resolves.toEqual(mockContent);
  });

  it('should retrieve home content by ID', async () => {
    const mockContent: HomeContent = { id: '1', title: 'Test Home 1', content: 'Content 1' };

    const promise = service.getHomeContentById('1');

    const req = httpTestingController.expectOne('/api/admin/home/1');
    expect(req.request.method).toEqual('GET');
    req.flush(mockContent);

    await expect(promise).resolves.toEqual(mockContent);
  });

  it('should create home content', async () => {
    const newContent: HomeContent = { id: '', title: 'New Home', content: 'New Content' };
    const mockResponse: HomeContent = { ...newContent, id: '2' };

    const promise = service.createHomeContent(newContent);

    const req = httpTestingController.expectOne('/api/admin/home');
    expect(req.request.method).toEqual('POST');
    expect(req.request.body).toEqual(newContent);
    req.flush(mockResponse);

    await expect(promise).resolves.toEqual(mockResponse);
  });

  it('should update home content', async () => {
    const updatedContent: HomeContent = { id: '1', title: 'Updated Home', content: 'Updated Content' };
    const mockResponse: HomeContent = { ...updatedContent };

    const promise = service.updateHomeContent(updatedContent);

    const req = httpTestingController.expectOne('/api/admin/home/1');
    expect(req.request.method).toEqual('PUT');
    expect(req.request.body).toEqual(updatedContent);
    req.flush(mockResponse);

    await expect(promise).resolves.toEqual(mockResponse);
  });

  it('should delete home content', async () => {
    const promise = service.deleteHomeContent('1');

    const req = httpTestingController.expectOne('/api/admin/home/1');
    expect(req.request.method).toEqual('DELETE');
    req.flush(null);

    await expect(promise).resolves.toBeNull();
});
