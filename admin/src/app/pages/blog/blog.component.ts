import { Component, inject, OnInit, signal } from '@angular/core';
import { BlogService } from '../../services/blog';
import { BlogPost, Tag } from '../../models/data-models'; // Import Tag
import { RouterLink } from '@angular/router'; // Import RouterLink for navigation

@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [RouterLink], // Removed AsyncPipe
  templateUrl: './blog.component.html',
  styleUrl: './blog.component.css',
})
export class BlogComponent implements OnInit {
  private readonly blogService = inject(BlogService);
  blogPosts = signal<BlogPost[]>([]);

  ngOnInit() {
    this.fetchBlogPosts();
  }

  fetchBlogPosts() {
    this.blogService.getAllBlogPosts().then((posts) => {
      this.blogPosts.set(posts);
    }).catch((error) => {
      console.error('Error fetching blog posts:', error);
    });
  }

  deletePost(id: string) {
    this.blogService.deleteBlogPost(id).then(() => {
      this.fetchBlogPosts(); // Refresh the list after deletion
    }).catch((error) => {
      console.error('Error deleting blog post:', error);
    });
  }

  getTagsAsString(tags: Tag[] | undefined): string {
    return tags ? tags.map(tag => tag.name).join(', ') : '';
  }
}
