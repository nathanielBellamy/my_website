import { Component, OnInit, inject, signal, effect, Renderer2 } from '@angular/core';
import { CommonModule, DOCUMENT } from '@angular/common';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { MarkdownComponent } from 'ngx-markdown';
import { Title, Meta } from '@angular/platform-browser';
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
  private readonly titleService = inject(Title);
  private readonly metaService = inject(Meta);
  private readonly renderer = inject(Renderer2);
  private readonly document = inject(DOCUMENT);

  blogPost = signal<BlogPost | null>(null);
  loading = signal(true);
  error = signal<string | null>(null);

  constructor() {
    effect(() => {
      const post = this.blogPost();
      if (post) {
        this.titleService.setTitle(`${post.title} - Nate Schieber`);
        const description = post.content.substring(0, 150).replace(/[#*`]/g, '') + '...';
        this.metaService.updateTag({ name: 'description', content: description });
        
        // Add Schema.org JSON-LD
        const script = this.renderer.createElement('script');
        script.type = 'application/ld+json';
        script.text = JSON.stringify({
          "@context": "https://schema.org",
          "@type": "BlogPosting",
          "headline": post.title,
          "datePublished": post.createdAt,
          "author": {
            "@type": "Person",
            "name": post.author.name
          },
          "articleBody": post.content
        });
        this.renderer.appendChild(this.document.head, script);
      }
    });
  }

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
