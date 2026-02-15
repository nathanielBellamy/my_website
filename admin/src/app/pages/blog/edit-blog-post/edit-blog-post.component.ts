import { Component, inject, OnInit, signal } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BlogFormComponent } from '../../../components/blog-form/blog-form.component';
import { BlogService } from '../../../services/blog.service';
import { BlogPost } from '../../../models/data-models';

@Component({
  selector: 'app-edit-blog-post',
  standalone: true,
  imports: [BlogFormComponent],
  templateUrl: './edit-blog-post.component.html',
  styleUrl: './edit-blog-post.component.css',
})
export class EditBlogPostComponent implements OnInit {
  private readonly blogService = inject(BlogService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  blogPost = signal<BlogPost | undefined>(undefined);

  ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.blogService.getBlogPostById(id).then((post) => {
        this.blogPost.set(post);
      }).catch((error) => {
        console.error('Error fetching blog post:', error);
      });
    }
  }

  async updatePost(post: BlogPost) {
    try {
      await this.blogService.updateBlogPost(post);
      await this.router.navigate(['/blog']);
    } catch (error) {
      console.error('Error updating blog post:', error);
    }
  }

  async goBack() {
    await this.router.navigate(['/blog']);
  }
}
