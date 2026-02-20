import { Component, OnInit, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { MarkdownComponent } from 'ngx-markdown';
import { BlogService } from '../../services/blog.service';
import { BlogPost } from '../../models/blog-post.model';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-blog-content-details',
  standalone: true,
  imports: [CommonModule, MarkdownComponent, RouterLink, ScrollFadeInDirective],
  templateUrl: './blog-content-details.component.html',
})
export class BlogContentDetailsComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly blogService = inject(BlogService);

  blogPost = signal<BlogPost | null>(null);
  loading = signal(true);
  error = signal<string | null>(null);

  async ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    if (!id) {
      this.error.set('No blog post ID provided');
      this.loading.set(false);
      return;
    }

    try {
      const post = await this.blogService.getById(id);
      this.blogPost.set(post);
    } catch (err) {
      this.error.set('Failed to load blog post');
    } finally {
      this.loading.set(false);
    }
  }
}
