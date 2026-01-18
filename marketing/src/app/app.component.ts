import { Component, OnInit } from '@angular/core';
import { NavbarComponent } from './components/navbar/navbar.component';
import { HeaderComponent } from './components/header/header.component';
import { RouterOutlet, Router, NavigationEnd } from '@angular/router';
import { filter } from 'rxjs/operators';
import { CommonModule } from '@angular/common'; // Import CommonModule for *ngIf

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [NavbarComponent, HeaderComponent, RouterOutlet, CommonModule], // Add CommonModule here
  template: `
    <app-navbar />
    <app-header *ngIf="showHeader"></app-header>
    <router-outlet />
  `,
  styles: [],
})
export class AppComponent implements OnInit {
  showHeader: boolean = true;

  constructor(private router: Router) {}

  ngOnInit() {
    this.router.events
      .pipe(filter(event => event instanceof NavigationEnd))
      .subscribe((event: NavigationEnd) => {
        this.showHeader = event.urlAfterRedirects === '/';
      });
  }
}
