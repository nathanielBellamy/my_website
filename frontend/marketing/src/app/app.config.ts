import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter, withRouterConfig, withInMemoryScrolling } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { MARKED_EXTENSIONS, provideMarkdown } from 'ngx-markdown';
import { markedImageResizeExtension } from './utils/markdown-image-renderer';
import { markedYouTubeExtension } from './utils/markdown-youtube-extension';

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
    provideMarkdown({
      markedExtensions: [
        {
          provide: MARKED_EXTENSIONS,
          useValue: markedImageResizeExtension,
          multi: true,
        },
        {
          provide: MARKED_EXTENSIONS,
          useValue: markedYouTubeExtension,
          multi: true,
        },
      ],
    }),
  ]
};
