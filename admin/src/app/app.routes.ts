import { Routes } from '@angular/router';
import { BlogComponent } from './pages/blog/blog.component';
import { CreateBlogPostComponent } from './pages/create-blog-post/create-blog-post.component';
import { EditBlogPostComponent } from './pages/edit-blog-post/edit-blog-post.component';

export const routes: Routes = [
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
];
