import { AuthService } from './auth.service';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

describe('AuthService', () => {
  let service: AuthService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        provideHttpClient(),
        provideHttpClientTesting(),
      ],
    });
    service = TestBed.inject(AuthService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('getChallenge', () => {
    it('should GET /api/auth/admin/challenge and return the challenge string', async () => {
      const promise = service.getChallenge();

      const req = httpMock.expectOne('/api/auth/admin/challenge');
      expect(req.request.method).toBe('GET');
      req.flush({ challenge: 'abc123' });

      const result = await promise;
      expect(result).toBe('abc123');
    });

    it('should reject when the request fails', async () => {
      const promise = service.getChallenge();

      const req = httpMock.expectOne('/api/auth/admin/challenge');
      req.flush('Server Error', { status: 500, statusText: 'Internal Server Error' });

      await expect(promise).rejects.toBeTruthy();
    });
  });

  describe('validatePassword', () => {
    it('should POST the hash to /api/auth/admin/password', async () => {
      const promise = service.validatePassword('hashed-pw');

      const req = httpMock.expectOne('/api/auth/admin/password');
      expect(req.request.method).toBe('POST');
      expect(req.request.body).toEqual({ hash: 'hashed-pw' });
      req.flush(null);

      await promise;
    });

    it('should reject when the password is invalid', async () => {
      const promise = service.validatePassword('wrong-hash');

      const req = httpMock.expectOne('/api/auth/admin/password');
      req.flush('Unauthorized', { status: 401, statusText: 'Unauthorized' });

      await expect(promise).rejects.toBeTruthy();
    });
  });

  describe('requestOtp', () => {
    it('should POST to /api/auth/admin/otp/request', async () => {
      const promise = service.requestOtp();

      const req = httpMock.expectOne('/api/auth/admin/otp/request');
      expect(req.request.method).toBe('POST');
      expect(req.request.body).toEqual({});
      req.flush(null);

      await promise;
    });

    it('should reject when OTP request fails', async () => {
      const promise = service.requestOtp();

      const req = httpMock.expectOne('/api/auth/admin/otp/request');
      req.flush('Error', { status: 500, statusText: 'Internal Server Error' });

      await expect(promise).rejects.toBeTruthy();
    });
  });

  describe('verifyOtp', () => {
    it('should POST the otp to /api/auth/admin/otp/verify', async () => {
      const promise = service.verifyOtp('123456');

      const req = httpMock.expectOne('/api/auth/admin/otp/verify');
      expect(req.request.method).toBe('POST');
      expect(req.request.body).toEqual({ otp: '123456' });
      req.flush(null);

      await promise;
    });

    it('should reject when OTP is invalid', async () => {
      const promise = service.verifyOtp('wrong');

      const req = httpMock.expectOne('/api/auth/admin/otp/verify');
      req.flush('Invalid OTP', { status: 401, statusText: 'Unauthorized' });

      await expect(promise).rejects.toBeTruthy();
    });
  });
});
