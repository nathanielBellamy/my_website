import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { AboutComponent } from './pages/about/about.component';
import { GrooveJrComponent } from './pages/groove-jr/groove-jr.component';
import { BlogComponent } from './pages/blog/blog.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'about', component: AboutComponent },
  { path: 'groovejr', component: GrooveJrComponent },
  { path: 'blog', component: BlogComponent },
  { path: '**', redirectTo: '' } // Redirect any unknown paths to home
];
