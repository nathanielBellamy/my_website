import { Component, OnInit, inject } from '@angular/core';
import { AboutStore } from './about.store';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { CardComponent } from '../../components/card/card.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { PageComponent } from '../../components/page/page.component';

@Component({
  selector: 'app-about',
  standalone: true,
  imports: [CommonModule, InfiniteScrollComponent, CardComponent, ScrollFadeInDirective, PageComponent],
  templateUrl: './about.component.html',
})
export class AboutComponent implements OnInit {
  protected readonly store = inject(AboutStore);
  private readonly router = inject(Router);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }

  viewContent(id: string) {
    this.router.navigate(['/about', id]);
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
