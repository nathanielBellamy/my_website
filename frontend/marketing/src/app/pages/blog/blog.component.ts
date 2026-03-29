import { Component, OnInit, OnDestroy, inject } from '@angular/core';
import { BlogStore } from './blog.store';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { CardComponent } from '../../components/card/card.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { Tag } from '../../models/blog-post.model';
import { Subject, debounceTime, distinctUntilChanged, takeUntil } from 'rxjs';
import { PageComponent } from '../../components/page/page.component';
import { encodeId } from '../../utils/id-encoder';
import { getSnippet } from '../../utils/snippet';

@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [CommonModule, InfiniteScrollComponent, CardComponent, ScrollFadeInDirective, PageComponent],
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
    this.router.navigate(['/blog', encodeId(id)]);
  }

  getSnippet(content: string): string {
    return getSnippet(content);
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
