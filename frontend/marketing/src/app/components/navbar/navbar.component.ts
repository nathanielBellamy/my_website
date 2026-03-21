import { Component, computed, inject, signal, WritableSignal } from '@angular/core';
import { Router, RouterLink, RouterLinkActive, NavigationEnd } from '@angular/router';
import { CommonModule } from '@angular/common';
import { toSignal } from '@angular/core/rxjs-interop';
import { filter, map } from 'rxjs/operators';

const PAGE_ROUTES = [
  '/',
  '/focus',
  '/work',
  '/about',
  '/groovejr',
  '/old-site-preview',
  '/blog'
];

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [RouterLink, RouterLinkActive, CommonModule],
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
  private readonly router = inject(Router);
  isOpen: WritableSignal<boolean> = signal(false);

  // Track current route URL
  private readonly currentUrl = toSignal(
    this.router.events.pipe(
      filter((event): event is NavigationEnd => event instanceof NavigationEnd),
      map(event => event.urlAfterRedirects.split('?')[0]) // ignore query params
    ),
    { initialValue: this.router.url.split('?')[0] }
  );

  // Compute the current index in our defined page sequence
  private readonly currentIndex = computed(() => {
    const url = this.currentUrl();
    // find index, exact match or match base path for blog sub-pages
    const index = PAGE_ROUTES.findIndex(route => 
      route === '/' ? url === '/' : url.startsWith(route)
    );
    return index;
  });

  readonly prevRoute = computed(() => {
    const idx = this.currentIndex();
    return PAGE_ROUTES[(idx + PAGE_ROUTES.length - 1) % PAGE_ROUTES.length];
  });

  readonly nextRoute = computed(() => {
    const idx = this.currentIndex();
    return PAGE_ROUTES[(idx + 1) % PAGE_ROUTES.length];
  });

  toggleMenu() {
    this.isOpen.update((value) => !value);
  }
}
