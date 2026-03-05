import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { MARKED_EXTENSIONS, provideMarkdown } from 'ngx-markdown';
import { markedImageResizeExtension } from './utils/markdown-image-renderer';

import { routes } from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideHttpClient(),
    provideMarkdown({
      markedExtensions: [
        {
          provide: MARKED_EXTENSIONS,
          useValue: markedImageResizeExtension,
          multi: true,
        },
      ],
    }),
  ]
};
