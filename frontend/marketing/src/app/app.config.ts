import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter, withRouterConfig, withInMemoryScrolling } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideMarkdown } from 'ngx-markdown';

import { routes } from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(
      routes,
      withRouterConfig({ onSameUrlNavigation: 'reload' }),
      withInMemoryScrolling({ scrollPositionRestoration: 'enabled' })
    ),
    provideHttpClient(),
    provideAnimationsAsync(),
    provideMarkdown()
  ]
};
