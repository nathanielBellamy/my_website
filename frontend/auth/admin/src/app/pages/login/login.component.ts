import { Component, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  private readonly authService = inject(AuthService);

  password = signal('');
  otp = signal('');
  step = signal<'password' | 'request' | 'verify'>('password');
  error = signal<string | null>(null);
  loading = signal(false);

  async sha256(message: string): Promise<string> {
    const msgBuffer = new TextEncoder().encode(message);
    const hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
  }

  async submitPassword() {
    this.loading.set(true);
    this.error.set(null);
    try {
      const challenge = await this.authService.getChallenge();
      // SHA256(password + challenge)
      // Note: Backend expects ADMIN_PW + challenge.
      // So checking matches.
      const hash = await this.sha256(this.password() + challenge);
      await this.authService.validatePassword(hash);
      this.step.set('request');
    } catch (err: any) {
      console.error(err);
      this.error.set('Invalid Password');
    } finally {
      this.loading.set(false);
    }
  }

  async requestOtp() {
    this.loading.set(true);
    this.error.set(null);
    try {
      await this.authService.requestOtp();
      this.step.set('verify');
    } catch (err: any) {
      this.error.set(err.message || 'Failed to request OTP');
    } finally {
      this.loading.set(false);
    }
  }

  async login() {
    this.loading.set(true);
    this.error.set(null);
    try {
      await this.authService.verifyOtp(this.otp());
      const returnTo = new URLSearchParams(window.location.search).get('return_to') || '/';
      window.location.href = returnTo;
    } catch (err: any) {
      this.error.set(err.message || 'Invalid OTP');
    } finally {
      this.loading.set(false);
    }
  }
}
