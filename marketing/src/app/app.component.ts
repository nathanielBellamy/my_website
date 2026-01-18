import { Component } from '@angular/core';
import { NavbarComponent } from './components/navbar/navbar.component';
import { HeaderComponent } from './components/header/header.component';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [NavbarComponent, HeaderComponent, RouterOutlet],
  template: `
    <app-navbar />
    <app-header />
    <router-outlet />
  `,
  styles: [],
})
export class AppComponent {
  title = 'marketing';
}
