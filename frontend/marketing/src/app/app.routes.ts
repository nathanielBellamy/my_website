import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { FocusComponent } from './pages/focus/focus.component';
import { WorkComponent } from './pages/work/work.component';
import { AboutComponent } from './pages/about/about.component';
import { AboutContentDetailsComponent } from './pages/about-content-details/about-content-details.component';
import { GrooveJrComponent } from './pages/groove-jr/groove-jr.component';
import { GrooveJrContentDetailsComponent } from './pages/groove-jr-content-details/groove-jr-content-details.component';
import { WorkContentDetailsComponent } from './pages/work-content-details/work-content-details.component';
import { OldSiteComponent } from './pages/old-site/old-site.component';
import { BlogComponent } from './pages/blog/blog.component';
import { BlogContentDetailsComponent } from './pages/blog-content-details/blog-content-details.component';
import { PrivacyPolicyComponent } from './pages/privacy-policy/privacy-policy.component';

export const routes: Routes = [
  { path: '', component: HomeComponent, title: 'Nate Schieber - Software Engineer' },
  // { path: 'focus', component: FocusComponent, title: 'Focus - Nate Schieber' },
  { path: 'work', component: WorkComponent, title: 'Work - Nate Schieber' },
  { path: 'about', component: AboutComponent, title: 'About - Nate Schieber' },
  { path: 'about/:id', component: AboutContentDetailsComponent, title: 'About - Nate Schieber' },
  { path: 'groovejr', component: GrooveJrComponent, title: 'Groove Jr. - Nate Schieber' },
  { path: 'groovejr/:id', component: GrooveJrContentDetailsComponent, title: 'Groove Jr. - Nate Schieber' },
  { path: 'work/:id', component: WorkContentDetailsComponent, title: 'Nate Schieber - Software Engineer' },
  { path: 'old-site-preview', component: OldSiteComponent, title: 'Old Site - Nate Schieber' },
  { path: 'blog', component: BlogComponent, title: 'Blog - Nate Schieber' },
  { path: 'blog/:id', component: BlogContentDetailsComponent, title: 'Blog Post - Nate Schieber' },
  { path: 'privacy-policy', component: PrivacyPolicyComponent, title: 'Privacy Policy - Nate Schieber' },
  { path: '**', redirectTo: '' }
];
