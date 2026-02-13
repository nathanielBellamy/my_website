import { Routes } from '@angular/router';
import { AllSectionsComponent } from './pages/all-sections/all-sections.component';

export const routes: Routes = [
  { path: '', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'about', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'groovejr', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: 'blog', component: AllSectionsComponent, runGuardsAndResolvers: 'always' },
  { path: '**', redirectTo: '' }
];
