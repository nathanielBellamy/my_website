import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { BlogFormComponent } from '../../../components/blog-form/blog-form.component';
import { BlogService } from '../../../services/blog.service';
import { BlogPost } from '../../../models/data-models';

@Component({
  selector: 'app-create-blog-post',
  standalone: true,
  imports: [BlogFormComponent],
  templateUrl: './create-blog-post.component.html',
  styleUrl: './create-blog-post.component.css',
})
export class CreateBlogPostComponent {
  private readonly blogService = inject(BlogService);
  private readonly router = inject(Router);

  async createPost(post: BlogPost) {
    try {
      await this.blogService.createBlogPost(post);
      await this.router.navigate(['/blog']);
    } catch (error) {
      console.error('Error creating blog post:', error);
    }
  }

  async goBack() {
    await this.router.navigate(['/blog']);
  }
}
