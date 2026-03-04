import { Routes } from '@angular/router';
import { BlogComponent } from './pages/blog/blog-content-list/blog.component';
import { CreateBlogPostComponent } from './pages/blog/create-blog-post/create-blog-post.component';
import { EditBlogPostComponent } from './pages/blog/edit-blog-post/edit-blog-post.component';
import { HomeContentListComponent } from './pages/home/home-content-list/home-content-list.component';
import { CreateHomeContentComponent } from './pages/home/create-home-content/create-home-content.component';
import { EditHomeContentComponent } from './pages/home/edit-home-content/edit-home-content.component';
import { GrooveJrContentListComponent } from './pages/groovejr/groove-jr-content-list/groove-jr-content-list.component';
import { CreateGrooveJrContentComponent } from './pages/groovejr/create-groove-jr-content/create-groove-jr-content.component';
import { EditGrooveJrContentComponent } from './pages/groovejr/edit-groove-jr-content/edit-groove-jr-content.component';
import { AboutContentListComponent } from './pages/about/about-content-list/about-content-list.component';
import { CreateAboutContentComponent } from './pages/about/create-about-content/create-about-content.component';
import { EditAboutContentComponent } from './pages/about/edit-about-content/edit-about-content.component';
import { GalleryComponent } from './pages/gallery/gallery.component';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full'
  },
  {
    path: 'home',
    component: HomeContentListComponent,
  },
  {
    path: 'home/new',
    component: CreateHomeContentComponent,
  },
  {
    path: 'home/:id/edit',
    component: EditHomeContentComponent,
  },
  {
    path: 'groovejr',
    component: GrooveJrContentListComponent,
  },
  {
    path: 'groovejr/new',
    component: CreateGrooveJrContentComponent,
  },
  {
    path: 'groovejr/:id/edit',
    component: EditGrooveJrContentComponent,
  },
  {
    path: 'about',
    component: AboutContentListComponent,
  },
  {
    path: 'about/new',
    component: CreateAboutContentComponent,
  },
  {
    path: 'about/:id/edit',
    component: EditAboutContentComponent,
  },
  {
    path: 'blog',
    component: BlogComponent,
  },
  {
    path: 'blog/new',
    component: CreateBlogPostComponent,
  },
  {
    path: 'blog/:id/edit',
    component: EditBlogPostComponent,
  },
  {
    path: 'gallery',
    component: GalleryComponent,
  },
];