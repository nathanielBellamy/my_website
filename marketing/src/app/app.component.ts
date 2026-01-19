import { Component, OnInit, inject } from '@angular/core';
import { NavbarComponent } from './components/navbar/navbar.component';
import { HeaderComponent } from './components/header/header.component';
import { RouterOutlet, Router, NavigationEnd } from '@angular/router';
import { filter } from 'rxjs/operators';
import { CommonModule } from '@angular/common'; // Import CommonModule for *ngIf
import { TrackerService } from './services/tracker.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [NavbarComponent, HeaderComponent, RouterOutlet, CommonModule], // Add CommonModule here
  template: `
    <app-navbar />
    @if (showHeader) {
      <app-header></app-header>
    }
    <router-outlet />
  `,
  styles: [],
})
export class AppComponent implements OnInit {
  showHeader: boolean = true;
  private readonly trackerService = inject(TrackerService);

  constructor(private router: Router) {}

  ngOnInit() {
    this.trackerService.trackIp();
    this.router.events
      .pipe(filter(event => event instanceof NavigationEnd))
      .subscribe((event: NavigationEnd) => {
        this.showHeader = event.urlAfterRedirects === '/';
      });
  }
}
