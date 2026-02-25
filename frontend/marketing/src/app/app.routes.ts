import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { FocusComponent } from './pages/focus/focus.component';
import { LatestPostsComponent } from './pages/latest-posts/latest-posts.component';
import { AboutComponent } from './pages/about/about.component';
import { GrooveJrComponent } from './pages/groove-jr/groove-jr.component';
import { OldSiteComponent } from './pages/old-site/old-site.component';
import { BlogComponent } from './pages/blog/blog.component';
import { BlogContentDetailsComponent } from './pages/blog-content-details/blog-content-details.component';
import { PrivacyPolicyComponent } from './pages/privacy-policy/privacy-policy.component';

export const routes: Routes = [
  { path: '', component: HomeComponent, title: 'Nate Schieber - Software Engineer', data: { animation: 0 } },
  { path: 'focus', component: FocusComponent, title: 'Focus - Nate Schieber', data: { animation: 1 } },
  { path: 'latest-posts', component: LatestPostsComponent, title: 'Latest Posts - Nate Schieber', data: { animation: 2 } },
  { path: 'about', component: AboutComponent, title: 'About - Nate Schieber', data: { animation: 3 } },
  { path: 'groovejr', component: GrooveJrComponent, title: 'Groove Jr. - Nate Schieber', data: { animation: 4 } },
  { path: 'old-site-preview', component: OldSiteComponent, title: 'Old Site - Nate Schieber', data: { animation: 5 } },
  { path: 'blog', component: BlogComponent, title: 'Blog - Nate Schieber', data: { animation: 6 } },
  { path: 'blog/:id', component: BlogContentDetailsComponent, title: 'Blog Post - Nate Schieber', data: { animation: 7 } },
  { path: 'privacy-policy', component: PrivacyPolicyComponent, title: 'Privacy Policy - Nate Schieber', data: { animation: 8 } },
  { path: '**', redirectTo: '' }
];

