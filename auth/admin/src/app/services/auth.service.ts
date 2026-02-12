import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly http = inject(HttpClient);

  async requestOtp(email: string): Promise<void> {
    await firstValueFrom(
      this.http.post<void>('/api/auth/admin/otp/request', { email })
    );
  }

  async verifyOtp(email: string, otp: string): Promise<void> {
    await firstValueFrom(
      this.http.post<void>('/api/auth/admin/otp/verify', { email, otp })
    );
  }
}
