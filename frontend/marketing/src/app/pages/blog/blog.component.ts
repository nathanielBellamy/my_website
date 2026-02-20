import { Component, OnInit, inject, OnDestroy } from '@angular/core';
import { BlogStore } from './blog.store';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { CardComponent } from '../../components/card/card.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { Tag } from '../../models/blog-post.model';
import { Subject, debounceTime, distinctUntilChanged, takeUntil } from 'rxjs';

@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [CommonModule, InfiniteScrollComponent, CardComponent, ScrollFadeInDirective],
  templateUrl: './blog.component.html',
})
export class BlogComponent implements OnInit, OnDestroy {
  protected readonly store = inject(BlogStore);
  private readonly router = inject(Router);
  private readonly searchSubject = new Subject<string>();
  private readonly destroy$ = new Subject<void>();

  ngOnInit() {
    this.store.loadMore();
    this.store.loadTags();

    this.searchSubject.pipe(
      debounceTime(500),
      distinctUntilChanged(),
      takeUntil(this.destroy$)
    ).subscribe(query => {
      this.store.searchTags(query);
    });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }

  onScroll() {
    this.store.loadMore();
  }

  viewPost(id: string) {
    this.router.navigate(['/blog', id]);
  }

  getSnippet(content: string): string {
    if (!content) return '';
    // simple snippet: first sentence or first 150 chars
    const firstPeriod = content.indexOf('.');
    if (firstPeriod > -1 && firstPeriod < 200) {
        return content.substring(0, firstPeriod + 1);
    }
    return content.length > 150 ? content.substring(0, 150) + '...' : content;
  }

  getTags(tags: Tag[]): string[] {
    return tags ? tags.map(t => t.name) : [];
  }

  onSearchTags(event: Event) {
    const input = event.target as HTMLInputElement;
    this.searchSubject.next(input.value);
  }

  onToggleTag(tagId: string) {
    this.store.toggleTag(tagId);
  }
}
