import { Routes } from '@angular/router';
import { AllSectionsComponent } from './pages/all-sections/all-sections.component';
import { BlogContentDetailsComponent } from './pages/blog-content-details/blog-content-details.component';
import { FocusComponent } from './pages/focus/focus.component';

export const routes: Routes = [
  { path: '', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'focus', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'latest-posts', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'about', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'groovejr', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'blog', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'blog/:id', component: BlogContentDetailsComponent },
  { path: '**', redirectTo: '' }
];
