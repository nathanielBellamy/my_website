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

  email = signal('');
  otp = signal('');
  step = signal<'email' | 'otp'>('email');
  error = signal<string | null>(null);
  loading = signal(false);

  async requestOtp() {
    this.loading.set(true);
    this.error.set(null);
    try {
      await this.authService.requestOtp(this.email());
      this.step.set('otp');
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
      await this.authService.verifyOtp(this.email(), this.otp());
      const returnTo = new URLSearchParams(window.location.search).get('return_to') || '/admin/';
      window.location.href = returnTo;
    } catch (err: any) {
      this.error.set(err.message || 'Invalid OTP');
    } finally {
      this.loading.set(false);
    }
  }
}
