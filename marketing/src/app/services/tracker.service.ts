import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment.localhost';

@Injectable({
  providedIn: 'root'
})
export class TrackerService {
  private readonly API_URL = `${environment.API_BASE_URL}/tracker`;
  private readonly IP_TRACKED_KEY = 'ip_tracked';

  async trackIp() {
    if (this.isIpTracked()) {
      return;
    }

    try {
      const ip = await this.getIpAddress();
      await this.sendIpAddress(ip);
      this.setIpTracked();
    } catch (error) {
      console.error('Failed to track IP address:', error);
    }
  }

  private async getIpAddress(): Promise<string> {
    const response = await fetch('https://api.ipify.org?format=json');
    if (!response.ok) {
      throw new Error('Failed to fetch IP address');
    }
    const data = await response.json();
    return data.ip;
  }

  private async sendIpAddress(ip: string): Promise<void> {
    const response = await fetch(this.API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ ip }),
    });

    if (!response.ok) {
      throw new Error('Failed to send IP address to the backend');
    }
  }

  private isIpTracked(): boolean {
    return localStorage.getItem(this.IP_TRACKED_KEY) === 'true';
  }

  private setIpTracked(): void {
    localStorage.setItem(this.IP_TRACKED_KEY, 'true');
  }
}
