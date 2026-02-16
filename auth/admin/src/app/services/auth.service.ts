import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly http = inject(HttpClient);

  async getChallenge(): Promise<string> {
    const res = await firstValueFrom(
      this.http.get<{ challenge: string }>('/api/auth/admin/challenge')
    );
    return res.challenge;
  }

  async validatePassword(hash: string): Promise<void> {
    await firstValueFrom(
      this.http.post<void>('/api/auth/admin/password', { hash })
    );
  }

  async requestOtp(): Promise<void> {
    await firstValueFrom(
      this.http.post<void>('/api/auth/admin/otp/request', {})
    );
  }

  async verifyOtp(otp: string): Promise<void> {
    await firstValueFrom(
      this.http.post<void>('/api/auth/admin/otp/verify', { otp })
    );
  }
}
