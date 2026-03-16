import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { CardComponent } from '../../components/card/card.component';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { LatestPostsStore } from './latest-posts.store';
import { PageComponent } from '../../components/page/page.component';

@Component({
  selector: 'app-latest-posts',
  standalone: true,
  imports: [CommonModule, CardComponent, InfiniteScrollComponent, ScrollFadeInDirective, PageComponent],
  templateUrl: './latest-posts.component.html',
})
export class LatestPostsComponent implements OnInit {
  protected readonly store = inject(LatestPostsStore);
  private readonly router = inject(Router);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }

  viewContent(id: string) {
    this.router.navigate(['/home', id]);
  }

  getSnippet(content: string): string {
    if (!content) return '';
    const firstPeriod = content.indexOf('.');
    if (firstPeriod > -1 && firstPeriod < 200) {
      return content.substring(0, firstPeriod + 1);
    }
    return content.length > 150 ? content.substring(0, 150) + '...' : content;
  }
}
