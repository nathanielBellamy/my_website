import { Component } from '@angular/core';
import { NavbarComponent } from './components/navbar/navbar.component';
import { HeaderComponent } from './components/header/header.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [NavbarComponent, HeaderComponent],
  template: `
    <app-navbar />
    <app-header />
  `,
  styles: [],
})
export class AppComponent {
  title = 'marketing';
}
