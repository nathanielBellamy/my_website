import { Component, signal, WritableSignal } from '@angular/core';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [RouterLink, CommonModule],
  templateUrl: './navbar.component.html',
  styles: [
    `
      :host {
        display: block;
        position: sticky;
        top: 0;
        z-index: 50;
      }
    `,
  ],
})
export class NavbarComponent {
  isOpen: WritableSignal<boolean> = signal(false);

  toggleMenu() {
    this.isOpen.update((value) => !value);
  }
}
