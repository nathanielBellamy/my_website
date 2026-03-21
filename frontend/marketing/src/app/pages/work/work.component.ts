import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { CardComponent } from '../../components/card/card.component';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { WorkStore } from './work.store';
import { PageComponent } from '../../components/page/page.component';
import { encodeId } from '../../utils/id-encoder';
import { getSnippet } from '../../utils/snippet';

@Component({
  selector: 'app-work',
  standalone: true,
  imports: [CommonModule, CardComponent, InfiniteScrollComponent, ScrollFadeInDirective, PageComponent],
  templateUrl: './work.component.html',
})
export class WorkComponent implements OnInit {
  protected readonly store = inject(WorkStore);
  private readonly router = inject(Router);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }

  viewContent(id: string) {
    this.router.navigate(['/work', encodeId(id)]);
  }

  getSnippet(content: string): string {
    return getSnippet(content);
  }
}
