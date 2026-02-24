import { Routes } from '@angular/router';
import { AllSectionsComponent } from './pages/all-sections/all-sections.component';
import { BlogContentDetailsComponent } from './pages/blog-content-details/blog-content-details.component';
import { PrivacyPolicyComponent } from './pages/privacy-policy/privacy-policy.component';

export const routes: Routes = [
  { path: '', component: AllSectionsComponent, title: 'Nate Schieber - Software Engineer', runGuardsAndResolvers: 'always' },
  { path: 'focus', component: AllSectionsComponent, title: 'Focus - Nate Schieber', runGuardsAndResolvers: 'always' },
  { path: 'latest-posts', component: AllSectionsComponent, title: 'Latest Posts - Nate Schieber', runGuardsAndResolvers: 'always' },
  { path: 'about', component: AllSectionsComponent, title: 'About - Nate Schieber', runGuardsAndResolvers: 'always' },
  { path: 'groovejr', component: AllSectionsComponent, title: 'Groove Jr. - Nate Schieber', runGuardsAndResolvers: 'always' },
  { path: 'old-site-preview', component: AllSectionsComponent, title: 'Old Site - Nate Schieber', runGuardsAndResolvers: 'always' },
  { path: 'blog', component: AllSectionsComponent, title: 'Blog - Nate Schieber', runGuardsAndResolvers: 'always' },
  { path: 'blog/:id', component: BlogContentDetailsComponent, title: 'Blog Post - Nate Schieber' },
  { path: 'privacy-policy', component: PrivacyPolicyComponent, title: 'Privacy Policy - Nate Schieber' },
  { path: '**', redirectTo: '' }
];
